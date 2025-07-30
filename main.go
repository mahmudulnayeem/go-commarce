package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecommarce/handlers"
)


func main(){
	mux:= http.NewServeMux()
	mux.HandleFunc("/",basePathHandler)
	mux.HandleFunc("/products",handlers.GetProducts)
	mux.HandleFunc("/product/create",handlers.CreteProduct)
	mux.HandleFunc("/product/update",handlers.UpdateProduct)
	mux.HandleFunc("/product/delete",handlers.DeleteProduct)

	fmt.Println("Server running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000",mux))
	
}

func basePathHandler(w http.ResponseWriter,r *http.Request){
fmt.Fprintln(w,"Hello go world")
}
