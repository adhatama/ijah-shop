package mysql

import (
	"ijah-shop/db"
	"ijah-shop/domain"

	uuid "github.com/satori/go.uuid"
)

type OutgoingProductRepositoryMysql struct {
	Querier db.Querier
}

func (r *OutgoingProductRepositoryMysql) Save(outProd *domain.OutgoingProduct, config ...*db.RepoConfig) error {
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
		})
	}

	_, err := querier.Exec(`INSERT INTO outgoing_product
		(uid, id, sku, quantity, price, type, description, date)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?) ON DUPLICATE KEY UPDATE
		id = ?, sku = ?, quantity = ?, price = ?, type = ?, description = ?, date = ?`,
		outProd.UID, outProd.ID, outProd.SKU, outProd.Quantity, outProd.Price,
		outProd.Type, outProd.Description, outProd.Date,
		outProd.ID, outProd.SKU, outProd.Quantity, outProd.Price,
		outProd.Type, outProd.Description, outProd.Date)
	if err != nil {
		return err
	}

	return nil
}

func (r *OutgoingProductRepositoryMysql) FindAll() ([]*domain.OutgoingProduct, error) {
	outProds := []*domain.OutgoingProduct{}

	rows, err := r.Querier.Query(`SELECT uid, id, sku, quantity, price, type, description, date
		FROM outgoing_product`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		outProd := &domain.OutgoingProduct{}
		var uidResult string

		rows.Scan(&uidResult, &outProd.ID, &outProd.SKU, &outProd.Quantity, &outProd.Price,
			&outProd.Type, &outProd.Description, &outProd.Date)

		uid, _ := uuid.FromString(uidResult)
		outProd.UID = uid

		outProds = append(outProds, outProd)
	}

	return outProds, nil
}

func (r *OutgoingProductRepositoryMysql) FindByID(id string, config ...*db.RepoConfig) ([]*domain.OutgoingProduct, error) {
	query := `SELECT uid, id, sku, quantity, price, type, description, date
		FROM outgoing_product WHERE id = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	outProds := []*domain.OutgoingProduct{}

	rows, err := querier.Query(query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		outProd := &domain.OutgoingProduct{}
		var uidResult string

		rows.Scan(&uidResult, &outProd.ID, &outProd.SKU, &outProd.Quantity, &outProd.Price,
			&outProd.Type, &outProd.Description, &outProd.Date)

		uid, _ := uuid.FromString(uidResult)
		outProd.UID = uid

		outProds = append(outProds, outProd)
	}

	return outProds, nil
}

func (r *OutgoingProductRepositoryMysql) FindByIDAndSKU(id, sku string, config ...*db.RepoConfig) (*domain.OutgoingProduct, error) {
	query := `SELECT uid, id, sku, quantity, price, type, description, date
		FROM outgoing_product WHERE id = ? AND sku = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	outProd := &domain.OutgoingProduct{}
	var uidResult string

	err := querier.QueryRow(query, id, sku).Scan(
		&uidResult, &outProd.ID, &outProd.SKU, &outProd.Quantity, &outProd.Price,
		&outProd.Type, &outProd.Description, &outProd.Date,
	)
	if err != nil {
		return nil, err
	}

	uid, _ := uuid.FromString(uidResult)
	outProd.UID = uid

	return outProd, nil
}

func (r *OutgoingProductRepositoryMysql) FindByUID(uid uuid.UUID, config ...*db.RepoConfig) (*domain.OutgoingProduct, error) {
	query := `SELECT uid, id, sku, quantity, price, type, description, date
		FROM outgoing_product WHERE uid = ?`
	querier := r.Querier

	for _, v := range config {
		v.Apply(&db.RepoConfigValue{
			Querier: &querier,
			Query:   &query,
		})
	}

	outProd := &domain.OutgoingProduct{}
	var uidResult string

	err := querier.QueryRow(query, uid.String()).Scan(
		&uidResult, &outProd.ID, &outProd.SKU, &outProd.Quantity, &outProd.Price,
		&outProd.Type, &outProd.Description, &outProd.Date,
	)
	if err != nil {
		return nil, err
	}

	outProd.UID = uid

	return outProd, nil
}
