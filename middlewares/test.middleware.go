package middlewares

import (
	"fmt"
	"net/http"
)

func Test(next http.Handler) http.Handler {

	test := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("This is test middleware")

		next.ServeHTTP(w, r)
	}

	testHandler := http.HandlerFunc(test)
	return testHandler

}
