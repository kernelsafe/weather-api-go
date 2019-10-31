package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/kernelsafe/weather-api/pkg/db"
	"github.com/kernelsafe/weather-api/pkg/router"
	"github.com/kernelsafe/weather-api/pkg/util"
)

// Routes returns all routes
func Routes() *chi.Mux {

	root := chi.NewRouter()
	dbClient := db.GetClient()
	cors := cors.New(cors.Options{
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		AllowCredentials: true,
	})

	root.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
		cors.Handler,
	)

	root.Route("/api", func(r chi.Router) {
		r.Mount("/v1/weather", router.WeatherRoutes(dbClient))
	})

	return root
}

func main() {
	router := Routes()
	log.Fatal(http.ListenAndServe(":"+util.GetEnv("PORT", "3000"), router))
}
