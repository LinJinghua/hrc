package service

import (
	"github.com/unrolled/render"
	"fmt"
	"net/http"
	"time"
)

// Customers is a middleware handler that provides information about customers of the travel agency.
type Customers struct {
}

// NewCustomers returns a new Customers instance
func NewCustomers() *Customers {
	return &Customers{}
}

func (l *Customers) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

	// fmt.Println("Start Customers service:", time.Now())
	start := time.Now()
	GetCustomers()
	fmt.Println("Finish Customers service:", time.Since(start))

	next(rw, r)
}

// GetCustomers .
func GetCustomers() *Customers {
	time.Sleep(CustomersServiceTime * time.Millisecond)
	return NewCustomers()
}

// CustomersHandler .
func CustomersHandler(formatter *render.Render) http.HandlerFunc {

	return func(w http.ResponseWriter, req *http.Request) {
		start := time.Now()
		GetCustomers()

		formatter.JSON(w, http.StatusOK, struct {
			ID       string `json:"customers"`
			Duration string `json:"duration"`
		}{ID: "8675309", Duration: time.Since(start).String()})
	}
}
