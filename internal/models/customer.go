package models

type Customer struct {
	CustomerID  string `gorm:"primaryKey"`
	FirstName   string `gorm:"not null"`
	LastName    string `gorm:"not null"`
	Email       string `gorm:"not null;unique"`
	PhoneNumber string `gorm:"not null"`
}
