package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {
	// Serve static files (CSS, JS, etc.)
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	// Routes
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/submit", submitHandler)

	fmt.Println("Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("templates/index.html"))
	tmpl.Execute(w, nil)
}

func submitHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		message := r.FormValue("message")

		fmt.Fprintf(w, "Thanks %s! Your message: \"%s\" was received.", name, message)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}
