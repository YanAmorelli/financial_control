package inout

import (
	"log"
	"net/http"
	"strconv"

	"github.com/YanAmorelli/financial_control/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
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
