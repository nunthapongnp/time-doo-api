package task

import (
	"net/http"
	"strconv"
	"time"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/task"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u task.TaskUsecase
}

func NewTaskHandler(u task.TaskUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/columns/:columnId/tasks")
	r.POST("", h.AddTask)
	r.GET("", h.GetTaskByColumn)
	r.PUT("/reorder", h.ReorderTask)

	t := rg.Group("/tasks")
	t.GET("/:taskId", h.GetTaskByID)
	t.PUT("/:taskId", h.EditTask)
	t.DELETE("/:taskId", h.RemoveTask)
	t.PUT("/tasks/move", h.MoveTaskToColumn)
}

func (h *Handler) AddTask(c *gin.Context) {
	columnID, _ := strconv.ParseInt(c.Param("columnId"), 10, 64)

	var input struct {
		ProjectID   int64      `json:"projectId"`
		Title       string     `json:"title"`
		Description string     `json:"description"`
		AssigneeID  *int64     `json:"assigneeId"`
		DueDate     *time.Time `json:"dueDate"`
		Position    int        `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	task := &domain.Task{
		ProjectID:   input.ProjectID,
		ColumnID:    columnID,
		Title:       input.Title,
		Description: input.Description,
		AssigneeID:  input.AssigneeID,
		DueDate:     input.DueDate,
		Position:    input.Position,
	}

	if err := h.u.AddTask(task); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, task)
}

func (h *Handler) GetTaskByColumn(c *gin.Context) {
	columnID, _ := strconv.ParseInt(c.Param("columnId"), 10, 64)

	tasks, err := h.u.GetTaskByColumn(columnID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, tasks)
}

func (h *Handler) GetTaskByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("taskId"), 10, 64)

	task, err := h.u.GetTaskByID(id)
	if err != nil {
		response.Error(c, err, http.StatusNotFound)
		return
	}

	response.OK(c, task)
}

func (h *Handler) EditTask(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("taskId"), 10, 64)

	var input struct {
		Title       string     `json:"title"`
		Description string     `json:"description"`
		AssigneeID  *int64     `json:"assigneeId"`
		DueDate     *time.Time `json:"dueDate"`
		ColumnID    int64      `json:"columnId"`
		Position    int        `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	task := &domain.Task{
		ID:          id,
		Title:       input.Title,
		Description: input.Description,
		AssigneeID:  input.AssigneeID,
		DueDate:     input.DueDate,
		ColumnID:    input.ColumnID,
		Position:    input.Position,
	}

	if err := h.u.EditTask(task); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, task)
}

func (h *Handler) RemoveTask(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("taskId"), 10, 64)

	if err := h.u.RemoveTask(id); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) ReorderTask(c *gin.Context) {
	columnID, _ := strconv.ParseInt(c.Param("columnId"), 10, 64)

	var input struct {
		TaskIDs []int64 `json:"taskIds"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	if err := h.u.ReorderTask(columnID, input.TaskIDs); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) MoveTaskToColumn(c *gin.Context) {
	var input struct {
		TaskID         int64   `json:"taskId"`
		FromColumnID   int64   `json:"fromColumnId"`
		ToColumnID     int64   `json:"toColumnId"`
		OrderedTaskIDs []int64 `json:"orderedTaskIds"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	if err := h.u.MoveTaskToColumn(input.TaskID, input.FromColumnID, input.ToColumnID, input.OrderedTaskIDs); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}
