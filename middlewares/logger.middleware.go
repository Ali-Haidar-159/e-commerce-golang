package middlewares

import (
	"fmt"
	"net/http"
)

func Logger(next http.Handler) http.Handler {

	logger := func(w http.ResponseWriter, r *http.Request) {

		//before execute the controller
		// start := time.Now()
		fmt.Println("Ami age print hobo")

		// execute the controller
		next.ServeHTTP(w, r)

		// after execute the controller
		// log.Println(r.Method, " ", r.URL.Path, " ", time.Since(start))
		fmt.Println("ami pore print hobo")
	}

	loggerHandler := http.HandlerFunc(logger)

	return loggerHandler

}
