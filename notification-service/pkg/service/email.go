package service

import (
	"fmt"
	"net/smtp"

	"github.com/sangeeth518/notification-service/pkg/config"
)

type EmailService interface {
	SendOrderConfirmation(to, userName string, orderId int, amount float64) error
}

type emailservice struct {
	config config.Config
}

func NewEmailService(cfg config.Config) EmailService {
	return &emailservice{
		config: cfg,
	}
}

func (e *emailservice) SendOrderConfirmation(to, userName string, orderId int, amount float64) error {
	subject := fmt.Sprintf("Order #%d Confirmed — BuyItNow", orderId)
	body := fmt.Sprintf(`
Hi %s,

Your order has been placed successfully!

Order ID:     #%d
Total Amount: ₹%.2f

Thank you for shopping with BuyItNow.

Regards,
BuyItNow Team
`, userName, orderId, amount)

	return e.sendEmail(to, subject, body)

}

func (e *emailservice) sendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth(
		"",
		e.config.SMTPEmail,
		e.config.SMTPPassword,
		e.config.SMTPHost,
	)
	msg := fmt.Sprintf(
		"From: %s <%s>\r\nTo: %s\r\nSubject: %s\r\nMIME-version: 1.0\r\nContent-Type: text/plain; charset=\"UTF-8\"\r\n\r\n%s",
		e.config.SMTPName,
		e.config.SMTPEmail,
		to,
		subject,
		body,
	)
	addr := fmt.Sprintf("%s:%d", e.config.SMTPHost, e.config.SMTPPort)

	err := smtp.SendMail(addr, auth, e.config.SMTPEmail, []string{to}, []byte(msg))
	if err != nil {
		return fmt.Errorf("failed to send email to %s: %w", to, err)
	}

	fmt.Printf("✅ Email sent to %s — %s\n", to, subject)
	return nil

}
