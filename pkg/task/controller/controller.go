package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/steelthedev/task-go/pkg/utils"
	"gorm.io/gorm"
)

type handler struct {
	DB *gorm.DB
}

func (h *handler) GetTasks(ctx *gin.Context) {
	tasks, err := GetTasksHelper(h.DB)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusNotFound, err.Error(), err.Error())
		return
	}
	utils.ApiSuccessIndented(ctx, http.StatusOK, "okay", &tasks)
}

func (h *handler) CreateTask(ctx *gin.Context) {
	body := &dto.CreateTask{}
	if err := ctx.BindJSON(&body); err != nil {
		utils.ApiErrorFunction(ctx, http.StatusBadRequest, "Bad Request was sent", err.Error())
		return
	}
	task, err := CreateTaskHelper(body, h.DB)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, "Internal server error", err.Error())
		return
	}

	utils.ApiSuccessIndented(ctx, http.StatusOK, "", &task)

}