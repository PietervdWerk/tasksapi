package api

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	tasksdb "github.com/pietervdwerk/tasksapi/internal/tasks"
	"github.com/pietervdwerk/tasksapi/pkg/openapi3"
)

// List all tasks
// (GET /tasks)
func (api *API) GetTasks(ctx context.Context, req GetTasksRequestObject) (GetTasksResponseObject, error) {
	dbTasks, err := api.tasksRepo.ListTasks(context.Background())
	if err != nil {
		return nil, err
	}

	var tasks GetTasks200JSONResponse
	for _, task := range dbTasks {
		tasks = append(tasks, tasksdb.TransformTask(task))
	}

	return tasks, nil
}

// Create a new task
// (POST /tasks)
func (api *API) PostTasks(ctx context.Context, req PostTasksRequestObject) (PostTasksResponseObject, error) {
	if req.Body.Title == "" {
		return PostTasks400JSONResponse{
			Message: "Title is required",
		}, nil
	}

	if req.Body.Status == "" {
		return PostTasks400JSONResponse{
			Message: "Status is required",
		}, nil
	}

	if req.Body.Status != openapi3.Completed && req.Body.Status != openapi3.Pending {
		return PostTasks400JSONResponse{
			Message: "Status must be 'completed' or 'pending'",
		}, nil
	}

	dbTask, err := api.tasksRepo.CreateTask(context.Background(), tasksdb.CreateTaskParams{
		ID:          uuid.New(),
		Title:       req.Body.Title,
		Description: req.Body.Description,
		Status:      req.Body.Status,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return PostTasks400JSONResponse{
			Message: "Task not found",
		}, nil
	} else if err != nil {
		return nil, err
	}

	return PostTasks201JSONResponse{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		Status:      dbTask.Status,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}, nil
}

// Delete a task
// (DELETE /tasks/{taskId})
func (api *API) DeleteTasksTaskId(ctx context.Context, req DeleteTasksTaskIdRequestObject) (DeleteTasksTaskIdResponseObject, error) {
	err := api.tasksRepo.DeleteTask(context.Background(), req.TaskId)
	if errors.Is(err, sql.ErrNoRows) {
		return DeleteTasksTaskId404JSONResponse{
			Message: "Task not found",
		}, nil
	} else if err != nil {
		return nil, err
	}

	return DeleteTasksTaskId204Response{}, nil
}

// Get a task by ID
// (GET /tasks/{taskId})
func (api *API) GetTasksTaskId(ctx context.Context, request GetTasksTaskIdRequestObject) (GetTasksTaskIdResponseObject, error) {
	task, err := api.tasksRepo.GetTask(context.Background(), request.TaskId)
	if errors.Is(err, sql.ErrNoRows) {
		return GetTasksTaskId404JSONResponse{
			Message: "Task not found",
		}, nil
	} else if err != nil {
		return nil, err
	}

	return GetTasksTaskId200JSONResponse{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      task.Status,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}, nil
}

// Update a task
// (PUT /tasks/{taskId})
func (api *API) PutTasksTaskId(ctx context.Context, request PutTasksTaskIdRequestObject) (PutTasksTaskIdResponseObject, error) {
	if request.Body.Title == "" {
		return PutTasksTaskId400JSONResponse{
			Message: "Title is required",
		}, nil
	}

	if request.Body.Status == "" {
		return PutTasksTaskId400JSONResponse{
			Message: "Status is required",
		}, nil
	}

	if request.Body.Status != openapi3.Completed && request.Body.Status != openapi3.Pending {
		return PutTasksTaskId400JSONResponse{
			Message: "Status must be 'completed' or 'pending'",
		}, nil
	}

	dbTask, err := api.tasksRepo.UpdateTask(context.Background(), tasksdb.UpdateTaskParams{
		ID:          request.TaskId,
		Title:       request.Body.Title,
		Description: request.Body.Description,
		Status:      request.Body.Status,
	})
	if errors.Is(err, sql.ErrNoRows) {
		return PutTasksTaskId404JSONResponse{
			Message: "Task not found",
		}, nil
	} else if err != nil {
		return nil, err
	}

	return PutTasksTaskId200JSONResponse{
		ID:          dbTask.ID,
		Title:       dbTask.Title,
		Description: dbTask.Description,
		Status:      dbTask.Status,
		CreatedAt:   dbTask.CreatedAt,
		UpdatedAt:   dbTask.UpdatedAt,
	}, nil
}
