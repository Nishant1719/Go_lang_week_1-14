package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/nishant/go-sqlite3-api/docs"

	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Simple Go SQLite API
// @version 1.0
// @description This is a simple REST API built with Go and SQLite.
// @host localhost:8080
// @BasePath /

type User struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Email     string `json:"email"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}

	createTable := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		firstname TEXT,
		lastname TEXT,
		email TEXT
	);`
	_, err = db.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}

	router := mux.NewRouter()

	router.HandleFunc("/users", createUser).Methods("POST")
	router.HandleFunc("/users", getUsers).Methods("GET")

	// Swagger endpoint
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

// createUser godoc
// @Summary Create a user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body User true "User Data"
// @Success 201 {object} User
// @Router /users [post]
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User
	json.NewDecoder(r.Body).Decode(&user)

	query := `INSERT INTO users (firstname, lastname, email) VALUES (?, ?, ?)`
	_, err := db.Exec(query, user.FirstName, user.LastName, user.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getUsers godoc
// @Summary Get all users
// @Description Fetch all users
// @Tags users
// @Produce json
// @Success 200 {array} User
// @Router /users [get]
func getUsers(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query(`SELECT firstname, lastname, email FROM users`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.FirstName, &user.LastName, &user.Email)
		users = append(users, user)
	}

	json.NewEncoder(w).Encode(users)
}
