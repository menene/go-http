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
	http.HandleFunc("/form", formHandler)

	log.Println("Server running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.html", nil)
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "about.html", nil)
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		renderTemplate(w, "form.html", nil)

	case http.MethodPost:
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Bad Request", http.StatusBadRequest)
			return
		}

		name := r.FormValue("name")

		data := map[string]string{
			"Name": name,
		}

		renderTemplate(w, "form.html", data)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func renderTemplate(w http.ResponseWriter, page string, data interface{}) {
	tmpl, err := template.ParseFiles(
		"./src/templates/layout.html",
		"./src/templates/"+page,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
