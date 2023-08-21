package models

import "gorm.io/gorm"

type Task struct {
  gorm.Model
  Name string `json:"task_name"`
  TimeSpent float64 `json:"time_spent"`
  StartDate float64 `json:"start_date"`
  EndDate float64 `json:"end_date"`
  Deadline float64 `json:"deadline"`
  Importance string `json:"importance"`
}

type Project struct {
  tasks []Task
}
