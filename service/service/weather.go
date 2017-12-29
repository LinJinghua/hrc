package service

import (
	"github.com/unrolled/render"
	"fmt"
	"net/http"
	"time"
)

// Weather is a middleware handler that provides weather forecast for a given destination.
type Weather struct {
}

// NewWeather returns a new Weather instance
func NewWeather() *Weather {
	return &Weather{}
}

func (l *Weather) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// fmt.Println("Start Weather service:", time.Now())
	start := time.Now()
	for i := DestinationsNumber; i > 0; i-- {
		GetWeather(nil)
	}
	fmt.Println("Finish Weather service:", time.Since(start))

	next(rw, r)
}

// GetWeather .
func GetWeather(d *Destinations) *Weather {
	time.Sleep(WeatherServiceTime * time.Millisecond)
	return NewWeather()
}

// WeatherHandler .
func WeatherHandler(formatter *render.Render) http.HandlerFunc {
	count := 0
	return func(w http.ResponseWriter, req *http.Request) {
		count++
		c := count
		start := time.Now()
		GetWeather(nil)

		formatter.JSON(w, http.StatusOK, struct {
			ID       int `json:"weather"`
			Duration string `json:"duration"`
		}{ID: c, Duration: time.Since(start).String()})
	}
}
