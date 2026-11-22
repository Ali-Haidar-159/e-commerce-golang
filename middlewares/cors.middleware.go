package middlewares

import "net/http"

func CorsMiddleware(next http.Handler) http.Handler {

	corsHandler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	}

	handler := http.HandlerFunc(corsHandler)

	return handler

}
