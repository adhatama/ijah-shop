package mysql

import (
	"ijah-shop/db"
	"ijah-shop/domain"

	uuid "github.com/satori/go.uuid"
)

type IncomingProductRepositoryMysql struct {
	Querier   db.Querier
	ForUpdate bool
}

func (r IncomingProductRepositoryMysql) InitTx(tx *db.Tx) db.IncomingProductRepository {
	r.Querier = tx.Tx

	return &r
}

func (r IncomingProductRepositoryMysql) IsForUpdate(val bool) db.IncomingProductRepository {
	r.ForUpdate = val

	return &r
}

func (r IncomingProductRepositoryMysql) Save(incProd *domain.IncomingProduct) error {
	_, err := r.Querier.Exec(`INSERT INTO incoming_product
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

func (r IncomingProductRepositoryMysql) FindAll() ([]*domain.IncomingProduct, error) {
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

func (r IncomingProductRepositoryMysql) FindByID(id string) ([]*domain.IncomingProduct, error) {
	incProds := []*domain.IncomingProduct{}

	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
	FROM incoming_product WHERE id = ?`

	if r.ForUpdate {
		query += " FOR UPDATE"
	}

	rows, err := r.Querier.Query(query, id)
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

func (r IncomingProductRepositoryMysql) FindByIDAndSKU(id, sku string) (*domain.IncomingProduct, error) {
	incProd := &domain.IncomingProduct{}
	var uidResult string

	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
	FROM incoming_product WHERE id = ? AND sku = ?`

	if r.ForUpdate {
		query += " FOR UPDATE"
	}

	err := r.Querier.QueryRow(query, id, sku).Scan(
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

func (r IncomingProductRepositoryMysql) FindByUID(uid uuid.UUID) (*domain.IncomingProduct, error) {
	incProd := &domain.IncomingProduct{}
	var uidResult string

	query := `SELECT uid, id, sku, quantity, received_quantity, created_date, status
	FROM incoming_product WHERE uid = ?`

	if r.ForUpdate {
		query += " FOR UPDATE"
	}

	err := r.Querier.QueryRow(query, uid.String()).Scan(
		&uidResult, &incProd.ID, &incProd.SKU, &incProd.Quantity, &incProd.ReceivedQuantity,
		&incProd.CreatedDate, &incProd.Status,
	)
	if err != nil {
		return nil, err
	}

	incProd.UID = uid

	return incProd, nil
}
