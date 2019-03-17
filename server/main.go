package main

import (
    "fmt"
		"github.com/julienschmidt/httprouter"
		"net/http"
		"log"
		"strconv"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func fib() func() int {
		a, b := 0, 1
		return func() int {
				c := a
				a, b = b, a+b
				return c
		}
}

func Fibonacci(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		digit := ps.ByName("digit")
		if n, err := strconv.Atoi(digit); err == nil {
				num := fib()
				for i := 0; i < n; i++ {
						if i == n - 1 {
								fmt.Fprintf(w, "%d", num())
						} else {
								fmt.Fprintf(w, "%d, ", num())
						}
				}
		} else {
				fmt.Fprintf(w, "%s is not an integer", digit)
		}
}

func main() {
		router := httprouter.New()
		router.GET("/api", Index)
		router.GET("/api/hello/:name", Hello)
		router.GET("/api/fibonacci/:digit", Fibonacci)

		log.Fatal(http.ListenAndServe(":8080", router))
}
