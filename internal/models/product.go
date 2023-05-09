package models

import (
	"github.com/shopspring/decimal"
)

type Product struct {
	ProductID string          `gorm:"primaryKey"`
	Name      string          `gorm:"not null"`
	Price     decimal.Decimal `gorm:"type:decimal(10,6);"`
}
