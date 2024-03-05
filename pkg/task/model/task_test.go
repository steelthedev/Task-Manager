package model

import (
	"testing"

	"github.com/steelthedev/task-go/pkg/task/dto"
	"github.com/stretchr/testify/assert"
)

func TestNewTaskFail(t *testing.T) {

	requestBody := &dto.CreateTask{
		Title:  "",
		Status: dto.Choice(StatusChoiceDone),
	}

	task, err := NewTask(requestBody)
	assert.Nil(t, task, "we were not expecting any task")
	assert.NotNil(t, err, "we were expecting errors ")
	assert.EqualValues(t, err.Error(), "title cannot be empty")
}

func TestNewTaskPass(t *testing.T) {

	requestBody := &dto.CreateTask{
		Title:  "A new task",
		Status: dto.Choice(StatusChoiceNotStarted),
	}

	task, err := NewTask(requestBody)
	assert.NotNil(t, task, "we were expecting task")
	assert.Nil(t, err, "we were not expecting any error")
	assert.EqualValues(t, task.Title, "A new task")
	assert.EqualValues(t, task.Status, dto.Choice(StatusChoiceNotStarted))
}
