package server

import (
	"log"
	"net/http"

	"github.com/diskordanz/consumer/service"

	"github.com/gorilla/mux"
)

type ConsumerAPI struct {
	router *mux.Router
	ss     service.Service
}

func NewAPI(ss *service.Service) *ConsumerAPI {
	return &ConsumerAPI{router: mux.NewRouter(), ss: *ss}
}

func (api *ConsumerAPI) AssignRoutes() {
	// to get franchise list -> api/franchises?count=<number>&offset=<number> | count and offset are used for pagination
	api.router.HandleFunc("/api/franchises", api.ListFranchises).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	//api.router.HandleFunc("/api/franchises", api.GetFranchisesByName).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}", "name", "{name}")

	api.router.HandleFunc("/api/franchises/{id}/locations", api.GetLocationsOfFranchise).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	//api.router.HandleFunc("/api/franchises/{id}/locations", api.GetLocationsOfFranchiseByName).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}", "name", "{name}")

	api.router.HandleFunc("/api/categories", api.ListCategories).Methods("GET")

	api.router.HandleFunc("/api/products", api.ListProducts).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	api.router.HandleFunc("/api/products/{id}", api.GetProduct).Methods("GET")

	api.router.HandleFunc("/api/consumers/{id}/profile", api.GetProfile).Methods("GET")

	//api.router.HandleFunc("/api/consumers/{id}/cart", api.GetCart).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")

	api.router.HandleFunc("/api/consumers/{id}/orders", api.ListOrders).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	api.router.HandleFunc("/api/consumers/{consumer_id}/orders/{order_id}", api.GetOrder).Methods("GET")

}

func (api *ConsumerAPI) Run(host string) {
	log.Printf("started listening on %s", host)
	log.Fatal(http.ListenAndServe(host, api.router))
}
