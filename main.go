package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ecommarce/models"
	"github.com/ecommarce/utils"
)



var productList []models.Product

func main(){
	mux:= http.NewServeMux()
	mux.HandleFunc("/",basePathHandler)
	mux.HandleFunc("/products",getProducts)
	mux.HandleFunc("/create-product",creteProduct)

	fmt.Println("Server running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000",mux))
	
}

func basePathHandler(w http.ResponseWriter,r *http.Request){
fmt.Fprintln(w,"Hello world")
}

func getProducts(w http.ResponseWriter,r *http.Request){
	utils.HandleCORS(w,r)
	utils.HandlePreflightReq(w,r)

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}
	utils.SendJSON(w,productList,http.StatusOK)
}

func creteProduct(w http.ResponseWriter, r *http.Request){
		utils.HandleCORS(w,r)
		utils.HandlePreflightReq(w,r)

		if r.Method != http.MethodPost{
			http.Error(w,"Please give me post request",400)
			return
		}

		var  newProduct models.Product	
		if err := json.NewDecoder(r.Body).Decode(&newProduct); err != nil {
			http.Error(w, "Invalid JSON format", http.StatusBadRequest)
			return
		}
		newProduct.Id=len(productList)+1
		newProduct.CreatedAt = time.Now()
		productList = append(productList, newProduct)
		utils.SendJSON(w,newProduct,http.StatusCreated)
	
}

func init(){
prd1:= models.Product{
	Id: 1,
	 Title: "Mango",
	 Description: "Mango is good",
	 Price: 45.3,
	ImageURL:"https://google.com",
}
productList = append(productList, prd1)
}