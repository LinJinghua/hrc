package service

import (
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"

	"github.com/LinJinghua/hrc/service/service"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	// n.Use(service.NewTravelAgency())
	// n := negroni.New(negroni.NewRecovery(), negroni.NewLogger(), service.NewCustomers(), service.NewDestinations(), service.NewWeather(), service.NewQuoting())
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			//fmt.Println(root)
		}
	}

	mx.NotFoundHandler = NotImplementedHandler()
	mx.MethodNotAllowedHandler = NotImplementedHandler()
	mx.HandleFunc("/login", loginHandler(formatter))
	mx.HandleFunc("/customers", service.CustomersHandler(formatter)).Methods("GET")
	mx.HandleFunc("/destinations", service.DestinationsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/weather", service.WeatherHandler(formatter)).Methods("GET")
	mx.HandleFunc("/quoting", service.QuotingHandler(formatter)).Methods("GET")
	mx.HandleFunc("/", homeHandler(formatter)).Methods("GET")
	mx.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(webRoot+"/assets/"))))
	// mx.PathPrefix("/").Handler(http.FileServer(http.Dir(webRoot + "/assets/")))

}
