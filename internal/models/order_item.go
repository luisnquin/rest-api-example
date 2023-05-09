package models

import "github.com/shopspring/decimal"

type OrderItem struct {
	OrderItemID string          `gorm:"primaryKey"`
	OrderID     string          `gorm:"foreignKey"`
	ProductID   string          `gorm:"foreignKey"`
	Quantity    uint32          `gorm:"not null"`
	Price       decimal.Decimal `gorm:"type:decimal(7,6);"`
}
