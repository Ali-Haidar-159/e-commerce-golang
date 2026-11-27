package router

import (
	"net/http"

	"ali.com/market/controller"
	"ali.com/market/middlewares"
)

func ProductRouter(productMux *http.ServeMux, middlewareManager *middlewares.Middleware) {
	productMux.Handle("GET /products", middlewareManager.With(http.HandlerFunc(controller.GetAllProducts)))
	productMux.Handle("GET /products/{productId}", middlewareManager.With(http.HandlerFunc(controller.GetSingleProduct)))
	productMux.Handle("POST /products", middlewareManager.With(http.HandlerFunc(controller.CreateProduct)))

}
