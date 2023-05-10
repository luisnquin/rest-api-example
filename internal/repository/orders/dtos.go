package orders

import "time"

type orderAndCustomer struct {
	CustomerId  string    `gorm:"customer_id"`
	OrderID     string    `gorm:"order_id"`
	Email       string    `gorm:"email"`
	FirstName   string    `gorm:"first_name"`
	LastName    string    `gorm:"last_name"`
	PhoneNumber string    `gorm:"phone_number"`
	CreatedAt   time.Time `gorm:"created_at"`
	UpdatedAt   time.Time `gorm:"updated_at"`
}
