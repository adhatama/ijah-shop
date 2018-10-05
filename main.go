package main

import (
	"ijah-shop/db"
	"ijah-shop/server"
	"log"

	_ "github.com/go-sql-driver/mysql"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Setup DB
	db := db.NewMysql()

	// Setup HTTP Server
	e := echo.New()

	// Middleware
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	s, err := server.NewServer(db)
	if err != nil {
		log.Fatal(err)
	}

	api := e.Group("/api")

	s.Mount(api)

	e.Logger.Fatal(e.Start(":1323"))
}
