package routes

import (
  "net/http"
  "mini-clickup/internal/api/handlers"
)

func SetupRoutes() *http.ServeMux {
    mux := http.NewServeMux()
    // Home page
    mux.HandleFunc("/", handlers.HomeHandler)
    // Add a new event
    mux.HandleFunc("/add", handlers.AddTask)

    return mux
}

