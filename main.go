package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	path = path[10:]
	n, e := strconv.Atoi(path)
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("invalid"))
		return
	}
	fmt.Fprint(w, fizzbuzz(n))
}

func fizzbuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%3 == 0:
		return "Fizz"
	case n%5 == 0:
		return "Buzz"
	default:
		return strconv.Itoa(n)
	}
}

func main() {
	http.HandleFunc("/fizzbuzz/", handler)
	http.ListenAndServe(":8080", nil)
}
