package main

import (
	"fmt"
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

func main() {
    tmpl := template.Must(template.ParseFiles("page.html"))
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        data := TodoPageData{
            PageTitle: "My TODO list",
            Todos: []Todo{
                {Title: "Task 1", Done: false},
                {Title: "Task 2", Done: true},
                {Title: "Task 3", Done: true},
            },
        }
        tmpl.Execute(w, data)
    })
    http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("hi")
    })
    http.ListenAndServe(":80", nil)
}
