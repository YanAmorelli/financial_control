package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Money struct {
	Id           int
	InOut        string `json:"inOut"`
	Name         string `json:"name"`
	InOutType    string `json:"inOutType"`
	Value        int    `json:"value"`
	Installments int    `json:"installments"`
	IsPlanned    bool   `json:"isPlanned"`
	Date         string `json:"date"`
}

type Balance struct {
	Id         int
	Difference int
	YearMonth  string
}

func PostInOut(c echo.Context) error {
	m := new(Money)
	if err := c.Bind(&m); err != nil {
		return c.JSON(http.StatusCreated, err)
	}

	return c.JSON(http.StatusCreated, m)
}

func main() {
	e := echo.New()

	e.POST("/money", PostInOut)

	e.Logger.Fatal(e.Start(":8080"))
}
