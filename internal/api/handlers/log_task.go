package handlers

import (
	"encoding/json"
	"mini-clickup/internal/api/models"
	"net/http"

	"gorm.io/gorm"
)

func LogTask(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	var ltsk models.TaskLog
	err := json.NewDecoder(r.Body).Decode(&ltsk)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = db.Create(&ltsk).Error
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
