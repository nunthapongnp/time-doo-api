package column

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/column"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u column.ColumnUsecase
}

func NewColumnHandler(u column.ColumnUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/projects/:projectId/columns")
	r.POST("", h.AddColumn)
	r.GET("", h.GetColumnByProject)
	r.PUT("/:columnId", h.EditColumn)
	r.DELETE("/:columnId", h.RemoveColumn)
	r.PUT("/reorder", h.ReorderColumn)
	r.GET("/tasks", h.GetColumnWithTasks)
}

func (h *Handler) AddColumn(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	var input struct {
		Name     string `json:"name"`
		Position int64  `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	col := &domain.Column{
		ProjectID: projectID,
		Name:      input.Name,
		Position:  input.Position,
	}

	if err := h.u.AddColumn(col); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, col)
}

func (h *Handler) GetColumnByProject(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	cols, err := h.u.GetColumnByProject(projectID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, cols)
}

func (h *Handler) EditColumn(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("columnId"), 10, 64)

	var input struct {
		Name     string `json:"name"`
		Position int64  `json:"position"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	col := &domain.Column{
		ID:       id,
		Name:     input.Name,
		Position: input.Position,
	}

	if err := h.u.EditColumn(col); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, col)
}

func (h *Handler) RemoveColumn(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("columnId"), 10, 64)

	if err := h.u.RemoveColumn(id); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) ReorderColumn(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	var input struct {
		ColumnIDs []int64 `json:"columnIds"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	if err := h.u.ReorderColumn(projectID, input.ColumnIDs); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) GetColumnWithTasks(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	cols, err := h.u.GetColumnWithTasks(projectID)
	if err != nil {
		response.Error(c, err, http.StatusBadRequest)
	}

	response.OK(c, cols)
}
