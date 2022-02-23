package models

import "gorm.io/gorm"

type Balance struct {
	gorm.Model
	Difference int
	YearMonth  string
}
