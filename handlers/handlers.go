package handlers

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func setupMiddleware(app *chi.Mux) {
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Compress(6, "application/json"))
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)
	app.Use(middleware.URLFormat)
	app.Use(middleware.Timeout(60 * time.Second))
}
func SetupRouter() *chi.Mux {
	server := NewServer()
	app := chi.NewRouter()
	setupMiddleware(app)
	server.setupEndpoints(app)
	return app
}
