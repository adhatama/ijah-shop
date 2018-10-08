package mysql

import (
	"ijah-shop/db"
	"ijah-shop/domain"

	uuid "github.com/satori/go.uuid"
)

type IncomingProductRepositoryMysql struct {
	Querier db.Querier
}

func (r *IncomingProductRepositoryMysql) Save(incProd *domain.IncomingProduct, config ...*db.RepoConfig) error {
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
		})
	}

	_, err := querier.Exec(`INSERT INTO incoming_product
		(uid, id, sku, quantity, received_quantity, price, created_date, status)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE
		id = ?, sku = ?, quantity = ?, received_quantity = ?, price = ?, created_date = ?, status = ?`,
		incProd.UID, incProd.ID, incProd.SKU, incProd.Quantity, incProd.ReceivedQuantity,
		incProd.Price, incProd.CreatedDate, incProd.Status,
		incProd.ID, incProd.SKU, incProd.Quantity, incProd.ReceivedQuantity,
		incProd.Price, incProd.CreatedDate, incProd.Status)
	if err != nil {
		return err
	}

	return nil
}

func (r *IncomingProductRepositoryMysql) FindAll() ([]*domain.IncomingProduct, error) {
	incProds := []*domain.IncomingProduct{}

	rows, err := r.Querier.Query(`SELECT uid, id, sku, quantity, received_quantity, created_date, status
		FROM incoming_product`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		incProd := &domain.IncomingProduct{}
		var uidResult string

		rows.Scan(&uidResult, &incProd.ID, &incProd.SKU, &incProd.Quantity, &incProd.ReceivedQuantity,
			&incProd.CreatedDate, &incProd.Status)

		uid, _ := uuid.FromString(uidResult)
		incProd.UID = uid

		incProds = append(incProds, incProd)
	}

	return incProds, nil
}

func (r *IncomingProductRepositoryMysql) FindByID(id string, config ...*db.RepoConfig) ([]*domain.IncomingProduct, error) {
	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
		FROM incoming_product WHERE id = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	incProds := []*domain.IncomingProduct{}

	rows, err := querier.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		incProd := &domain.IncomingProduct{}
		var uidResult string

		rows.Scan(&uidResult, &incProd.ID, &incProd.SKU, &incProd.Quantity, &incProd.ReceivedQuantity,
			&incProd.CreatedDate, &incProd.Status)

		uid, _ := uuid.FromString(uidResult)
		incProd.UID = uid

		incProds = append(incProds, incProd)
	}

	return incProds, nil
}

func (r *IncomingProductRepositoryMysql) FindByIDAndSKU(id, sku string, config ...*db.RepoConfig) (*domain.IncomingProduct, error) {
	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
		FROM incoming_product WHERE id = ? AND sku = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	incProd := &domain.IncomingProduct{}
	var uidResult string

	err := querier.QueryRow(query, id, sku).Scan(
		&uidResult, &incProd.ID, &incProd.SKU, &incProd.Quantity, &incProd.ReceivedQuantity,
		&incProd.CreatedDate, &incProd.Status,
	)
	if err != nil {
		return nil, err
	}

	uid, _ := uuid.FromString(uidResult)
	incProd.UID = uid

	return incProd, nil
}

func (r *IncomingProductRepositoryMysql) FindByUID(uid uuid.UUID, config ...*db.RepoConfig) (*domain.IncomingProduct, error) {
	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
		FROM incoming_product WHERE uid = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	incProd := &domain.IncomingProduct{}
	var uidResult string

	err := querier.QueryRow(query, uid.String()).Scan(
		&uidResult, &incProd.ID, &incProd.SKU, &incProd.Quantity, &incProd.ReceivedQuantity,
		&incProd.CreatedDate, &incProd.Status,
	)
	if err != nil {
		return nil, err
	}

	incProd.UID = uid

	return incProd, nil
}
