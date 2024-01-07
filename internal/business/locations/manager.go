package locations

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/database/sqlc"
	"github.com/luisnquin/server-example/internal/server"
)

type (
	Manager struct {
		// repository repository
		querier sqlc.Querier
	}

	// repository struct {
	// 	orders orders.Repository
	// }
)

func NewManager(querier sqlc.Querier) *Manager {
	return &Manager{
		querier: querier,
		// repository{orders: orders.NewRepository(db)},
	}
}

func (m Manager) RegisterHandlers(srv server.Registerer) {
	srv.RegisterHandler("/locations/v1/cities", http.MethodGet, m.GetCitiesHandler(), false, false)
}
