package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)


func ExportTasks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  var tsks []models.Task
  taskIDStr := r.URL.Query().Get("user")
  db.Where("user = ?", taskIDStr).Find(&tsks)

  jsonData, err := json.Marshal(tsks)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonData)
}
