package publisher

import "github.com/sangeeth518/go-Ecommerce/pkg/utils/models"

type Publisher interface {
	PublishOrderPlaced(event models.OrderPlacedEvent) error
}
