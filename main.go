package main

import (
	"log"

	"github.com/YanAmorelli/financial_control/database"
	"github.com/YanAmorelli/financial_control/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dsn := "host=localhost port=5432 user=postgres password=123456789 dbname=postgres"
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbClient := handlers.DBClient{DB: db}

	// e.POST("/newuser", dbClient.NewUser)
	e.POST("/transaction", dbClient.CreateEntries)
	// e.GET("/transaction/:id", dbClient.ReadEntriesByID)
	// e.PUT("/transaction/:id", dbClient.UpdateEntries)
	// e.DELETE("/transaction/:id", dbClient.DeleteEntries)

	e.Logger.Fatal(e.Start(":8080"))
}
