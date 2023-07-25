package router

import (
	"go-shopping-cart/middleware"

	"github.com/gorilla/mux"
)

//Router is exported and used in main.go


func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/newproduct", middleware.CreateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/getproduct/{id}", middleware.GetProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/getallproduct", middleware.GetAllProduct).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/updateproduct", middleware.UpdateProduct).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/deleteproduct/{id}", middleware.DeleteProduct).Methods("POST", "OPTIONS")
return router
}
