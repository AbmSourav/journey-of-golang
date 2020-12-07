package main

import (
	"html/template"
	"net/http"
)

func handeler(resw http.ResponseWriter, req *http.Request) {
	// template := template.Must(template.ParseFiles("./webDev/besics/02.template/public/index.gohtml"))
	template := template.Must(template.ParseFiles("./webDev/besics/02.template/public/index.html"))
	template.Execute(resw, "Universe")
}

func main() {
	http.HandleFunc("/", handeler)
	http.ListenAndServe(":3050", nil)
}
