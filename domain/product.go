package domain

import (
	"errors"
	"time"
)

type Product struct {
	SKU            string
	Name           string
	AvailableStock int
	Orders         []Order
	Outgoings      []Outgoing
}

type Order struct {
	ID                string
	Quantity          int
	ReceivedQuantity  int
	Price             int
	Date              time.Time
	Status            OrderStatus
	LastOrderReceived time.Time
}

type Outgoing struct {
	ID          string
	Quantity    int
	Price       int
	Type        OutgoingType
	Description string
	Date        time.Time
}

type OrderStatus string

const (
	OrderStatusOrdered           OrderStatus = "ORDERED"
	OrderStatusPartiallyReceived OrderStatus = "PARTIALLY_RECEIVED"
	OrderStatusCompleted         OrderStatus = "COMPLETED"
)

type OutgoingType string

const (
	OutgoingStatusSales   OutgoingType = "SALES"
	OutgoingStatusProblem OutgoingType = "PROBLEM"
	OutgoingStatusSample  OutgoingType = "SAMPLE"
)

type ProductService interface {
	IsSKUAlreadyCreated(sku string) bool
	IsOrderIDAlreadyCreated(id string) bool
	IsOutgoingIDAlreadyCreated(id string) bool
}

func NewProduct(productService ProductService, sku, name string) (*Product, error) {
	if productService.IsSKUAlreadyCreated(sku) {
		return nil, errors.New("SKU is already created")
	}

	if name == "" {
		return nil, errors.New("Name cannot be empty")
	}

	return &Product{
		SKU:  sku,
		Name: name,
	}, nil
}

func (p *Product) ChangeName(name string) error {
	if name == "" {
		return errors.New("Name cannot be empty")
	}

	p.Name = name

	return nil
}

func (p *Product) NewOrder(productService ProductService, id string, quantity, price int, date time.Time) error {
	if productService.IsOrderIDAlreadyCreated(id) {
		return errors.New("Order ID is already created")
	}

	order := Order{
		ID:       id,
		Quantity: quantity,
		Price:    price,
		Status:   OrderStatusOrdered,
		Date:     date,
	}

	p.Orders = append(p.Orders, order)

	return nil
}

func (p *Product) ReceiveOrder(id string, quantity int, date time.Time) error {
	isIDFound := false

	for i, v := range p.Orders {
		if v.ID == id {
			p.Orders[i].ReceivedQuantity += quantity
			p.Orders[i].LastOrderReceived = date
			isIDFound = true
		}
	}

	if !isIDFound {
		return errors.New("Order ID is not found")
	}

	return nil
}

func (p *Product) Sell(productService ProductService, id string, quantity, price int, description string, date time.Time) error {
	if productService.IsOutgoingIDAlreadyCreated(id) {
		return errors.New("Outgoing ID is already created")
	}

	out := Outgoing{
		ID:          id,
		Quantity:    quantity,
		Price:       price,
		Type:        OutgoingStatusSales,
		Description: description,
		Date:        date,
	}

	p.Outgoings = append(p.Outgoings, out)

	return nil
}

func (p *Product) ProblemArise(quantity int, description string, date time.Time) error {
	out := Outgoing{
		Quantity:    quantity,
		Type:        OutgoingStatusProblem,
		Description: description,
		Date:        date,
	}

	p.Outgoings = append(p.Outgoings, out)

	return nil
}

func (p *Product) GiveSample(quantity int, description string, date time.Time) error {
	out := Outgoing{
		Quantity:    quantity,
		Type:        OutgoingStatusSample,
		Description: description,
		Date:        date,
	}

	p.Outgoings = append(p.Outgoings, out)

	return nil
}
