package cmd

import (
	"fmt"
	"net/http"

	"ali.com/market/middlewares"
	"ali.com/market/router"
	"ali.com/market/utils"
)

func getPort() string {
	const PORT int = 3000

	return fmt.Sprintf(":%d", PORT)
}

func Serve() {

	// create the router
	productMux := http.NewServeMux()

	// CORS and preflight handle
	middlewareManager := middlewares.NewMiddleware()
	middlewareManager.Use(middlewares.Logger, middlewares.Test)

	// request-response cycle
	router.ProductRouter(productMux, middlewareManager)

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
