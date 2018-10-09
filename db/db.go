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

type RepoConfig struct {
	Tx *Tx
}

type RepoConfigValue struct {
	Querier *Querier
	Query   *string
}

func (c *RepoConfig) Apply(value *RepoConfigValue) {
	if c.Tx != nil && value.Querier != nil {
		*value.Querier = c.Tx.Tx

		if value.Query != nil {
			q := *value.Query + " FOR UPDATE"
			*value.Query = q
		}
	}
}

func NewMysql() *DB {
	db, err := sql.Open("mysql", "root:root@/gotest?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}

	return &DB{db}
}

type ProductRepository interface {
	Save(product *domain.Product, config ...*RepoConfig) error
	FindAll() ([]*domain.Product, error)
	FindBySKU(sku string, config ...*RepoConfig) (*domain.Product, error)
}

type IncomingProductRepository interface {
	Save(product *domain.IncomingProduct, config ...*RepoConfig) error
	FindAll() ([]*domain.IncomingProduct, error)
	FindByID(id string, config ...*RepoConfig) ([]*domain.IncomingProduct, error)
	FindByIDAndSKU(id, sku string, config ...*RepoConfig) (*domain.IncomingProduct, error)
	FindByUID(uid uuid.UUID, config ...*RepoConfig) (*domain.IncomingProduct, error)
}

type OutgoingProductRepository interface {
	Save(product *domain.OutgoingProduct, config ...*RepoConfig) error
	FindAll() ([]*domain.OutgoingProduct, error)
	FindByID(id string, config ...*RepoConfig) ([]*domain.OutgoingProduct, error)
	FindByIDAndSKU(id, sku string, config ...*RepoConfig) (*domain.OutgoingProduct, error)
	FindByUID(uid uuid.UUID, config ...*RepoConfig) (*domain.OutgoingProduct, error)
}
