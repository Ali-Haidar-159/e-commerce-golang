package main

import (
	"fmt"
	"net/http"

	"ali.com/market/controller"
)

func getPort() string {
	const PORT int = 3000

	return fmt.Sprintf(":%d", PORT)
}

func globalRouter(mux *http.ServeMux) http.Handler {

	handleAllReq := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		mux.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(handleAllReq)
	return handler

}

func main() {

	// create the router
	productMux := http.NewServeMux()

	// request-response cycle

	// productMux.Handle("OPTIONS /products", http.HandlerFunc(controller.GetAllProducts))
	productMux.Handle("GET /products", http.HandlerFunc(controller.GetAllProducts))

	productMux.Handle("OPTIONS /create", http.HandlerFunc(controller.CreateProduct))
	productMux.Handle("POST /create", http.HandlerFunc(controller.CreateProduct))

	// attach the global router with server
	handler := globalRouter(productMux)

	// create the server obj
	server := &http.Server{
		Addr:    getPort(),
		Handler: handler,
	}

	// run the server
	fmt.Println("Marketing Server Is Running At http://localhost:3000")
	err := server.ListenAndServe()

	if err != nil && err != http.ErrServerClosed {
		fmt.Println("Find An Error To Start A Server : ", err)
	}
}
