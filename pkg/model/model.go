package model

import (
	"errors"
	"math"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WeatherRequest struct definition
type WeatherRequest struct {
	Temperature float64   `bson:"temperature" json:"temperature"`
	Date        time.Time `bson:"date" json:"date"`
}

// Bind validates the request body
func (req *WeatherRequest) Bind(r *http.Request) error {
	if math.IsNaN(req.Temperature) {
		return errors.New("temperature is not a number")
	}

	if req.Date.IsZero() {
		req.Date = time.Now()
	}

	return nil
}

// WeatherResponse struct definition
type WeatherResponse struct {
	Temperature float32            `bson:"temperature" json:"temperature"`
	ID          primitive.ObjectID `bson:"_id, omitempty" json:"_id,omitempty"`
	Date        time.Time          `bson:"date" json:"date"`
}
