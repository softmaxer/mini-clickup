package handlers

import (
	"net/http"

	"gorm.io/gorm"
)

func AnalyseHandler(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  taskIDStr := r.URL.Query().Get("user")
  url := "http://127.0.0.1:8050/analyse?user=" + taskIDStr
  http.Redirect(w, r, url, http.StatusSeeOther)
}
