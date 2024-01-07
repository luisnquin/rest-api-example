package locations

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/api"
	"github.com/luisnquin/server-example/internal/database/sqlc"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/luisnquin/server-example/internal/server"
)

func (m Manager) GetCitiesHandler() server.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p server.Params) {
		params := parseGetCitiesParams(r)

		var (
			cities any
			err    error
		)

		if params.CountryCode != "" {
			cities, err = m.querier.GetCitiesByCountry(r.Context(), sqlc.GetCitiesByCountryParams{
				CountryCode: params.CountryCode,
				Limit:       params.Count,
			})
		} else {
			cities, err = m.querier.GetCities(r.Context(), params.Count)
		}
		if err != nil {
			log.Err(err).Msg("error trying to get cities")
			api.SendInternalServerError(w)

			return
		}

		api.Response(w, http.StatusOK, api.StdResponse{
			Message: api.Success,
			Data:    cities,
		})
	}
}
