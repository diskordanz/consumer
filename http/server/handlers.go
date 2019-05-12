package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
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
func parseToken(token string) (uint64, error) {
	var jwtKey = []byte("my_secret_key")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return 0, err
	}

	userID, err := strconv.ParseUint(claims["user_id"].(string), 10, 64)

	return userID, nil
}
func (api *ConsumerAPI) GetConsumer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])

	consumer, err := api.ss.GetConsumer(int(id))

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	WriteEntityAndHeader(&w, consumer)
}

func (api *ConsumerAPI) GetCart(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	count, _ := strconv.Atoi(requestVariables["count"])
	offset, _ := strconv.Atoi(requestVariables["offset"])
	id, _ := strconv.Atoi(requestVariables["id"])

	cart, err := api.ss.GetCart(id, count, offset)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	enableCors(&w)
	WriteEntityAndHeader(&w, cart)
}

func (api *ConsumerAPI) GetCartItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)

	requestVariables := mux.Vars(r)
	id, _ := strconv.Atoi(requestVariables["id"])
	product_id, _ := strconv.Atoi(requestVariables["product_id"])

	item, _ := api.ss.GetCartItem(id, product_id)

	WriteEntityAndHeader(&w, item)
}

func (api *ConsumerAPI) CreateCartItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var reqCartItem model.CartItem
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&reqCartItem); err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := api.ss.CreateCartItem(reqCartItem)
	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	WriteEntityAndHeader(&w, result)
}

func (api *ConsumerAPI) UpdateCartItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var reqCartItem model.CartItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqCartItem); err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	result, err := api.ss.UpdateCartItem(reqCartItem)

	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	WriteEntityAndHeader(&w, result)
}

func (api *ConsumerAPI) DeleteCartItem(w http.ResponseWriter, r *http.Request) {
	requestVariables := mux.Vars(r)
	enableCors(&w)
	id, _ := strconv.Atoi(requestVariables["id"])
	var reqItem model.CartItem
	reqItem.ID = uint64(id)

	err := api.ss.DeleteCartItem(reqItem)

	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (api *ConsumerAPI) CreateOrder(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var reqOrder model.Order
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqOrder); err != nil {
		fmt.Println(reqOrder.Status)
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	orderResult, err := api.ss.CreateOrder(reqOrder)

	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}

	WriteEntityAndHeader(&w, orderResult)
}

func (api *ConsumerAPI) CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
	var reqOrder model.OrderItem
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&reqOrder); err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	orderResult, err := api.ss.CreateOrderItem(reqOrder)

	if err != nil {
		WriteErrorEntityAndHeader(&w, err, http.StatusBadRequest)
		return
	}

	WriteEntityAndHeader(&w, orderResult)
}
func (api *ConsumerAPI) CreateConsumer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	WriteEntityAndHeader(&w, consumerResult)
}

func (api *ConsumerAPI) UpdateConsumer(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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

func (api *ConsumerAPI) ListOrders(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(&w)

	requestVariables := mux.Vars(r)
	orderID, _ := strconv.Atoi(requestVariables["id"])

	order, err := api.ss.GetOrder(orderID)

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

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, Access-Control-Allow-Origin")
}
func (api *ConsumerAPI) CORS(w http.ResponseWriter, r *http.Request) {
	enableCors(&w)
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
	enableCors(w)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		return
	}
	writer.WriteHeader(statusCode)
	writer.Write(b)
}
