package model

import "gorm.io/gorm"

type Choice string

const (
	StatusChoiceDone       Choice = "done"
	StatusChoiceDoing      Choice = "doing"
	StatusChoiceNotStarted Choice = "todo"
)

type Task struct {
	gorm.Model
	Title  string `json:"title"`
	Status Choice `json:"status"`
}
