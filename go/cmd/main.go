package main

import (
	"net/http"

	function "github.com/pokutuna/example-adc-with-scope/go"
)

func main() {
	http.HandleFunc("/", function.App)
	http.ListenAndServe(":3000", nil)
}
