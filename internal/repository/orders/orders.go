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

func (r Repository) PaginatedSearch(page, limit int) ([]relations.DetailedOrder, error) {
	offset := repository.CreateOffset(page, limit)

	var ordersWithCustomer []orderAndCustomer

	err := r.db.Table("orders").Select("orders.order_id", "c.customer_id", "c.email",
		"c.first_name", "c.last_name", "c.phone_number", "orders.created_at", "orders.updated_at").
		Joins("LEFT JOIN customers AS c ON orders.customer_id = c.customer_id").
		Offset(offset).Limit(limit).Find(&ordersWithCustomer).Error
	if err != nil {
		return nil, err
	}

	detailedOrders := make([]relations.DetailedOrder, len(ordersWithCustomer))

	for i, oc := range ordersWithCustomer {
		detailedOrder, err := r.getDetailedOrder(oc)
		if err != nil {
			return nil, err
		}

		detailedOrders[i] = detailedOrder
	}

	return detailedOrders, nil
}

func (r Repository) GetByID(orderId string) (relations.DetailedOrder, error) {
	orderId = strings.TrimSpace(orderId)

	var orderAndCustomer orderAndCustomer

	err := r.db.Table("orders").Select("orders.order_id", "c.customer_id", "c.email",
		"c.first_name", "c.last_name", "c.phone_number", "orders.created_at", "orders.updated_at").
		Joins("LEFT JOIN customers AS c ON orders.customer_id = c.customer_id").
		Where("orders.order_id = ?", orderId).First(&orderAndCustomer).Error
	if err != nil {
		return relations.DetailedOrder{}, err
	}

	return r.getDetailedOrder(orderAndCustomer)
}

func (r Repository) getDetailedOrder(oc orderAndCustomer) (relations.DetailedOrder, error) {
	var orderItems []models.OrderItem

	err := r.db.Where("order_items.order_id = ?", oc.OrderID).Find(&orderItems).Error
	if err != nil {
		return relations.DetailedOrder{}, err
	}

	items := make([]relations.OrderItem, len(orderItems))

	for j, item := range orderItems {
		var product models.Product

		err = r.db.Where("products.product_id = ?", item.ProductID).First(&product).Error
		if err != nil {
			return relations.DetailedOrder{}, err
		}

		items[j] = relations.OrderItem{
			ID: item.OrderItemID,
			Product: relations.OrderProduct{
				ID:    product.ProductID,
				Name:  product.Name,
				Price: product.Price,
			},
			Quantity: item.Quantity,
		}
	}

	return relations.DetailedOrder{
		ID:    oc.OrderID,
		Items: items,
		Customer: relations.OrderCustomer{
			FirstName:   oc.FirstName,
			LastName:    oc.LastName,
			Email:       oc.Email,
			PhoneNumber: oc.PhoneNumber,
		},
		CreatedAt: oc.CreatedAt,
		UpdatedAt: oc.UpdatedAt,
	}, nil
}
