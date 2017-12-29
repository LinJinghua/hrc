package service

import (
	"fmt"
	"net/http"
	"time"
)

// TravelAgency is a middleware handler that provides the travel agency.
type TravelAgency struct {
}

// NewTravelAgency returns a new TravelAgency instance
func NewTravelAgency() *TravelAgency {
	return &TravelAgency{}
}

func (t *TravelAgency) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	next(rw, r)

	r.ParseForm()
	fmt.Println(r.Form, r.Form.Get("m"))
	if r.Form.Get("m") == "s" {
		start := time.Now()
		fmt.Println("[Sync Service]  |   Start  |", start)
		fmt.Fprintln(rw, "[Sync Service]  |   Start  |", start)
		t.GetTravelAgency()
		fmt.Fprintln(rw, "[Sync Service]  |  Finish  |", time.Now())
		fmt.Fprintln(rw, "[Sync Service]  | Duration |", time.Since(start))
	} else {
		start := time.Now()
		fmt.Println("[ASync Service]  |   Start  |", start)
		fmt.Fprintln(rw, "[Async Service] |   Start  |", start)
		t.GetTravelAgencyAsync()
		fmt.Fprintln(rw, "[Async Service] |  Finish  |", time.Now())
		fmt.Fprintln(rw, "[Async Service] | Duration |", time.Since(start))
	}

}

// GetTravelAgency .
func (t *TravelAgency) GetTravelAgency() {

	customer := GetCustomers()
	dests := GetDest(customer)
	for _, dest := range dests {
		GetWeather(&dest)
	}
	for _, dest := range dests {
		GetQuoting(&dest)
	}
	
}

// GetTravelAgencyAsync .
func (t *TravelAgency) GetTravelAgencyAsync() {

	custCh := make(chan *Customers, 1)
	destsCh := make(chan []Destinations, 1)
	weaCh := make(chan *Weather, 1)
	quotCh := make(chan *Quoting, 1)


	go func ()  {
		custCh <- GetCustomers()
	}()

	go func ()  {
		destsCh <- GetDest(<- custCh)
	}()

	dests := <- destsCh

	go func ()  {	
		for i := len(dests) - 1; i >= 0; i-- {
			go func (dest *Destinations)  {
				weaCh <- GetWeather(dest)
			}(&dests[i])
			go func (dest *Destinations)  {
				quotCh <- GetQuoting(dest)
			}(&dests[i])
		}
	}()

	for i := len(dests) << 1; i > 0; i-- {
		select {
		case <- weaCh:
		case <- quotCh:
		}
	}

}
