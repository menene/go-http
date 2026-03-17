package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Message string `json:"message"`
}

var teams []Team

func main() {
	loadTeams()

	http.HandleFunc("/api/teams", teamsHandler)
	http.HandleFunc("/api/teams/", teamByIDHandler)
	http.HandleFunc("/api/ping", pingHandler)

	log.Println("REST API running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

// ---------- LOAD DATA ----------

func loadTeams() {
	file, err := os.ReadFile("./data/teams.json")
	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(file, &teams)
	if err != nil {
		log.Fatal(err)
	}
}

// ---------- HANDLERS ----------

// /api/teams
func teamsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		writeJSON(w, http.StatusOK, teams)

	case http.MethodPost:
		var newTeam Team

		err := json.NewDecoder(r.Body).Decode(&newTeam)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if newTeam.Name == "" {
			http.Error(w, "Name is required", http.StatusBadRequest)
			return
		}

		newTeam.ID = generateNextID()
		teams = append(teams, newTeam)

		writeJSON(w, http.StatusCreated, newTeam)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// /api/teams/{id}
func teamByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	switch r.Method {

	case http.MethodGet:
		team, found := findTeam(id)
		if !found {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		writeJSON(w, http.StatusOK, team)

	case http.MethodPut:
		var updated Team

		err := json.NewDecoder(r.Body).Decode(&updated)
		if err != nil {
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		for i, t := range teams {
			if t.ID == id {
				teams[i].Name = updated.Name
				writeJSON(w, http.StatusOK, teams[i])
				return
			}
		}

		http.Error(w, "Not Found", http.StatusNotFound)

	case http.MethodDelete:
		for i, t := range teams {
			if t.ID == id {
				teams = append(teams[:i], teams[i+1:]...)
				writeJSON(w, http.StatusOK, Message{Message: "Deleted"})
				return
			}
		}

		http.Error(w, "Not Found", http.StatusNotFound)

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// ---------- HELPERS ----------

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]
	return strconv.Atoi(idStr)
}

func findTeam(id int) (Team, bool) {
	for _, t := range teams {
		if t.ID == id {
			return t, true
		}
	}
	return Team{}, false
}

func generateNextID() int {
	max := 0
	for _, t := range teams {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, Message{Message: "pong"})
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}