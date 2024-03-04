package services

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/steelthedev/task-go/pkg/task/model"
	"github.com/steelthedev/task-go/pkg/utils"
	"gorm.io/gorm"
)

func CreateTask(ctx *gin.Context, db *gorm.DB) (*model.Task, *utils.AppError) {

	body := &dto.CreateTask{}
	if err := ctx.BindJSON(&body); err != nil {
		return nil, &utils.AppError{
			Message:    "Invalid body request",
			StatusCode: http.StatusBadRequest,
			Error:      err.Error(),
		}
	}

	newTask := model.Task{
		Title:  body.Title,
		Status: model.StatusChoiceNotStarted,
	}

	if result := db.Create(&newTask).Table("task"); result.Error != nil {
		return nil, &utils.AppError{
			Message:    "An internal Error occured",
			StatusCode: http.StatusInternalServerError,
			Error:      result.Error.Error(),
		}
	}
	return &newTask, nil

}
