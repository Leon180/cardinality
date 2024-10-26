package main

import (
	"cardinality/algo"
	"net/http"
)

func main() {

	http.HandleFunc("/morris", algo.NewMorrisCounterHandle().ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
