package domain

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

type OutgoingProduct struct {
	UID         uuid.UUID
	ID          string
	SKU         string
	Quantity    int
	Price       int
	Type        string
	Description string
	Date        time.Time
}

func NewOutgoingProduct(id, sku string, quantity, price int,
	outgoingType, description string, date time.Time) (*OutgoingProduct, error) {

	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	return &OutgoingProduct{
		UID:         uid,
		ID:          id,
		SKU:         sku,
		Quantity:    quantity,
		Price:       price,
		Type:        outgoingType,
		Description: description,
		Date:        date,
	}, nil
}
