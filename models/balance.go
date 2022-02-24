package models

type MonthlyBalance struct {
	Id        uint   `json:"-" gorm:"column:id"`
	Amount    int    `json:"amount" gorm:"column:difference"`
	YearMonth string `json:"yearMonth" gorm:"column:year_month"`
	TotalId   int    `json:"-"`
}

type TotalBalance struct {
	Id     uint
	Amount int `gorm:"column:total"`
}
