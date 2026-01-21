package main

import (
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/", homeHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	fmt.Println("Servidor rodando em http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}