package db

import (
	"database/sql"
	"ijah-shop/domain"
	"log"

	uuid "github.com/satori/go.uuid"
)

type Querier interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
}

type DB struct {
	*sql.DB
}

func (db *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return db.Exec(query, args)
}

func (db *DB) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return db.Query(query, args)
}

func (db *DB) QueryRow(query string, args ...interface{}) *sql.Row {
	return db.QueryRow(query, args)
}

type Tx struct {
	*sql.Tx
}

func (tx *Tx) Exec(query string, args ...interface{}) (sql.Result, error) {
	return tx.Exec(query, args)
}

func (tx *Tx) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return tx.Query(query, args)
}

func (tx *Tx) QueryRow(query string, args ...interface{}) *sql.Row {
	return tx.QueryRow(query, args)
}

func (db *DB) Begin() (*Tx, error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return nil, err
	}
	return &Tx{tx}, nil
}

func NewMysql() *DB {
	db, err := sql.Open("mysql", "root:root@/gotest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

type ProductRepository interface {
	InitTx(tx *Tx) ProductRepository

	Save(product *domain.Product) error
	FindAll() ([]*domain.Product, error)
	FindBySKU(sku string) (*domain.Product, error)
}

type IncomingProductRepository interface {
	InitTx(tx *Tx) IncomingProductRepository
	IsForUpdate(val bool) IncomingProductRepository

	Save(product *domain.IncomingProduct) error
	FindAll() ([]*domain.IncomingProduct, error)
	FindByID(id string) ([]*domain.IncomingProduct, error)
	FindByIDAndSKU(id, sku string) (*domain.IncomingProduct, error)
	FindByUID(uid uuid.UUID) (*domain.IncomingProduct, error)
}
