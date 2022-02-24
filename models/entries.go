package models

type Entries struct {
	Id               uint   `json:"-" gorm:"column:id"`
	MonthlyBalanceID uint   `json:"balance_id" gorm:"column:balance_id"`
	Title            string `json:"title" gorm:"column:title"`
	Type             string `json:"type" gorm:"column:type"`
	Amount           int    `json:"amount" gorm:"column:amount"`
	Installments     int    `json:"installments" gorm:"column:installments"`
	IsPlanned        bool   `json:"isPlanned" gorm:"column:is_planned"`
	Date             string `json:"date" gorm:"column:date"`
}
