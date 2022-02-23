package main

import (
	"log"

	"github.com/YanAmorelli/financial_control/database"
	inout "github.com/YanAmorelli/financial_control/handlers/InOut"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	dsn := "host=localhost port=5432 user= password= dbname="
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbClient := inout.DBClient{DB: db}

	e.POST("/transaction", dbClient.CreateInOut)
	e.GET("/transaction/:id", dbClient.ReadInOutByID)
	e.PUT("/transaction/:id", dbClient.UpdateInOut)
	e.DELETE("/transaction/:id", dbClient.DeleteInOut)

	e.Logger.Fatal(e.Start(":8080"))
}
