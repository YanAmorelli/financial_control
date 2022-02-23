package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	InOut        string `json:"inOut"`
	Name         string `json:"name"`
	InOutType    string `json:"inOutType"`
	Value        int    `json:"value"`
	Installments int    `json:"installments"`
	IsPlanned    bool   `json:"isPlanned"`
	Date         string `json:"date"`
}
