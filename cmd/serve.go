package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecommarce/handlers"
	"github.com/ecommarce/utils"
)

func Serve() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(basePathHandler))
	// products handlers
	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))
	mux.Handle("GET /products/{id}", http.HandlerFunc(handlers.GetProductById))
	mux.Handle("POST /products", http.HandlerFunc(handlers.CreteProduct))
	mux.Handle("PUT /products", http.HandlerFunc(handlers.UpdateProduct))
	mux.Handle("DELETE /product", http.HandlerFunc(handlers.DeleteProduct))

	globalRouter := utils.GlobalRouter(mux)
	fmt.Println("Server running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", globalRouter))
}

func basePathHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello go world")
}
