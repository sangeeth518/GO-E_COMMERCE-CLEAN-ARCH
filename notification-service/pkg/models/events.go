package models

import "time"

type OrderPlacedEvent struct {
	EventType   string    `json:"event_type"`
	OrderId     int       `json:"order_id"`
	UserId      int       `json:"user_id"`
	UserEmail   string    `json:"user_email"`
	UserPhone   string    `json:"user_phone"`
	UserName    string    `json:"user_name"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

// NotificationRecord stored in Redis
type NotificationRecord struct {
	UserId  int       `json:"user_id"`
	Type    string    `json:"type"`
	Message string    `json:"message"`
	OrderId int       `json:"order_id,omitempty"`
	SentAt  time.Time `json:"sent_at"`
	Status  string    `json:"status"`
}
