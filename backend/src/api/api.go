package api

import (
	"net/http"

	"github.com/ArtemII9a/portfolio/src/errors"
	"github.com/ArtemII9a/portfolio/src/handlers"
	"github.com/ArtemII9a/portfolio/src/middleware"
)

type API struct {
	version string
	port    string
}

func NewAPIServer(version, port string) *API {
	return &API{version: version, port: port}
}

func (api *API) Run() error {
	defer errors.RecoverFromPanic()

	mux := http.NewServeMux()
	ApplyMuxMiddleware := middleware.NewMiddlewareStack(
		middleware.Logging,
	)
	middlewaredMux := ApplyMuxMiddleware(mux)

	handler := handlers.New()
	mux.Handle("GET /", handler)

	server := &http.Server{
		Addr:    api.port,
		Handler: middlewaredMux,
	}

	return server.ListenAndServe()
}
