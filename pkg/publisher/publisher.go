package publisher

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sangeeth518/go-Ecommerce/pkg/config"
	"github.com/sangeeth518/go-Ecommerce/pkg/utils/models"
)

type rabbitMQPublisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

func NewPublisher(cfg config.Config) (Publisher, error) {
	var conn *amqp.Connection
	var err error

	for i := 1; i <= 5; i++ {
		conn, err = amqp.Dial(cfg.RabbitMQURL)
		if err == nil {
			break
		}
		log.Printf(" RabbitMQ connection attempt %d/5 failed: %v", i, err)
		time.Sleep(3 * time.Second)
	}
	if err != nil {
		log.Printf("RabbitMQ unavailable, Notifications disabled %v", err)
		return nil, nil
	}

	ch, err := conn.Channel()
	if err != nil {
		return nil, fmt.Errorf("faied to open channel: %w", err)
	}
	queueName := "notification_queue"

	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to declare queue %w", err)
	}

	return &rabbitMQPublisher{
		conn:    conn,
		channel: ch,
		queue:   queueName,
	}, nil

}

func (p *rabbitMQPublisher) PublishOrderPlaced(event models.OrderPlacedEvent) error {
	return p.publish(event)
}

func (p *rabbitMQPublisher) publish(body any) error {
	data, err := json.Marshal(body)
	if err != nil {
		return fmt.Errorf("failed to marshal event: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return p.channel.PublishWithContext(
		ctx,
		"",
		p.queue,
		false,
		false,
		amqp.Publishing{
			ContentType:  "application/json",
			DeliveryMode: amqp.Persistent,
			Body:         data,
		},
	)
}

func (p *rabbitMQPublisher) Close() {
	if p.conn != nil {
		p.channel.Close()
	}
	if p.channel != nil {
		p.conn.Close()
	}
}
