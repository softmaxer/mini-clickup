package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)


func ExportLogs(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
  var tsklogs []models.TaskLog
  taskIDStr := r.URL.Query().Get("task_name")
  db.Where("task_name = ?", taskIDStr).Find(&tsklogs)

  jsonData, err := json.Marshal(tsklogs)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }

  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusOK)
  w.Write(jsonData)
}
