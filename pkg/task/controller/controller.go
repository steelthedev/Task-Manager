package controller

import (
	"net/http"
	"strconv"

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

func (h *handler) DeleteTask(ctx *gin.Context) {
	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, "Errror occured somewhere", "Error occured while converting ID to int")
		return
	}
	err = DeleteTaskHelper(ID, h.DB)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	utils.ApiSuccessIndented(ctx, http.StatusOK, "Deleted succesfully", "")

}

func (h *handler) UpdateTask(ctx *gin.Context) {
	body := &dto.CreateTask{}
	if err := ctx.BindJSON(&body); err != nil {
		utils.ApiErrorFunction(ctx, http.StatusBadRequest, "Bad Request was sent", err.Error())
		return
	}

	id := ctx.Param("id")
	ID, err := strconv.Atoi(id)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, "Errror occured somewhere", "Error occured while converting ID to int")
		return
	}

	err = UpdateTaskHelper(ID, body, h.DB)
	if err != nil {
		utils.ApiErrorFunction(ctx, http.StatusInternalServerError, err.Error(), err.Error())
		return
	}

	utils.ApiSuccessIndented(ctx, http.StatusOK, "Updated Succesfully", "")

}
