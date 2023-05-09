package server

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/luisnquin/blind-creator-rest-api-test/internal/config"
)

type (
	Server struct {
		pathMethods map[string][]string
		config      config.App
		handlers    []handlerInfo
	}

	handlerInfo struct {
		handler    Handler
		method     string
		path       string
		withToken  bool
		withApiKey bool
	}

	Handler func(w http.ResponseWriter, r *http.Request, p Params)
)

type Registerer interface {
	RegisterHandler(path, method string, handler Handler, apiKey, token bool)
}

type Params = httprouter.Params

func New(config config.App) Server {
	return Server{
		pathMethods: make(map[string][]string),
		config:      config,
	}
}

func (s Server) Start() error {
	router := httprouter.New()

	for _, endpoint := range s.handlers {
		// handler := LogMiddleware(endpoint.handler)
		handler := endpoint.handler

		router.Handle(endpoint.method, endpoint.path, httprouter.Handle(handler))
	}

	return http.ListenAndServe(s.config.Server.Port(), router)
}
