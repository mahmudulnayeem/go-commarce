package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ecommarce/middleware"
	"github.com/ecommarce/utils"
)

func Serve() {
	mux := http.NewServeMux()
	manager := middleware.NewManager()
	manager.Use(middleware.Logger)
	initRoutes(mux, manager)

	globalRouter := utils.GlobalRouter(mux)
	fmt.Println("Server running on localhost:3000")
	log.Fatal(http.ListenAndServe(":3000", globalRouter))
}

func basePathHandler(w http.ResponseWriter, r *http.Request) {
	utils.SendJSON(w, "Pong", http.StatusOK)
}
