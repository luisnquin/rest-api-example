package relations

import (
	"time"

	"github.com/shopspring/decimal"
)

type (
	DetailedOrder struct {
		ID        string        `json:"id"`
		Items     []OrderItem   `json:"items"`
		Customer  OrderCustomer `json:"customer"`
		CreatedAt time.Time     `json:"createdAt"`
		UpdatedAt time.Time     `json:"updatedAt"`
	}

	OrderItem struct {
		ID       string       `json:"id"`
		Product  OrderProduct `json:"product"`
		Quantity uint32       `json:"quantity"`
	}

	OrderCustomer struct {
		FirstName   string `json:"firstName"`
		LastName    string `json:"lastName"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phoneNumber"`
	}

	OrderProduct struct {
		ID    string          `json:"id"`
		Name  string          `json:"name"`
		Price decimal.Decimal `json:"price"`
	}
)

/*
Desired output

[
		{
			"id": "",
			"items": [
				{
					"id": "-",
					"product": {
						"id": "-"
						"name": "-"
						"price": 0.00
					}
					"quantity": 0
				}
			],
			"customer": {
				"firstName": "-",
				"lastName": "-",
				"email": "-",
				"phoneNumber": ""
			},
			"createdAt": "",
			"updatedAt": ""
		}
	]
*/
