package handlers

import (
	"github.com/go-chi/chi"
)

func (s *Server) setupEndpoints(app *chi.Mux) {
	app.Route("/api", func(apiRoutes chi.Router) {
		apiRoutes.Route("/users", func(userRoutes chi.Router) {
		})
	})
}
