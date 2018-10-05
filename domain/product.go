package domain

type Product struct {
	SKU            string
	Name           string
	AvailableStock int
}

type ProductService interface {
	IsSKUAlreadyCreated(sku string) bool
}

func NewProduct(productService ProductService, sku, name string, availableStock int) (*Product, error) {
	return &Product{
		SKU:            sku,
		Name:           name,
		AvailableStock: availableStock,
	}, nil
}

func (p *Product) ChangeName(name string) error {
	p.Name = name

	return nil
}

func (p *Product) AddStock(quantity int) error {
	p.AvailableStock = p.AvailableStock + quantity

	return nil
}
