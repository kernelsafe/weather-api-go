package main

import (
	"compress/flate"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/kernelsafe/weather-api-go/pkg/db"
	"github.com/kernelsafe/weather-api-go/pkg/router"
	"github.com/kernelsafe/weather-api-go/pkg/util"
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
		middleware.Compress(flate.DefaultCompression),
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
