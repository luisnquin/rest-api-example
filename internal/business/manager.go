package business

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/repository/orders"
	"github.com/luisnquin/server-example/internal/server"
	"gorm.io/gorm"
)

type (
	Manager struct {
		repository repository
	}

	repository struct {
		orders orders.Repository
	}
)

func NewManager(db *gorm.DB) Manager {
	return Manager{
		repository{orders: orders.NewRepository(db)},
	}
}

func (m Manager) RegisterHandlers(srv server.Registerer) {
	srv.RegisterHandler("/orders", http.MethodGet, m.PaginatedOrdersHandler(), false, false)
	srv.RegisterHandler("/orders/:id", http.MethodGet, m.GetOneOrderHandler(), false, false)
}
