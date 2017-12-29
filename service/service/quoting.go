package service

import (
	"github.com/unrolled/render"
	"fmt"
	"net/http"
	"time"
)

// Quoting is a middleware handler that provides price calculation for a customer to travel to a recommended destination.
type Quoting struct {
}

// NewQuoting returns a new Quoting instance
func NewQuoting() *Quoting {
	return &Quoting{}
}

func (l *Quoting) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// fmt.Println("Start Quoting service:", time.Now())
	start := time.Now()
	for i := DestinationsNumber; i > 0; i-- {
		GetQuoting(nil)
	}
	fmt.Println("Finish Quoting service:", time.Since(start))

	next(rw, r)
}

// GetQuoting .
func GetQuoting(d *Destinations) *Quoting {
	time.Sleep(QuotingServiceTime * time.Millisecond)
	return NewQuoting()
}

// QuotingHandler .
func QuotingHandler(formatter *render.Render) http.HandlerFunc {
	count := 0
	return func(w http.ResponseWriter, req *http.Request) {
		count++
		c := count
		start := time.Now()
		GetQuoting(nil)

		formatter.JSON(w, http.StatusOK, struct {
			ID       int `json:"quoting"`
			Duration string `json:"duration"`
		}{ID: c, Duration: time.Since(start).String()})
	}
}
