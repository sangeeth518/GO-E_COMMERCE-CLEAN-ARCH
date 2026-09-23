package main

import (
	"fmt"
	"log"

	"github.com/redis/go-redis/v9"
	"github.com/sangeeth518/notification-service/pkg/config"
	"github.com/sangeeth518/notification-service/pkg/consumer"
	"github.com/sangeeth518/notification-service/pkg/repository"
	"github.com/sangeeth518/notification-service/pkg/service"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("can't load config", err)
	}
	log.Printf("RabbitMQ URL: '%s'", cfg.RabbitMQURL)
	redisClient := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
	})
	notfRepo := repository.NewNotificationRepository(redisClient)
	emailSvc := service.NewEmailService(cfg)

	c, err := consumer.NewConsumer(cfg.RabbitMQURL, emailSvc, notfRepo)
	if err != nil {
		log.Fatal("cannot create consumer ", err)
	}

	log.Println("Notification service started")

	if err := c.StartConsuming(); err != nil {
		log.Fatal("Consumer error", err)
	}

}
