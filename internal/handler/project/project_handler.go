package project

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/project"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u project.ProjectUsecase
}

func NewProjectHandler(u project.ProjectUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/projects")
	r.POST("/", h.AddProject)
	r.GET("/:projectId", h.GetProjectByID)
	r.GET("/tenant/:tenantId", h.GetProjectByTenant)
	r.PUT("/", h.EditProject)
	r.DELETE("/:projectId", h.RemoveProject)
	r.GET("/:projectId/members", h.GetProjectMembers)
}

func (h *Handler) AddProject(c *gin.Context) {
	var p domain.Project
	if err := c.ShouldBindJSON(&p); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	if err := h.u.AddProject(&p); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, p)
}

func (h *Handler) GetProjectByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)
	project, err := h.u.GetProjectByID(id)
	if err != nil {
		response.Error(c, err, http.StatusNotFound)
		return
	}

	response.OK(c, project)
}

func (h *Handler) GetProjectByTenant(c *gin.Context) {
	tenantID, _ := strconv.ParseInt(c.Param("tenantId"), 10, 64)
	projects, err := h.u.GetProjectByTenant(tenantID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, projects)
}

func (h *Handler) EditProject(c *gin.Context) {
	var p domain.Project
	if err := c.ShouldBindJSON(&p); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	if err := h.u.EditProject(&p); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, p)
}

func (h *Handler) RemoveProject(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)
	if err := h.u.RemoveProject(id); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) GetProjectMembers(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	members, err := h.u.GetProjectMembers(projectID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, members)
}
