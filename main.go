package main

import (
	"net/http"

	"github.com/tquigly1287/testdevworkspace/mortgage"
)

func main() {
	http.HandleFunc("/", mortgage.GetBalance)
	http.ListenAndServe(":8080", nil)
}