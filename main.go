package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	// static files
	css := http.FileServer(http.Dir("./src/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	// assets
	assets := http.FileServer(http.Dir("./src/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	log.Println("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	renderTemplate(w, "about.html")
}

func renderTemplate(w http.ResponseWriter, page string) {
	tmpl, err := template.ParseFiles(
		"./src/templates/layout.html",
		"./src/templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
