package main

import (
	"net/http"
	"strings"
)

type OrderRequest struct {
	ProductName string `json:"product_name"`
	Price       *int   `json:"price"`
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, "working good", nil)
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {

	var order OrderRequest

	if err := readJSON(w, r, &order); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if strings.TrimSpace(order.ProductName) == "" {
		writeJSONError(w, http.StatusBadRequest, "Product name is reqired")
		return
	}

	if order.Price == nil {
		writeJSONError(w, http.StatusBadRequest, "price is required")
		return
	}

	writeJSON(w, http.StatusOK, "order created", nil)

}
