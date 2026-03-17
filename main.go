package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type Team struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Message struct {
	Message string `json:"message"`
}

var db *sql.DB

func main() {
	connectDB()

	http.HandleFunc("/api/teams", teamsHandler)
	http.HandleFunc("/api/teams/", teamByIDHandler)
	http.HandleFunc("/api/ping", pingHandler)

	log.Println("REST DB API running on :80")
	log.Fatal(http.ListenAndServe(":80", nil))
}

func connectDB() {
	connStr := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" password=" + os.Getenv("DB_PASSWORD") +
		" dbname=" + os.Getenv("DB_NAME") +
		" sslmode=disable"

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

// ---------- HANDLERS ----------

func teamsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		rows, err := db.Query("SELECT id, name FROM teams")
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		defer rows.Close()

		var teams []Team
		for rows.Next() {
			var t Team
			rows.Scan(&t.ID, &t.Name)
			teams = append(teams, t)
		}

		writeJSON(w, 200, teams)

	case http.MethodPost:
		var t Team
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			http.Error(w, "Invalid JSON", 400)
			return
		}

		if t.Name == "" {
			http.Error(w, "Name required", 400)
			return
		}

		err = db.QueryRow(
			"INSERT INTO teams(name) VALUES($1) RETURNING id",
			t.Name,
		).Scan(&t.ID)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		writeJSON(w, 201, t)

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

func teamByIDHandler(w http.ResponseWriter, r *http.Request) {
	id, err := extractID(r.URL.Path)
	if err != nil {
		http.Error(w, "Invalid ID", 400)
		return
	}

	switch r.Method {

	case http.MethodGet:
		var t Team
		err := db.QueryRow(
			"SELECT id, name FROM teams WHERE id=$1",
			id,
		).Scan(&t.ID, &t.Name)

		if err == sql.ErrNoRows {
			http.Error(w, "Not Found", 404)
			return
		}

		writeJSON(w, 200, t)

	case http.MethodPut:
		var t Team
		json.NewDecoder(r.Body).Decode(&t)

		_, err := db.Exec(
			"UPDATE teams SET name=$1 WHERE id=$2",
			t.Name, id,
		)

		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		t.ID = id
		writeJSON(w, 200, t)

	case http.MethodDelete:
		_, err := db.Exec("DELETE FROM teams WHERE id=$1", id)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		writeJSON(w, 200, Message{Message: "Deleted"})

	default:
		http.Error(w, "Method Not Allowed", 405)
	}
}

// ---------- HELPERS ----------

func extractID(path string) (int, error) {
	parts := strings.Split(path, "/")
	return strconv.Atoi(parts[len(parts)-1])
}

func pingHandler(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, 200, Message{Message: "pong"})
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(payload)
}