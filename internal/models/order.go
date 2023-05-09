package models

import (
	"database/sql"
	"time"
)

type Order struct {
	OrderID    string         `gorm:"primaryKey"`
	CustomerID sql.NullString `gorm:"foreignKey"`
	CreatedAt  time.Time      `gorm:"not null"`
	UpdatedAt  time.Time      `gorm:"not null"`
}
