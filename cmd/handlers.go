package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type OrderRequest struct {
	ProductName string `json:"product_name"`
	Price       int    `json:"price"`
}

func (app *application) health(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, "working good", nil)
}

func (app *application) createOrder(w http.ResponseWriter, r *http.Request) {

	var order OrderRequest
	newId := uuid.NewString()

	if err := readJSON(w, r, &order); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid json")
		return
	}

	if strings.TrimSpace(order.ProductName) == "" {
		writeJSONError(w, http.StatusBadRequest, "Product name is reqired")
		return
	}

	if order.Price < 0 {
		writeJSONError(w, http.StatusBadRequest, "price is required")
		return
	}

	res, err := app.store.Order.Create(r.Context(), newId, order.ProductName, order.Price)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("error creating order: %v", err))
		return
	}

	writeJSON(w, http.StatusOK, "order created", res)

}

func (app *application) payOrder(w http.ResponseWriter, r *http.Request) {

	orderId := chi.URLParam(r, "id")

	if strings.TrimSpace(orderId) == "" {
		writeJSONError(w, http.StatusBadRequest, "order id is reqired")
		return
	}

	status, err := app.store.Order.GetStatus(r.Context(), orderId)
	if err != nil {
		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("%v", err))
		return
	}

	if status == "paid" {
		writeJSON(w, http.StatusOK, "already paid", nil)
		return
	}

	res, err := app.store.Order.Pay(r.Context(), orderId)
	if err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			writeJSONError(w, http.StatusNotFound, "order not found")
			return
		}

		writeJSONError(w, http.StatusInternalServerError, fmt.Sprintf("failed to pay order %v", err))
		return
	}

	app.pool.Enqueue(orderId)

	writeJSON(w, http.StatusOK, "order paid successfully", res)

}

func (app *application) checkStatus(w http.ResponseWriter, r *http.Request) {
	orderId := chi.URLParam(r, "id")

	if strings.TrimSpace(orderId) == "" {
		writeJSONError(w, http.StatusBadRequest, "order id is reqired")
		return
	}

	status, err := app.store.Order.GetStatus(r.Context(), orderId)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, fmt.Sprintf("%v", err))
		return
	}

	writeJSON(w, http.StatusOK, fmt.Sprintf("%v", status), nil)

}
