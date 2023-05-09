package orders

import (
	"strings"

	"github.com/luisnquin/blind-creator-rest-api-test/internal/models"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/models/relations"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/repository"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository { return Repository{db: db} }

// TODO: use JOIN instead

func (r Repository) PaginatedSearch(page, limit int) ([]relations.DetailedOrder, error) {
	offset := repository.CreateOffset(page, limit)

	var orders []models.Order

	err := r.db.Offset(offset).Limit(limit).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	detailedOrders := make([]relations.DetailedOrder, len(orders))

	for index, order := range orders {
		var customer models.Customer

		err := r.db.Table("customers").Where("customer_id = ?", order.CustomerID).First(&customer).Error
		if err != nil {
			return nil, err
		}

		var orderItems []models.OrderItem

		err = r.db.Table("order_items").Where("order_id = ?", order.OrderID).Find(&orderItems).Error
		if err != nil {
			return nil, err
		}

		items := make([]relations.OrderItem, len(orderItems))

		for index, item := range orderItems {
			var product models.Product

			err := r.db.Table("products").Where("product_id = ?", item.ProductID).First(&product).Error
			if err != nil {
				return nil, err
			}

			items[index] = relations.OrderItem{
				ID: item.OrderItemID,
				Product: relations.OrderProduct{
					ID:    product.ProductID,
					Name:  product.Name,
					Price: product.Price,
				},
				Quantity: item.Quantity,
			}
		}

		detailedOrders[index] = relations.DetailedOrder{
			ID:    order.OrderID,
			Items: items,
			Customer: relations.OrderCustomer{
				FirstName:   customer.FirstName,
				LastName:    customer.LastName,
				Email:       customer.Email,
				PhoneNumber: customer.PhoneNumber,
			},
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		}
	}

	return detailedOrders, nil
}

func (r Repository) GetByID(orderId string) (relations.DetailedOrder, error) {
	orderId = strings.TrimSpace(orderId)

	var order models.Order

	err := r.db.Table("orders").Where("order_id = ?", orderId).Find(&order).Error
	if err != nil {
		return relations.DetailedOrder{}, err
	}

	var customer models.Customer

	err = r.db.Table("customers").Where("customer_id = ?", order.CustomerID).First(&customer).Error
	if err != nil {
		return relations.DetailedOrder{}, err
	}

	var orderItems []models.OrderItem

	err = r.db.Table("order_items").Where("order_id = ?", order.OrderID).Find(&orderItems).Error
	if err != nil {
		return relations.DetailedOrder{}, err
	}

	items := make([]relations.OrderItem, len(orderItems))

	for index, item := range orderItems {
		var product models.Product

		err := r.db.Table("products").Where("product_id = ?", item.ProductID).First(&product).Error
		if err != nil {
			return relations.DetailedOrder{}, err
		}

		items[index] = relations.OrderItem{
			ID: item.OrderItemID,
			Product: relations.OrderProduct{
				ID:    product.ProductID,
				Name:  product.Name,
				Price: product.Price,
			},
			Quantity: item.Quantity,
		}
	}

	detailedOrder := relations.DetailedOrder{
		ID:    order.OrderID,
		Items: items,
		Customer: relations.OrderCustomer{
			FirstName:   customer.FirstName,
			LastName:    customer.LastName,
			Email:       customer.Email,
			PhoneNumber: customer.PhoneNumber,
		},
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}

	return detailedOrder, nil
}
