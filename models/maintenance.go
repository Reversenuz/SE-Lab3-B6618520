package models

import "gorm.io/gorm"

type Maintenance struct {
	gorm.Model
	Description string  `json:"description"`
	Date        string  `json:"date"`
	Cost        float64 `json:"cost"`
}
