package client

import (
	"io/ioutil"
	"fmt"
	"time"
	"net/http"

	"github.com/LinJinghua/hrc/service/service"
)

// Client .
type Client struct {
}

// NewClient returns a new Client instance
func NewClient() *Client {
	return &Client{}
}

// GetService .
func (c *Client) GetService() {
	start := time.Now()
	fmt.Println("[Sync Service]  |   Start  |", start)

	customer := GetCustomers()
	dests := GetDest(customer)
	for _, dest := range dests {
		GetWeather(&dest)
	}
	for _, dest := range dests {
		GetQuoting(&dest)
	}
	
	fmt.Println("[Sync Service]  |  Finish  |", time.Now())
	fmt.Println("[Sync Service]  | Duration |", time.Since(start))
}

// GetServiceAsync .
func (c *Client) GetServiceAsync() {
	start := time.Now()
	fmt.Println("[Async Service] |   Start  |", start)

	custCh := make(chan *service.Customers, 1)
	destsCh := make(chan []service.Destinations, 1)
	weaCh := make(chan *service.Weather, 1)
	quotCh := make(chan *service.Quoting, 1)


	go func ()  {
		custCh <- GetCustomers()
	}()

	go func ()  {
		destsCh <- GetDest(<- custCh)
	}()

	dests := <- destsCh

	go func ()  {	
		for i := len(dests) - 1; i >= 0; i-- {
			go func (dest *service.Destinations)  {
				weaCh <- GetWeather(dest)
			}(&dests[i])
			go func (dest *service.Destinations)  {
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

	fmt.Println("[Async Service] |  Finish  |", time.Now())
	fmt.Println("[Async Service] | Duration |", time.Since(start))
}

// GetCustomers .
func GetCustomers() *service.Customers {
	httpGet("http://localhost:8080/customers")
	return nil
}

// GetDest .
func GetDest(c *service.Customers) []service.Destinations {
	httpGet("http://localhost:8080/destinations")
	return make([]service.Destinations, service.DestinationsNumber)
}

// GetWeather .
func GetWeather(d *service.Destinations) *service.Weather {
	httpGet("http://localhost:8080/weather")
	return nil
}

// GetQuoting .
func GetQuoting(d *service.Destinations) *service.Quoting {
	httpGet("http://localhost:8080/quoting")
	return nil
}

func httpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, er := ioutil.ReadAll(resp.Body)
	if er != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
