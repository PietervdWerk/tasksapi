package tasks

import (
	"github.com/pietervdwerk/tasksapi/pkg/openapi3"
)

func TransformTask(task Task) openapi3.Task {
	return openapi3.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}
