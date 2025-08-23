
package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	_ "github.com/mattn/go-sqlite3"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var db *sql.DB

func main() {
	var err error
	db, err = sql.Open("sqlite3", "tasks.db")
	if err != nil {
		panic(err)
	}

	// Create table if not exists
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS tasks (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		title TEXT,
		completed BOOLEAN
	)`)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/tasks", tasksHandler)
	http.HandleFunc("/tasks/", taskHandler)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "frontend.html")
	})

	println("Server running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func tasksHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		rows, err := db.Query("SELECT id, title, completed FROM tasks")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		var allTasks []Task
		for rows.Next() {
			var t Task
			err := rows.Scan(&t.ID, &t.Title, &t.Completed)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			allTasks = append(allTasks, t)
		}
		if err := json.NewEncoder(w).Encode(allTasks); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPost:
		var t Task
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		res, err := db.Exec("INSERT INTO tasks (title, completed) VALUES (?, ?)", t.Title, t.Completed)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		id, _ := res.LastInsertId()
		t.ID = int(id)
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/tasks/"):] // get id from path
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	switch r.Method {
	case http.MethodGet:
		var t Task
		err := db.QueryRow("SELECT id, title, completed FROM tasks WHERE id = ?", id).Scan(&t.ID, &t.Title, &t.Completed)
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if err := json.NewEncoder(w).Encode(t); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodPut:
		var updated Task
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err := db.Exec("UPDATE tasks SET title = ?, completed = ? WHERE id = ?", updated.Title, updated.Completed, id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		updated.ID = id
		if err := json.NewEncoder(w).Encode(updated); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	case http.MethodDelete:
		_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
