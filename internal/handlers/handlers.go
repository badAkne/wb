package handlers

import "net/http"

type OrderHandler interface {
	GetOrder(w http.ResponseWriter, r *http.Request)
}
