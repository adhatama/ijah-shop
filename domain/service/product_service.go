package service

import "ijah-shop/db"

type ProductService struct {
	ProductRepo  db.ProductRepository
	OrderRepo    db.IncomingProductRepository
	OutgoingRepo db.OutgoingProductRepository
}

func (s ProductService) IsSKUAlreadyCreated(sku string) bool {
	prod, _ := s.ProductRepo.FindBySKU(sku)
	if prod != nil {
		return true
	}

	return false
}

func (s ProductService) IsOrderIDAlreadyCreated(id string) bool {
	order, _ := s.OrderRepo.FindByID(id)
	if order != nil {
		return true
	}

	return false
}

func (s ProductService) IsOutgoingIDAlreadyCreated(id string) bool {
	out, _ := s.OutgoingRepo.FindByID(id)
	if out != nil {
		return true
	}

	return false
}
