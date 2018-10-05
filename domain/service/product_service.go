package service

import "ijah-shop/db"

type ProductService struct {
	ProductRepo db.ProductRepository
}

func (s ProductService) IsSKUAlreadyCreated(sku string) bool {
	return false
}
