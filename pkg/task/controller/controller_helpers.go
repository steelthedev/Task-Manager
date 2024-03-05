package controller

import (
	"fmt"

	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/steelthedev/task-go/pkg/task/model"
	"gorm.io/gorm"
)

func GetTasksHelper(db *gorm.DB) (*[]model.Task, error) {
	var tasks []model.Task
	if result := db.Find(&tasks); result.Error != nil {
		return nil, fmt.Errorf("cannot get result from db")
	}
	return &tasks, nil
}

func CreateTaskHelper(requestBody *dto.CreateTask, db *gorm.DB) (*model.Task, error) {
	newTask, err := model.NewTask(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error occured while creating new task object")
	}

	if result := db.Create(&newTask); result.Error != nil {
		return nil, fmt.Errorf("cannot insert into db")
	}
	return newTask, nil
}
