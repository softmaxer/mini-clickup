package api

import (
	"mini-clickup/internal/api/routes"
	"net/http"
)

func StartServer() {
	mux := routes.SetupRoutes()
  http.ListenAndServe(":8080", mux)
}
