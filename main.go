package main

import (
	"fmt"
	"net/http"

	"ali.com/market/controller"
	"ali.com/market/middlewares"
)

func getPort() string {
	const PORT int = 3000

	return fmt.Sprintf(":%d", PORT)
}

func main() {

	// create the router
	productMux := http.NewServeMux()

	// request-response cycle

	// productMux.Handle("OPTIONS /products", http.HandlerFunc(controller.GetAllProducts))
	productMux.Handle("GET /products", middlewares.CorsMiddleware(http.HandlerFunc(controller.GetAllProducts)))

	productMux.Handle("OPTIONS /create", http.HandlerFunc(controller.CreateProduct))
	productMux.Handle("POST /create", http.HandlerFunc(controller.CreateProduct))

	// create the server obj
	server := &http.Server{
		Addr:    getPort(),
		Handler: productMux,
	}

	// run the server
	fmt.Println("Marketing Server Is Running At http://localhost:3000")
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Find An Error To Start A Server : ", err)
	}
}
