package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)

func RemoveTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var tsk models.Task
	err := json.NewDecoder(r.Body).Decode(&tsk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
    return
	}

  var tsklogs []models.TaskLog
  var tsks []models.Task
  db.Where("task_name = ?", tsk.TaskName).Find(&tsklogs)
  db.Where("task_name = ?", tsk.TaskName).Find(&tsks)
  db.Delete(&tsklogs)
  db.Delete(&tsks)
}
