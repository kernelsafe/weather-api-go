package router

import (
	"net/http"

	"github.com/kernelsafe/weather-api-go/pkg/db"
	"github.com/kernelsafe/weather-api-go/pkg/model"
	"github.com/kernelsafe/weather-api-go/pkg/service"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const collectionName = "weather"

// WeatherRoutes returns all Weather's endpoints
func WeatherRoutes(dbClient *mongo.Client) *chi.Mux {
	router := chi.NewRouter()

	router.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "id field is required!", http.StatusBadRequest)
			return
		}
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "invalid id!", http.StatusBadRequest)
			return
		}
		res, err := service.GetOne(dbClient, collectionName, objID)

		if err != nil || res == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		render.JSON(w, r, res)
	})

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := service.GetAll(dbClient, collectionName)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, res)
	})

	router.Post("/", func(w http.ResponseWriter, r *http.Request) {
		weather := &model.WeatherRequest{}
		if err := render.Bind(r, weather); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		res, err := db.AddDocument(dbClient, collectionName, weather)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		res, err = service.GetOne(dbClient, collectionName, res.(primitive.ObjectID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		render.JSON(w, r, res)
	})

	router.Delete("/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		if id == "" {
			http.Error(w, "id field is required!", http.StatusBadRequest)
			return
		}
		objID, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			http.Error(w, "invalid id!", http.StatusBadRequest)
			return
		}

		res, err := service.GetOne(dbClient, collectionName, objID)

		if err != nil || res == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		err = service.DeleteOne(dbClient, collectionName, objID)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		render.Status(r, http.StatusCreated)
		render.DefaultResponder(w, r, res)
	})

	return router
}
