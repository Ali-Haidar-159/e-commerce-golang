package cmd

import (
	"fmt"
	"net/http"

	"ali.com/market/controller"
	"ali.com/market/middlewares"
	"ali.com/market/utils"
)

func getPort() string {
	const PORT int = 3000

	return fmt.Sprintf(":%d", PORT)
}

func Serve() {

	// create the router
	productMux := http.NewServeMux()

	// request-response cycle

	productMux.Handle("GET /products", middlewares.Test(middlewares.Logger(http.HandlerFunc(controller.GetAllProducts))))

	productMux.Handle("GET /products/{productId}", middlewares.Test(middlewares.Logger(http.HandlerFunc(controller.GetSingleProduct))))

	productMux.Handle("POST /products", middlewares.Test(middlewares.Logger(http.HandlerFunc(controller.CreateProduct))))

	// attach the global router with server
	handler := utils.GlobalRouter(productMux)

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
