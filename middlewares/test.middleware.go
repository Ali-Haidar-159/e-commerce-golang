package middlewares

import (
	"fmt"
	"net/http"
)

func Test(next http.Handler) http.Handler {

	test := func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("ami age print hobo -- test middleware")

		next.ServeHTTP(w, r)

		fmt.Println("Ami pore print hobo -- test middleware")
	}

	testHandler := http.HandlerFunc(test)
	return testHandler

}
