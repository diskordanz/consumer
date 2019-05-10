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
	//franchises
	api.router.HandleFunc("/api/franchises", api.ListFranchises).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	api.router.HandleFunc("/api/franchises/{id}", api.GetFranchise).Methods("GET")
	//locations
	api.router.HandleFunc("/api/franchises/{id}/locations", api.GetLocationsOfFranchise).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	//categories
	api.router.HandleFunc("/api/categories", api.ListCategories).Methods("GET")
	//products
	api.router.HandleFunc("/api/products", api.ListProducts).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}", "name", "{name}")
	api.router.HandleFunc("/api/categories/{id}/products", api.ListProductsByCategory).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}", "name", "{name}")
	api.router.HandleFunc("/api/products/{id}", api.GetProduct).Methods("GET")
	//consumers
	api.router.HandleFunc("/api/consumers/{id}/cart", api.GetCart).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	api.router.HandleFunc("/api/consumers/{id}/cart", api.CreateCartItem).Methods("POST")
	api.router.HandleFunc("/api/consumers/{id}/cart", api.CORS).Methods("OPTIONS")
	api.router.HandleFunc("/api/consumers/{id}/cart/{item_id}", api.UpdateCartItem).Methods("PUT")
	api.router.HandleFunc("/api/consumers/{id}/cart/{item_id}", api.CORS).Methods("OPTIONS")
	api.router.HandleFunc("/api/consumers/{id}/cart", api.GetCartItem).Methods("GET").Queries("product_id", "{product_id}")

	api.router.HandleFunc("/api/consumers/{id}", api.GetConsumer).Methods("GET")
	api.router.HandleFunc("/api/consumers", api.CreateConsumer).Methods("POST")
	api.router.HandleFunc("/api/consumers", api.CORS).Methods("OPTIONS")

	api.router.HandleFunc("/api/consumers/{id}", api.UpdateConsumer).Methods("PUT")
	//orders
	api.router.HandleFunc("/api/consumers/{id}/orders", api.ListOrders).Methods("GET").Queries("count", "{count:[0-9]+}", "offset", "{offset:[0-9]+}")
	api.router.HandleFunc("/api/consumers/{consumer_id}/orders/{order_id}", api.GetOrder).Methods("GET")
	//healthcheck
	api.router.HandleFunc("/api/healthcheck", api.Healthcheck).Methods("GET")
}

func (api *ConsumerAPI) Run(host string) {
	log.Printf("started listening on %s", host)
	log.Fatal(http.ListenAndServe(host, api.router))
}
