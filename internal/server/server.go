package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/luisnquin/server-example/internal/config"
	"github.com/luisnquin/server-example/internal/log"
	"github.com/samber/lo"
)

type (
	Server struct {
		OnBeforeStart func()
		pathMethods   map[string][]string
		config        config.App
		handlers      []handlerInfo
	}

	handlerInfo struct {
		handler    HandlerFunc
		method     string
		path       string
		withToken  bool
		withApiKey bool
	}

	HandlerFunc func(w http.ResponseWriter, r *http.Request, p Params)
)

type Registerer interface {
	RegisterHandler(path, method string, handler HandlerFunc, apiKey, token bool)
}

type registerer interface {
	RegisterHandlers(srv Registerer)
}

type Params = httprouter.Params

func New(config config.App) *Server {
	return &Server{
		pathMethods: make(map[string][]string),
		config:      config,
	}
}

func (s Server) Start() error {
	router := httprouter.New()

	for _, endpoint := range s.handlers {
		// handler := LogMiddleware(endpoint.handler)
		handler := s.logIncomingRequestsMiddleware(endpoint.handler)

		log.Debug().Msgf("HTTP Route %s %s has been registered ", endpoint.method, endpoint.path)

		router.Handle(endpoint.method, endpoint.path, httprouter.Handle(handler))
	}

	if s.OnBeforeStart != nil {
		go s.OnBeforeStart()
	}

	return http.ListenAndServe(s.config.Server.Port(), router)
}

func (s *Server) RegisterBatch(rs ...registerer) {
	lo.ForEach(rs, func(r registerer, _ int) { r.RegisterHandlers(s) })
}

func (s *Server) RegisterHandler(path, method string, handler HandlerFunc, apiKey, token bool) {
	s.validateHandlerRequest(path, method, handler)

	s.pathMethods[path] = append(s.pathMethods[path], method)

	s.handlers = append(s.handlers, handlerInfo{
		handler:    handler,
		withApiKey: apiKey,
		method:     method,
		withToken:  token,
		path:       path,
	})
}

func (s Server) validateHandlerRequest(path, method string, handler HandlerFunc) {
	switch {
	case path == "":
		panic("no path provided")
	case strings.ToLower(path) != path:
		panic(fmt.Sprintf("path '%s' must be in lowercase", path))
	case path[0] != '/':
		panic(fmt.Sprintf("the path '%s' must start with a slash", path))
	case !strings.Contains("abcdefghijklmnopqrstuvwxyz", string(path[len(path)-1])):
		panic(fmt.Sprintf("the path '%s' must end with a valid letter", path))
	}

	validHttpMethods := []string{http.MethodGet, http.MethodPost, http.MethodPatch, http.MethodPut, http.MethodDelete}

	if !lo.Contains(validHttpMethods, method) {
		panic(fmt.Sprintf("http method '%s' not supported for path %s", method, path))
	}

	if handler == nil {
		panic(fmt.Sprintf("http handler for path and method %s - %s is nil", path, method))
	}

	if lo.Contains(s.pathMethods[path], method) {
		panic(fmt.Sprintf("http method '%s' with path '%s' already used", method, path))
	}
}
