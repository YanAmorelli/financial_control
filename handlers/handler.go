package handlers

import (
	"net/http"
	"strconv"

	"github.com/YanAmorelli/financial_control/models"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type DBClient struct {
	DB *gorm.DB
}

func (db *DBClient) CreateEntries(c echo.Context) error {
	entries := new(models.Entries)
	mBalance := new(models.MonthlyBalance)

	if err := c.Bind(&entries); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	dayOne := "01"
	monthsFisrtDay := entries.Date[:8] + dayOne
	if err := db.DB.Table("monthly_balance").Where("year_month = ?", monthsFisrtDay).FirstOrCreate(&mBalance, models.MonthlyBalance{
		Amount:    0,
		YearMonth: monthsFisrtDay,
	}).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	entries.MonthlyBalanceID = mBalance.Id
	mBalance.Amount += entries.Amount

	if err := db.DB.Table("monthly_balance").Save(mBalance).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if err := db.DB.Table("entries").Create(&entries).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	c.JSON(http.StatusCreated, "Success")

	return nil
}

func (db *DBClient) ReadEntriesByID(c echo.Context) error {
	var transaction models.Entries
	transactionId, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.First(&transaction, transactionId).Error; err != nil {
		return c.JSON(http.StatusNotFound, err)
	}
	c.JSON(http.StatusOK, transaction)

	return nil
}

func (db *DBClient) UpdateEntries(c echo.Context) error {
	var transaction models.Entries
	transactionId, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.First(&transaction, transactionId).Error; err != nil {
		c.JSON(http.StatusNotFound, err)
	}

	if err := c.Bind(&transaction); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	if err := db.DB.Save(transaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, transaction)

	return nil
}

func (db *DBClient) DeleteEntries(c echo.Context) error {
	var transaction models.Entries
	transactionId, _ := strconv.Atoi(c.Param("id"))

	if err := db.DB.Delete(&transaction, transactionId).Error; err != nil {
		c.JSON(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, "Success: Deleted")

	return nil
}
