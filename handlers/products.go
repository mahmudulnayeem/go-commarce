package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/ecommarce/models"
	"github.com/ecommarce/utils"
)
var productList []models.Product

func init() {
	productList = append(productList, models.Product{
		Id:          1,
		Title:       "Mango",
		Description: "Mango is good",
		Price:       45.3,
		ImageURL:    "https://google.com",
		CreatedAt:   time.Now(),
	})
}

func GetProducts(w http.ResponseWriter,r *http.Request){
	utils.HandleCORS(w,r)
	utils.HandlePreflightReq(w,r)

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusBadRequest)
		return
	}
	utils.SendJSON(w,productList,http.StatusOK)
}

func CreteProduct(w http.ResponseWriter, r *http.Request){
		utils.HandleCORS(w,r)
		utils.HandlePreflightReq(w,r)

		if r.Method != http.MethodPost{
			http.Error(w,"Only post method is allowed",http.StatusBadRequest)
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

func UpdateProduct(w http.ResponseWriter, r *http.Request){
		utils.HandleCORS(w,r)
		utils.HandlePreflightReq(w,r)

		if r.Method != http.MethodPut{
			http.Error(w,"Only put method is allowed",http.StatusNotFound)
			return
		}

		idStr := r.URL.Query().Get("id")
		id,err:= strconv.Atoi(idStr)
		
		if err!=nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}
			var updatedProduct models.Product
		if err := json.NewDecoder(r.Body).Decode(&updatedProduct); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		found := false
		for i , p:= range productList{
			if p.Id==id {
			productList[i].Title = updatedProduct.Title
			productList[i].Description = updatedProduct.Description
			productList[i].Price = updatedProduct.Price
			productList[i].ImageURL = updatedProduct.ImageURL
			// Set UpdatedAt
			now := time.Now()
			productList[i].UpdatedAt = &now

			found = true
			utils.SendJSON(w, productList[i], http.StatusOK)
			break
			}
		}

	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
	
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
		utils.HandleCORS(w,r)
		utils.HandlePreflightReq(w,r)

		if r.Method != http.MethodDelete{
			http.Error(w,"Only delete method is allowed",http.StatusNotFound)
			return
		}

		idStr := r.URL.Query().Get("id")
		id,err:= strconv.Atoi(idStr)
		
		if err!=nil {
			http.Error(w, "Invalid product ID", http.StatusBadRequest)
			return
		}

		found := false
		for i , p:= range productList{
			if p.Id==id {
			productList = append(productList[:i], productList[i+1:]... )
			found = true
			utils.SendJSON(w,"Deleted", http.StatusNoContent)
			break
			}
		}

	if !found {
		http.Error(w, "Product not found", http.StatusNotFound)
	}
	
}