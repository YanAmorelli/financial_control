package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/YanAmorelli/financial_control/database"
	"github.com/YanAmorelli/financial_control/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
	Tr models.Transaction
}

func (db *DBClient) CreateInOut(c echo.Context) error {
	transaction := new(models.Transaction)
	if err := c.Bind(&transaction); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Create(&transaction).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, transaction)
}

func (db *DBClient) ReadInOutByID(c echo.Context) error {
	var transaction models.Transaction
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	if err := db.DB.First(&transaction, transactionId).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (db *DBClient) UpdateInOut(c echo.Context) error {
	var transaction models.Transaction
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	if err = db.DB.First(&transaction, transactionId).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	if err = c.Bind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err = db.DB.Save(transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (db *DBClient) DeleteInOut(c echo.Context) error {
	var transaction models.Transaction
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	if err = db.DB.Delete(&transaction, transactionId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Success: Deleted")
}

func main() {
	e := echo.New()

	dsn := "host=localhost port=5432 user= password= dbname="
	db, err := database.ConnectDatabase(dsn)
	if err != nil {
		log.Fatal(err)
	}

	dbClient := DBClient{DB: db}

	e.POST("/transaction", dbClient.CreateInOut)
	e.GET("/transaction/:id", dbClient.ReadInOutByID)
	e.PUT("/transaction/:id", dbClient.UpdateInOut)
	e.DELETE("/transaction/:id", dbClient.DeleteInOut)

	e.Logger.Fatal(e.Start(":8080"))
}
