package main

import (
	"embed"
	"io/fs"
	"log"

	"net/http"

	"github.com/go-chi/chi/v5"

	"go-http/internal/api/v1/modules/cron_module"
	"go-http/internal/api/v1/modules/exec_module"
)

//go:embed web/dist
var efs embed.FS

func (s *server) router(app *chi.Mux) {
	api := chi.NewRouter()
	apiPath := s.app.Conf().Backend.APIPath
	app.Mount(apiPath, api)

	cron_module.Router(api)
	exec_module.Router(api)

	dist, err := fs.Sub(efs, "web/dist")
	if err != nil {
		log.Fatalf("dist file server error: %v", err)
	}

	app.Handle("/*", http.StripPrefix("/", http.FileServer(http.FS(dist))))
}
