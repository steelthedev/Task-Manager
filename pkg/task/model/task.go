package model

import (
	"fmt"

	"github.com/steelthedev/task-go/pkg/task/dto"
	"gorm.io/gorm"
)

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

func NewTask(requestBody *dto.CreateTask) (*Task, error) {
	if err := ValidateNewTaskRequestBody(requestBody); err != nil {
		return nil, err
	}

	newTask := &Task{
		Title:  requestBody.Title,
		Status: StatusChoiceNotStarted,
	}
	return newTask, nil

}

func ValidateNewTaskRequestBody(requestBody *dto.CreateTask) error {
	if len(requestBody.Title) < 1 {
		return fmt.Errorf("title cannot be empty")
	}
	return nil
}
