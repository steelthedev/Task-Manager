package services

import (
	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/steelthedev/task-go/pkg/task/model"
	"github.com/steelthedev/task-go/pkg/task/repository"
	"github.com/steelthedev/task-go/pkg/utils"
)

type TaskServices struct {
	Storage repository.Storage
}

func (s *TaskServices) CreateTask(requestBody dto.CreateTask) (*model.Task, *utils.AppError) {
	task, err := model.NewTask(&requestBody)
	if err != nil {
		return nil, &utils.AppError{
			Message: "Error occured while sending request body to repository",
			Error:   err.Error(),
		}
	}
	task, err = s.Storage.CreateTask(task)
	if err != nil {
		return nil, &utils.AppError{
			Message: "Error Occured during saving layer",
			Error:   err.Error(),
		}
	}
	return task, nil
}
