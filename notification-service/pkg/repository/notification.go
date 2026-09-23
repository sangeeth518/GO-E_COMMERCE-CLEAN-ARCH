package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/sangeeth518/notification-service/pkg/models"
)

type NotificationRepository interface {
	SaveNotification(record models.NotificationRecord) error
	GetUserNotifications(userId int) ([]models.NotificationRecord, error)
}

type notificationRepo struct {
	client *redis.Client
}

func NewNotificationRepository(client *redis.Client) NotificationRepository {
	return &notificationRepo{
		client: client,
	}
}

func (r *notificationRepo) SaveNotification(record models.NotificationRecord) error {
	ctx := context.Background()
	data, err := json.Marshal(record)
	if err != nil {
		return err
	}
	key := fmt.Sprintf("notifications:%d", record.UserId)

	err = r.client.LPush(ctx, key, data).Err()
	if err != nil {
		return err
	}
	r.client.LTrim(ctx, key, 0, 49)

	r.client.Expire(ctx, key, 30*24*time.Hour)
	return nil

}

func (r *notificationRepo) GetUserNotifications(userId int) ([]models.NotificationRecord, error) {
	ctx := context.Background()
	key := fmt.Sprintf("notifications:%d", userId)

	results, err := r.client.LRange(ctx, key, -1, 0).Result()
	if err != nil {
		return nil, err
	}

	var notifications []models.NotificationRecord
	for _, result := range results {
		var record models.NotificationRecord
		if err := json.Unmarshal([]byte(result), &record); err == nil {
			notifications = append(notifications, record)
		}
	}
	return notifications, nil
}
