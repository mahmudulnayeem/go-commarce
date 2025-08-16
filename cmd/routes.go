package cmd

import (
	"net/http"

	"github.com/ecommarce/handlers"
	"github.com/ecommarce/middleware"
)

// initRoutes initializes the routes for the application
// and applies the middleware to each route.
// It sets up the HTTP handlers for various endpoints.
func initRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	mux.Handle("GET /ping", manager.With(http.HandlerFunc(basePathHandler)))
	productRoutes(mux, manager)
}

// Products routes
func productRoutes(mux *http.ServeMux, manager *middleware.Manager) {
	// Get all products
	mux.Handle("GET /products", manager.With(http.HandlerFunc(handlers.GetProducts)))
	// Get product by product id
	mux.Handle("GET /products/{id}", manager.With(http.HandlerFunc(handlers.GetProductById)))
	// Create a new product
	mux.Handle("POST /products", manager.With(http.HandlerFunc(handlers.CreteProduct)))
	// Update a product
	mux.Handle("PUT /products/{id}", manager.With(http.HandlerFunc(handlers.UpdateProduct)))
	// Delete a product
	mux.Handle("DELETE /products/{id}", manager.With(http.HandlerFunc(handlers.DeleteProduct)))
}
