package handlers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/YanAmorelli/financial_control/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func (db *DBClient) NewUser(c echo.Context) error {
	totalBalance := new(models.TotalBalance)
	totalBalance.Amount = 0
	if err := c.Bind(&totalBalance); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Table("total_balance").Create(&totalBalance).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	monthlyBalance := new(models.MonthlyBalance)
	monthlyBalance.Amount = 0
	monthlyBalance.YearMonth = time.Now().Format("2006-01-02")
	monthlyBalance.TotalId = 1

	if err := c.Bind(&monthlyBalance); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Table("monthly_balance").Create(&monthlyBalance).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Message: Success")
}

func (db *DBClient) CreateEntries(c echo.Context) error {
	entries := new(models.Entries)
	var balance models.MonthlyBalance

	if err := c.Bind(&entries); err != nil {
		log.Println("entries", err)
		return c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Table("entries").Create(&entries).Error; err != nil {
		log.Println(err)
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := db.DB.Table("monthly_balance").First(&balance, entries.MonthlyBalanceID).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	if balance.Id == entries.MonthlyBalanceID {
		balance.Amount += entries.Amount
	}

	if err := db.DB.Table("monthly_balance").Save(balance).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, "Success")
}

func (db *DBClient) ReadEntriesByID(c echo.Context) error {
	var transaction models.Entries
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	if err := db.DB.First(&transaction, transactionId).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	return c.JSON(http.StatusOK, transaction)
}

func (db *DBClient) UpdateEntries(c echo.Context) error {
	var transaction models.Entries
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

func (db *DBClient) DeleteEntries(c echo.Context) error {
	var transaction models.Entries
	transactionId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		log.Println(err)
	}

	if err = db.DB.Delete(&transaction, transactionId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, "Success: Deleted")
}
