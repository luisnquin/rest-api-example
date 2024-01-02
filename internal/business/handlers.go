package business

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/luisnquin/server-example/internal/api"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
	"gorm.io/gorm"
)

func (m Manager) PaginatedOrdersHandler() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		page, limit := 1, 20

		query := r.URL.Query()

		if p := query.Get("page"); p != "" {
			p, err := strconv.Atoi(p)
			if err == nil && p > 0 {
				page = p
			}
		}

		if l := query.Get("limit"); l != "" {
			l, err := strconv.Atoi(l)
			if err == nil && l > 0 {
				limit = l
			}
		}

		orders, err := m.repository.orders.PaginatedSearch(page, limit)
		if err != nil {
			log.Err(err).Msg("error trying to retrieve paginated and detailed orders")
			api.SendInternalServerError(w)

			return
		}

		api.Response(w, http.StatusOK, api.StdResponse{
			Message: api.Success,
			Data:    orders,
		})
	}
}

func (m Manager) GetOneOrderHandler() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		orderId := p.ByName("id")

		order, err := m.repository.orders.GetByID(orderId)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				log.Debug().Msgf("order %q not found in db", orderId)
				api.SendNotFound(w, fmt.Sprintf("order '%s' not found", orderId))
			} else {
				log.Err(err).Msgf("unable to get order %q from db", orderId)
				api.SendInternalServerError(w)
			}

			return
		}

		api.Response(w, http.StatusOK, api.StdResponse{
			Message: api.Success,
			Data:    order,
		})
	}
}
