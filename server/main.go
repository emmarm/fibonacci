package main

import (
    "fmt"
		"github.com/julienschmidt/httprouter"
		"net/http"
		"log"
		"strconv"
		"encoding/json"
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

		n, err := strconv.Atoi(digit)
		type Response struct {
				Message   string		`json:"message"`
				Error		  string 		`json:"error"`
		}
		res := &Response{
			Message: "",
			Error: ""}

		if n > 93 {
				res.Error = "Please enter a smaller number"
		} else if n < 1 {
				res.Error = "Number must be a positive integer"
		} else if err != nil {
				res.Error = digit + "is not an integer"
		} else {
				num := fib()
				for i := 0; i < n; i++ {
						if i == n - 1 {
								res.Message += strconv.FormatInt(int64(num()), 10)
						} else {
								res.Message += strconv.FormatInt(int64(num()), 10) + ", "
						}
				}
		}

		resJson, err := json.Marshal(res)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(resJson)
}

func main() {
		router := httprouter.New()
		router.GET("/api", Index)
		router.GET("/api/hello/:name", Hello)
		router.GET("/api/fibonacci/:digit", Fibonacci)

		log.Fatal(http.ListenAndServe(":8080", router))
}
