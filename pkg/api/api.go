package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-pg/pg/v10"
)

func StartAPI(pgdb *pg.DB) *chi.Mux {
	// get the router
	r := chi.NewRouter()

	// add midleware
	// in this case we will store our DB to use it later
	r.Use(middleware.Logger, middleware.WithValue("DB", pgdb))

	r.Route("/comments", func(r chi.Router) {
		r.Get("/", getComments)
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("up and running"))
	})

	return r
}

func getComments(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("comments"))
}
