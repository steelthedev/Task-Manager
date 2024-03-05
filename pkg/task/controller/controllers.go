package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/steelthedev/task-go/pkg/task/services"
	"github.com/steelthedev/task-go/pkg/utils"
)

var (
	ControllerHandler handler
)

type handler struct {
	services.TaskServices
}

func (h handler) CreateTask(ctx *gin.Context) {

	body := dto.CreateTask{}
	if err := ctx.BindJSON(&body); err != nil {
		utils.ApiErrorFunction(ctx, http.StatusBadRequest, "Bad Request", err.Error())
		return
	}
	task, err := h.TaskServices.CreateTask(body)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, err.Message, err.Error)
		return
	}
	utils.ApiSuccessIndented(ctx, http.StatusCreated, "Created", &task)
}
