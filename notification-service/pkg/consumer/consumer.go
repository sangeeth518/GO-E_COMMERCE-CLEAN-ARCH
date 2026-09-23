package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sangeeth518/notification-service/pkg/models"
	"github.com/sangeeth518/notification-service/pkg/repository"
	"github.com/sangeeth518/notification-service/pkg/service"
)

type Consumer struct {
	channel      *amqp.Channel
	emailService service.EmailService
	repo         repository.NotificationRepository
}

func NewConsumer(amqpURL string, emailService service.EmailService, repo repository.NotificationRepository) (*Consumer, error) {
	var conn *amqp.Connection
	var err error

	for i := 1; i <= 5; i++ {
		conn, err = amqp.Dial(amqpURL)
		if err == nil {
			break
		}
		log.Printf(" RabbitMQ connection attempt %d/5 failed: %v", i, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, err
	}
	_, err = ch.QueueDeclare(
		"notification_queue",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, err
	}

	log.Println("✅ Notification consumer connected to RabbitMQ")

	return &Consumer{channel: ch, emailService: emailService, repo: repo}, nil
}

func (c *Consumer) StartConsuming() error {

	msgs, err := c.channel.Consume(
		"notification_queue", "",
		false, false, false, false, nil,
	)
	if err != nil {
		return err
	}
	for message := range msgs {
		msg := message
		go func() {
			var event models.OrderPlacedEvent
			if err := json.Unmarshal(msg.Body, &event); err != nil {
				log.Printf("failed to unmarshal: %v", err)
				msg.Nack(false, false)
				return

			}
			if err := c.emailService.SendOrderConfirmation(event.UserEmail, event.UserName, event.OrderId, event.TotalAmount); err != nil {
				log.Printf("❌ email failed: %v", err)
				msg.Nack(false, true) // requeue — retry
				return
			}

			record := models.NotificationRecord{
				UserId:  event.UserId,
				Type:    "order_placed",
				Message: fmt.Sprintf("Order #%d confirmed for ₹%s", event.OrderId, event.UserEmail),
				OrderId: event.OrderId,
				SentAt:  time.Now(),
				Status:  "sent",
			}
			if err := c.repo.SaveNotification(record); err != nil {
				log.Printf("failed to save notification: %v", err)
			}
			msg.Ack(false)
			log.Printf("✅ notification sent to %s for order #%d", event.UserEmail, event.OrderId)
		}()

	}
	return nil
}
