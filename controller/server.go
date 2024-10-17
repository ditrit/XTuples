package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/supertokens/supertokens-golang/supertokens"

	"fmt"
	"go-http/pkg/app"
)

type server struct {
	app *app.App
}

func NewServer(app *app.App) *server {
	return &server{app}
}

// Start Starts the server on the specified host and port + with the defined routes
func (s *server) Start() {
	conf := s.app.Conf()
	backend := chi.NewRouter()

	backend.Use(
		middleware.Logger,
		middleware.Recoverer,
		middleware.AllowContentType("application/json"),
		middleware.ContentCharset("UTF-8", "Latin-1", ""),
		cors.Handler(cors.Options{
			AllowedOrigins: []string{conf.Backend.Host},
			AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowedHeaders: append([]string{"content-type"},
				supertokens.GetAllCORSHeaders()...),
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	backend.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			supertokens.Middleware(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
				next.ServeHTTP(rw, req)
			})).ServeHTTP(w, r)
		})
	})

	s.router(backend)

	address := fmt.Sprintf("%v:%v", conf.Backend.Host, conf.Backend.Port)
	http.ListenAndServe(address, backend)
}
