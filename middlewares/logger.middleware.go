package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logger(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter, r *http.Request) {

		//before execute the controller
		start := time.Now()

		// execute the controller
		next.ServeHTTP(w, r)

		// after execute the controller
		log.Println(r.Method, " ", r.URL.Path, " ", time.Since(start))
	}

	loggerHandler := http.HandlerFunc(logger)

	return loggerHandler

}
