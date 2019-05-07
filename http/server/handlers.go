package server

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/diskordanz/consumer/service/model"
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
	enableCors(&w)
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
	enableCors(&w)
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

	enableCors(&w)
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
	enableCors(&w)
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
	enableCors(&w)
	WriteEntityAndHeader(&w, locations)
}

func (api *ConsumerAPI) ListCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := api.ss.ListCategories()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enableCors(&w)
	WriteEntityAndHeader(&w, categories)
}

func (api *ConsumerAPI) ListProducts(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])
	name, _ := requestVariables["name"]

	products, err := api.ss.ListProducts(name, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enableCors(&w)
	WriteEntityAndHeader(&w, products)
}

func (api *ConsumerAPI) ListProductsByCategory(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])
	id, _ := strconv.Atoi(requestVariables["id"])
	name, _ := requestVariables["name"]

	products, err := api.ss.ListProductsByCategory(uint64(id), name, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enableCors(&w)
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
	enableCors(&w)
	WriteEntityAndHeader(&w, product)
}

func (api *ConsumerAPI) GetConsumer(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])

	consumer, err := api.ss.GetConsumer(id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enableCors(&w)
	WriteEntityAndHeader(&w, consumer)
}

func (api *ConsumerAPI) CreateConsumer(w http.ResponseWriter, r *http.Request) {

	var reqConsumer model.Consumer
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqConsumer); err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	consumerResult, err := api.ss.CreateConsumer(reqConsumer)

	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	enableCorsForCons(&w)
	WriteEntityAndHeader(&w, consumerResult)
}

func (api *ConsumerAPI) UpdateConsumer(w http.ResponseWriter, r *http.Request) {
	enableCorsForCons(&w)

	dec := json.NewDecoder(r.Body)
	var reqConsumer model.Consumer
	err := dec.Decode(&reqConsumer)
	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}

	consumerResult, err := api.ss.UpdateConsumer(reqConsumer)
	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}

	WriteEntityAndHeader(&w, consumerResult)
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
	enableCors(&w)
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
	enableCors(&w)
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
	enableCors(&w)
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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func enableCorsForCons(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
}

func (api *ConsumerAPI) Healthcheck(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	if api == nil || api.ss == nil || api.ss.Healthcheck() != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func WriteErrorEntityAndHeader(w *http.ResponseWriter, errorEntity error, statusCode int) {
	b, err := json.Marshal(errorEntity)
	writer := *w
	enableCorsForCons(w)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(statusCode)
	writer.Write(b)
}
