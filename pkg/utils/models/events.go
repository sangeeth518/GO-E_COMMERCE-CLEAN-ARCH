package models

import "time"

// OrderPlacedEvent is published to RabbitMQ when an order is placed
// Notification service reads this same struct
type OrderPlacedEvent struct {
	EventType   string    `json:"event_type"` // "order.placed"
	OrderId     int       `json:"order_id"`
	UserId      int       `json:"user_id"`
	UserEmail   string    `json:"user_email"`
	UserPhone   string    `json:"user_phone"`
	UserName    string    `json:"user_name"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
}

// UserSignupEvent is published when a new user signs up
type UserSignupEvent struct {
	EventType string `json:"event_type"` // "user.signup"
	UserId    int    `json:"user_id"`
	UserEmail string `json:"user_email"`
	UserName  string `json:"user_name"`
}
