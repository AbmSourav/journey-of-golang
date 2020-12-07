package main

import (
	"net/http"
)

func main() {
	// we got index, and about two html file in /static folder
	// so routes are: / , /about.html
	http.Handle("/", http.FileServer(http.Dir("./webDev/besics/01.renderStaticHTML/static/")))
	http.ListenAndServe(":3050", nil)
}
