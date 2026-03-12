package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Deporte string `json:"deporte"`
	Liga    string `json:"liga"`
	Sede    string `json:"sede"`
	Year    int    `json:"year"`
	Titulos int    `json:"titulos"`
}

var filePath = "data/liga-nacional.json"

func loadTeams() ([]Team, error) {

	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var teams []Team

	err = json.Unmarshal(data, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func saveTeams(teams []Team) error {

	data, err := json.MarshalIndent(teams, "", " ")
	if err != nil {
		return err
	}

	return os.WriteFile(filePath, data, 0644)
}

func errorJSON(w http.ResponseWriter, code int, msg string) {

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	json.NewEncoder(w).Encode(map[string]string{
		"error": msg,
	})
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {

	path := strings.TrimPrefix(r.URL.Path, "/api/items")

	if path == "" || path == "/" {
		handleCollection(w, r)
		return
	}

	idStr := strings.Trim(path, "/")

	id, err := strconv.Atoi(idStr)
	if err != nil {
		errorJSON(w, 400, "invalid id")
		return
	}

	handleItem(w, r, id)
}

func handleCollection(w http.ResponseWriter, r *http.Request) {

	teams, err := loadTeams()
	if err != nil {
		errorJSON(w, 500, "error loading data")
		return
	}

	switch r.Method {

	case http.MethodGet:

		idParam := r.URL.Query().Get("id")

		if idParam != "" {

			id, _ := strconv.Atoi(idParam)

			for _, t := range teams {
				if t.ID == id {
					json.NewEncoder(w).Encode(t)
					return
				}
			}

			errorJSON(w, 404, "team not found")
			return
		}

		json.NewEncoder(w).Encode(teams)

	case http.MethodPost:

		var team Team

		err := json.NewDecoder(r.Body).Decode(&team)
		if err != nil {
			errorJSON(w, 400, "invalid json")
			return
		}

		teams = append(teams, team)

		saveTeams(teams)

		w.WriteHeader(201)
		json.NewEncoder(w).Encode(team)

	default:

		errorJSON(w, 405, "method not allowed")

	}
}

func handleItem(w http.ResponseWriter, r *http.Request, id int) {

	teams, err := loadTeams()
	if err != nil {
		errorJSON(w, 500, "error loading data")
		return
	}

	index := -1

	for i, t := range teams {

		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		errorJSON(w, 404, "team not found")
		return
	}

	switch r.Method {

	case http.MethodGet:

		json.NewEncoder(w).Encode(teams[index])

	case http.MethodPut:

		var updated Team

		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			errorJSON(w, 400, "invalid json")
			return
		}

		teams[index] = updated

		saveTeams(teams)

		json.NewEncoder(w).Encode(updated)

	case http.MethodPatch:

		var patch map[string]interface{}

		json.NewDecoder(r.Body).Decode(&patch)

		if name, ok := patch["name"].(string); ok {
			teams[index].Name = name
		}

		if titulos, ok := patch["titulos"].(float64); ok {
			teams[index].Titulos = int(titulos)
		}

		saveTeams(teams)

		json.NewEncoder(w).Encode(teams[index])

	case http.MethodDelete:

		teams = append(teams[:index], teams[index+1:]...)

		saveTeams(teams)

		json.NewEncoder(w).Encode(map[string]string{
			"message": "deleted",
		})

	default:

		errorJSON(w, 405, "method not allowed")
	}
}

func main() {

	port := "3012"

	http.HandleFunc("/api/items", itemsHandler)
	http.HandleFunc("/api/items/", itemsHandler)

	fmt.Println("Server running on port", port)

	http.ListenAndServe(":"+port, nil)
}
