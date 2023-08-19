package main

import (
	"api-mux/configs"
	"api-mux/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadgin .env file")
	}
	configs.Connection()
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Antes de passar a solicitação para o próximo handler, registre informações sobre a solicitação
		fmt.Printf("Received request: %s %s\n", r.Method, r.URL.Path)

		// Chame o próximo handler
		next.ServeHTTP(w, r)
	})
}

func main() {
	muxRouter := mux.NewRouter()
	muxRouter.Use(loggingMiddleware)
	routes.AppRoutes(muxRouter)

	fmt.Println("API is on port :8080")
	log.Fatal(http.ListenAndServe("127.0.0.1:8080", muxRouter))

	// configs.DB.AutoMigrate(&models.Users{})
}
