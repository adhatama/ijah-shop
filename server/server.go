package server

import (
	"ijah-shop/db"
	"ijah-shop/db/mysql"
	"ijah-shop/domain"
	"ijah-shop/domain/service"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo"
	log "github.com/mgutz/logxi/v1"
)

type Server struct {
	DB                  *db.DB
	ProductRepo         db.ProductRepository
	ProductService      domain.ProductService
	IncomingProductRepo db.IncomingProductRepository
}

func NewServer(db *db.DB) (*Server, error) {
	productRepo := &mysql.ProductRepositoryMysql{
		Querier: db.DB,
	}

	incomingProductRepo := &mysql.IncomingProductRepositoryMysql{
		Querier: db.DB,
	}

	productService := &service.ProductService{
		ProductRepo: productRepo,
	}

	server := &Server{
		DB:                  db,
		ProductRepo:         productRepo,
		ProductService:      productService,
		IncomingProductRepo: incomingProductRepo,
	}

	return server, nil
}

func (s *Server) Mount(e *echo.Group) {
	e.POST("/products", s.SaveProduct)
	e.GET("/products", s.GetProducts)
	e.GET("/products/:sku", s.GetProductBySKU)

	e.POST("/incoming_products", s.SaveIncomingProduct)
	e.POST("/receive_orders", s.ReceiveOrder)
}

func (s *Server) SaveProduct(c echo.Context) error {
	sku := c.FormValue("sku")
	name := c.FormValue("name")
	initialQuantity, err := strconv.Atoi(c.FormValue("initial_quantity"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	product, err := domain.NewProduct(s.ProductService, sku, name, initialQuantity)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = s.ProductRepo.Save(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	data := make(map[string]interface{})
	data["data"] = product

	return c.JSON(http.StatusOK, data)
}

func (s *Server) GetProducts(c echo.Context) error {
	products, err := s.ProductRepo.FindAll()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	data := make(map[string]interface{})
	data["data"] = products

	return c.JSON(http.StatusOK, data)
}

func (s *Server) GetProductBySKU(c echo.Context) error {
	product, err := s.ProductRepo.FindBySKU(c.Param("sku"))
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}

	data := make(map[string]interface{})
	data["data"] = product

	return c.JSON(http.StatusOK, data)
}

func (s *Server) SaveIncomingProduct(c echo.Context) error {
	id := c.FormValue("id")
	sku := c.FormValue("sku")

	quantity, err := strconv.Atoi(c.FormValue("quantity"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	price, err := strconv.Atoi(c.FormValue("price"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	createdDate, err := time.Parse("2006-01-02", c.FormValue("created_date"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	incProd, err := domain.NewIncomingProduct(id, sku, quantity, price, createdDate)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = s.IncomingProductRepo.Save(incProd)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	data := make(map[string]interface{})
	data["data"] = incProd

	return c.JSON(http.StatusOK, data)
}

func (s *Server) ReceiveOrder(c echo.Context) error {
	id := c.FormValue("id")
	sku := c.FormValue("sku")

	quantity, err := strconv.Atoi(c.FormValue("quantity"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	date, err := time.Parse("2006-01-02", c.FormValue("date"))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Begin processing

	tx, err := s.DB.Begin()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	repoConfig := &db.RepoConfig{
		Tx: &db.Tx{tx},
	}

	// Gathers data

	incProd, err := s.IncomingProductRepo.FindByIDAndSKU(id, sku, repoConfig)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	prod, err := s.ProductRepo.FindBySKU(sku, repoConfig)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Do calculation for each domain

	err = incProd.ReceiveOrder(quantity, date)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = prod.AddStock(quantity)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	// Persists

	err = s.IncomingProductRepo.Save(incProd, repoConfig)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	err = s.ProductRepo.Save(prod, repoConfig)
	if err != nil {
		log.Error(err.Error())
		tx.Rollback()
		return c.JSON(http.StatusInternalServerError, err)
	}

	tx.Commit()

	// Compose response

	data := make(map[string]interface{})
	data["data"] = incProd

	return c.JSON(http.StatusOK, data)
}
