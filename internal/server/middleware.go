package server

import (
	"net/http"

	"github.com/luisnquin/server-example/internal/log"
)

func (s Server) logIncomingRequestsMiddleware(next HandlerFunc) HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request, p Params) {
		log.Info().
			Str("user-agent", r.UserAgent()).Int64("length", r.ContentLength).
			Str("method", r.Method).Str("path", r.URL.Path).Send()

		next(w, r, p)
	}
}
