package models

import "gorm.io/gorm"

type Task struct {
	gorm.Model
	User      string  `json:"user"`
	TaskName  string  `json:"task_name"`
	Deadline  float64 `json:"deadline"`
	Finished  bool    `json:"finished"`
	StartDate float64 `json:"start_date"`
	EndDate   float64 `json:"end_date"`
  Importance string `json:"importance"`
}

type TaskLog struct {
	gorm.Model
	User      string  `json:"user"`
	TaskName  string  `json:"task_name"`
	TimeSpent float64 `json:"time_spent"`
	Date      float64 `json:"date"`
}
