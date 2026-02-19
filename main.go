package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Server running on :80")

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

	fmt.Fprint(w, "<h1>Home</h1><p>Welcome to branch 03</p>")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	fmt.Fprint(w, "<h1>About</h1><p>Now using net/http</p>")
}
