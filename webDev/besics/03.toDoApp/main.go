package main

import (
	"html/template"
	"net/http"
)

type Todo struct {
    Title string
    Done  bool
}

type TodoPageData struct {
    PageTitle string
    Todos     []Todo
}

func dataMap(resw http.ResponseWriter, template template.Template) {
	data := TodoPageData{
		PageTitle: "ToDo List",
		Todos: []Todo{
			{Title: "Task 1", Done: true},
			{Title: "Task 2", Done: true},
			{Title: "Task 3", Done: false},
		},
	}

	template.Execute(resw, data)
}

func handeler(resw http.ResponseWriter, req *http.Request) {
	template := template.Must(template.ParseFiles("./webDev/besics/03.toDoApp/public/index.html"))
	
	dataMap(resw, *template)
}

func main() {
	http.HandleFunc("/", handeler)
	http.ListenAndServe(":3050", nil)
}
