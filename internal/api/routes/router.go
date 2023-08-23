package routes

import (
	"log"
	"mini-clickup/internal/api/handlers"
	"mini-clickup/internal/database"
	"net/http"
)

func SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	db, err := database.SetupDatabase()
	if err != nil {
		log.Fatalln("Error setting up database.")
	}
	// Home page
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.HomeHandler(db, w, r)
	})
	// Add a new event
	mux.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		handlers.AddTask(db, w, r)
	})

	mux.HandleFunc("/log", func(w http.ResponseWriter, r *http.Request) {
		handlers.LogTask(db, w, r)
	})

	mux.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		handlers.RemoveTask(db, w, r)
	})

  mux.HandleFunc("/export-logs", func(w http.ResponseWriter, r *http.Request) {
    handlers.ExportLogs(db, w, r)
  })

  mux.HandleFunc("/export-tasks", func(w http.ResponseWriter, r *http.Request) {
    handlers.ExportTasks(db, w, r)
  })

	return mux
}
