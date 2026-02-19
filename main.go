package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server running on :80")

	// Static files
	css := http.FileServer(http.Dir("./src/css"))
	http.Handle("/css/", http.StripPrefix("/css/", css))

	assets := http.FileServer(http.Dir("./src/assets"))
	http.Handle("/assets/", http.StripPrefix("/assets/", assets))

	// Routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/about", aboutHandler)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		panic(err)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./src/index.html")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./src/about.html")
}
