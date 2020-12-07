package main

import (
	"fmt"
	"net/http"
)

func handeler(resw http.ResponseWriter, req *http.Request) {
	fmt.Fprint(resw, "<h1>Hello Universe</h1>")
}

func main() {
	http.HandleFunc("/", handeler)
	http.ListenAndServe(":3050", nil)
}
