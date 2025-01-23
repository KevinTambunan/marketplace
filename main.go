package main

import (
	"log"
	"marketplace/config"
	"marketplace/handlers"
	middlewares "marketplace/middleware"
	"marketplace/services"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	config.ConnectDB()
	defer config.CloseDB()

	port := ":8080"
	router := mux.NewRouter()

	userService := services.NewUserService(config.DB)
	userHandler := handlers.NewUserHandler(userService)

	router.HandleFunc("/login", userHandler.Login).Methods("POST")
	router.HandleFunc("/register", userHandler.Register).Methods("POST")
	router.HandleFunc("/users", middlewares.JWTMiddleware(userHandler.GetUserProfile)).Methods("GET")

	categoryService := services.NewCategoryService(config.DB)
	categoryHandler := handlers.NewCategoryHandler(categoryService)

	router.HandleFunc("/categories", categoryHandler.GetAllCategories).Methods("GET")
	router.HandleFunc("/categories", categoryHandler.CreateCategory).Methods("POST")
	router.HandleFunc("/categories/{id}", categoryHandler.GetCategoryById).Methods("GET")
	router.HandleFunc("/categories/{id}", categoryHandler.UpdateCategory).Methods("PUT")
	router.HandleFunc("/categories/{id}", categoryHandler.DeleteCategory).Methods("DELETE")

	statusService := services.NewStatusService(config.DB)
	statusHandler := handlers.NewStatusHandler(statusService)

	router.HandleFunc("/statuses", statusHandler.GetAllStatuses).Methods("GET")
	router.HandleFunc("/statuses", statusHandler.CreateStatus).Methods("POST")
	router.HandleFunc("/statuses/{id}", statusHandler.GetStatusById).Methods("GET")
	router.HandleFunc("/statuses/{id}", statusHandler.UpdateStatus).Methods("PUT")
	router.HandleFunc("/statuses/{id}", statusHandler.DeleteStatus).Methods("DELETE")

	productService := services.NewProductService(config.DB)
	productHandler := handlers.NewProductHandler(productService)

	router.HandleFunc("/products", productHandler.GetAllProducts).Methods("GET")
	router.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	router.HandleFunc("/products/{id}", productHandler.GetProductById).Methods("GET")
	router.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	router.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	orderService := services.NewOrderService(config.DB)
	orderHandler := handlers.NewOrderHandler(orderService)

	router.HandleFunc("/orders", orderHandler.GetAllOrders).Methods("GET")
	router.HandleFunc("/orders", orderHandler.CreateOrder).Methods("POST")
	router.HandleFunc("/orders/{id}", orderHandler.GetOrderById).Methods("GET")
	router.HandleFunc("/orders/{id}", orderHandler.UpdateOrder).Methods("PUT")
	router.HandleFunc("/orders/{id}", orderHandler.DeleteOrder).Methods("DELETE")

	log.Println("Server running at port", port)
	log.Fatal(http.ListenAndServe(port, router))
}
