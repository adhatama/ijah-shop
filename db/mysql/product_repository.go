package mysql

import (
	"ijah-shop/db"
	"ijah-shop/domain"
)

type ProductRepositoryMysql struct {
	Querier db.Querier
}

func (r *ProductRepositoryMysql) Save(product *domain.Product, config ...*db.RepoConfig) error {
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
		})
	}

	_, err := querier.Exec(`INSERT INTO product (sku, name, available_quantity) VALUES (?, ?, ?)
	ON DUPLICATE KEY UPDATE sku = ?, name = ?, available_quantity = ?`,
		product.SKU, product.Name, product.AvailableStock,
		product.SKU, product.Name, product.AvailableStock)
	if err != nil {
		return err
	}

	return nil
}

func (r *ProductRepositoryMysql) FindAll() ([]*domain.Product, error) {
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

func (r *ProductRepositoryMysql) FindBySKU(sku string, config ...*db.RepoConfig) (*domain.Product, error) {
	query := `SELECT sku, name, available_quantity
		FROM product WHERE sku = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	p := &domain.Product{}
	err := querier.QueryRow(query, sku).Scan(&p.SKU, &p.Name, &p.AvailableStock)
	if err != nil {
		return nil, err
	}

	return p, nil
}
