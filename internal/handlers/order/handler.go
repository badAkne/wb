package order

import (
	"encoding/json"
	"net/http"
	"wb/internal/service"

	"github.com/gorilla/mux"
)

type OrderHandler struct {
	service service.OrderService
}

func NewOrderHandlers(service service.OrderService) OrderHandler {
	return OrderHandler{service: service}
}

func (h *OrderHandler) GetOrder(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	uid := vars["uid"]

	order, found, err := h.service.GetOrder(r.Context(), uid)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if !found {
		http.Error(w, "Order Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err = json.NewEncoder(w).Encode(order); err != nil {
		http.Error(w, "failed to encode a json", 500)
		return
	}
}
