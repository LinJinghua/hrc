package service

import (
	"github.com/unrolled/render"
	"fmt"
	"net/http"
	"time"
)

// Destinations is a middleware handler that provides a list of visited and recommended destinations for an authenticated customer.
type Destinations struct {
}

// NewDestinations returns a new Destinations instance
func NewDestinations() *Destinations {
	return &Destinations{}
}

func (l *Destinations) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// fmt.Println("Start Destinations service:", time.Now())
	start := time.Now()
	GetDest(nil)
	fmt.Println("Finish Destinations service:", time.Since(start))

	next(rw, r)
}

// GetDest .
func GetDest(c *Customers) []Destinations {
	time.Sleep(DestinationsServiceTime * time.Millisecond)
	return make([]Destinations, DestinationsNumber)
}

// DestinationsHandler .
func DestinationsHandler(formatter *render.Render) http.HandlerFunc {
	
	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		GetDest(nil)

		formatter.JSON(w, http.StatusOK, struct {
			ID       string `json:"destinations"`
			Duration string `json:"duration"`
		}{ID: "Destinations", Duration: time.Since(start).String()})
	}
}
