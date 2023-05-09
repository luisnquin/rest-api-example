package models

type OrderItem struct {
	OrderItemID string `gorm:"primaryKey"`
	OrderID     string `gorm:"foreignKey"`
	ProductID   string `gorm:"foreignKey"`
	Quantity    uint32 `gorm:"not null"`
}
