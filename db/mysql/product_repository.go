package mysql

import (
	"ijah-shop/db"
	"ijah-shop/domain"
)

type ProductRepositoryMysql struct {
	Querier   db.Querier
	ForUpdate bool
}

func (r ProductRepositoryMysql) InitTx(tx *db.Tx) db.ProductRepository {
	r.Querier = tx.Tx

	return &r
}

func (r ProductRepositoryMysql) Save(product *domain.Product) error {
	_, err := r.Querier.Exec(`INSERT INTO product (sku, name, available_quantity) VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE sku = ?, name = ?, available_quantity = ?`,
		product.SKU, product.Name, product.AvailableStock,
		product.SKU, product.Name, product.AvailableStock)
	if err != nil {
		return err
	}

	return nil
}

func (r ProductRepositoryMysql) FindAll() ([]*domain.Product, error) {
	products := []*domain.Product{}

	rows, err := r.Querier.Query("SELECT sku, name, available_quantity FROM product")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		p := &domain.Product{}

		rows.Scan(&p.SKU, &p.Name, &p.AvailableStock)

		products = append(products, p)
	}

	return products, nil
}

func (r ProductRepositoryMysql) FindBySKU(sku string) (*domain.Product, error) {
	p := &domain.Product{}
	err := r.Querier.QueryRow(`SELECT sku, name, available_quantity
		FROM product WHERE sku = ? FOR UPDATE`, sku).Scan(&p.SKU, &p.Name, &p.AvailableStock)
	if err != nil {
		return nil, err
	}

	return p, nil
}
