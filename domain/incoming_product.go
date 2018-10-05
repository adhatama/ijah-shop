package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type IncomingProduct struct {
	UID              uuid.UUID
	ID               string
	SKU              string
	Quantity         int
	ReceivedQuantity int
	Price            int
	CreatedDate      time.Time
	Status           string
	History          []IncomingProductHistory
}

type IncomingProductHistory struct {
	ReceivedQuantity int
	ReceivedDate     time.Time
}

func NewIncomingProduct(id, sku string, quantity, price int, createdDate time.Time) (*IncomingProduct, error) {
	uid, _ := uuid.NewV4()

	return &IncomingProduct{
		UID:         uid,
		ID:          id,
		SKU:         sku,
		Quantity:    quantity,
		Price:       price,
		Status:      "CREATED",
		CreatedDate: createdDate,
	}, nil
}

func (ip *IncomingProduct) ReceiveOrder(quantity int, date time.Time) error {
	totalQuantity := ip.ReceivedQuantity + quantity

	ip.History = append(ip.History, IncomingProductHistory{
		ReceivedQuantity: quantity,
		ReceivedDate:     date,
	})

	ip.ReceivedQuantity = totalQuantity

	if ip.ReceivedQuantity >= ip.Quantity {
		ip.Status = "COMPLETED"
	} else {
		ip.Status = "PARTIALLY_RECEIVED"
	}

	return nil
}
