package main

import (
	"log"
	"net/http"

	"github.com/YanAmorelli/financial_control/database"
	"github.com/YanAmorelli/financial_control/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func (db *DBClient) PostInOut(c echo.Context) error {
	t := new(models.Transaction)
	if err := c.Bind(&t); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Create(&t).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, t)
}

func main() {
	e := echo.New()

	dsn := "host=localhost port=5432 user= password= dbname="
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbClient := DBClient{DB: db}

	e.POST("/transaction", dbClient.PostInOut)

	e.Logger.Fatal(e.Start(":8080"))
}
