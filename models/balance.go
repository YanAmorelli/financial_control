package models

import (
	"gorm.io/gorm"
)

// import "gorm.io/gorm"

type MonthlyBalance struct {
	Id        int    `json:"-" gorm:"column:id"`
	Amount    int    `json:"amount" gorm:"column:amount"`
	YearMonth string `json:"yearMonth" gorm:"column:year_month"`
	// TotalId   int    `json:"-"`
}

// type TotalBalance struct {
// 	Id     uint
// 	Amount int `gorm:"column:total"`
// }

func UpdateBalances(db *gorm.DB, mBalance *MonthlyBalance, entries *Entries) error {
	if err := db.Table("monthly_balance").First(&mBalance, entries.MonthlyBalanceID).Error; err != nil {
		return err
	}

	if mBalance.Id == entries.MonthlyBalanceID {
		mBalance.Amount += entries.Amount
	}

	if err := db.Table("monthly_balance").Save(mBalance).Error; err != nil {
		return err
	}

	return nil
}
