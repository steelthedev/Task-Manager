package repository

import (
	"github.com/steelthedev/task-go/pkg/task/model"
	"github.com/steelthedev/task-go/pkg/utils"
	"gorm.io/gorm"
)

type Storage interface {
	CreateTask(task *model.Task) (*model.Task, error)
}

type Repository struct {
	DB *gorm.DB
}

func (r Repository) CreateTask(task *model.Task) (*model.Task, *utils.AppError) {

	if result := r.DB.Create(&task).Table("task"); result.Error != nil {
		return nil, &utils.AppError{
			Message: "Error occured while adding to database",
			Error:   result.Error.Error(),
		}
	}
	return task, nil
}
