package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (api *ConsumerAPI) ListFranchises(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])

	franchises, err := api.ss.ListFranchises(count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, franchises)
}

func (api *ConsumerAPI) GetFranchise(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])

	franchise, err := api.ss.GetFranchise(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, franchise)
}

func (api *ConsumerAPI) GetFranchisesByName(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)

	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])

	name := requestVariables["name"]

	franchiseList, err := api.ss.SearchFranchisesByName(count, offset, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	log.Printf("request on getFranchisesByName with name=%s", name)
	WriteEntityAndHeader(&w, franchiseList)
}

func (api *ConsumerAPI) GetLocationsOfFranchise(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	id, _ := strconv.Atoi(requestVariables["id"])
	offset, _ := strconv.Atoi(requestVariables["offset"])

	locations, err := api.ss.GetLocationsOfFranchise(id, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, locations)
}

func (api *ConsumerAPI) GetLocationsOfFranchiseByName(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	id, _ := strconv.Atoi(requestVariables["id"])
	offset, _ := strconv.Atoi(requestVariables["offset"])
	name := requestVariables["name"]

	locations, err := api.ss.GetLocationsOfFranchiseByName(id, count, offset, name)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, locations)
}

func (api *ConsumerAPI) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := api.ss.ListCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, categories)
}

func (api *ConsumerAPI) ListProducts(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])

	products, err := api.ss.ListProducts(count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, products)
}

func (api *ConsumerAPI) GetProduct(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])

	product, err := api.ss.GetProduct(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, product)
}

func (api *ConsumerAPI) GetProfile(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])

	profile, err := api.ss.GetProfile(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, profile)
}

/*
func (api *ConsumerAPI) GetCart(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])

	cart, err := api.ss.GetCart(id, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, cart)
}*/

func (api *ConsumerAPI) ListOrders(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])
	id, _ := strconv.Atoi(requestVariables["id"])

	orders, err := api.ss.ListOrders(id, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, orders)
}

func (api *ConsumerAPI) GetOrder(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	consumerID, _ := strconv.Atoi(requestVariables["consumer_id"])
	orderID, _ := strconv.Atoi(requestVariables["order_id"])

	order, err := api.ss.GetOrder(consumerID, orderID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, order)
}

func WriteEntityAndHeader(w *http.ResponseWriter, entity interface{}) {
	b, err := json.Marshal(entity)
	writer := *w
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.Write(b)
}
