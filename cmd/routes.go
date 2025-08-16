package cmd

import (
	"net/http"

	"github.com/ecommarce/handlers"
	"github.com/ecommarce/middleware"
)

func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /ping", manager.With(http.HandlerFunc(basePathHandler)))
	// products handlers
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreteProduct), middleware.Logger))
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))
}
