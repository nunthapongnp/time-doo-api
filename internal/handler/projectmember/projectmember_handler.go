package projectmember

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/projectmember"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u projectmember.ProjectMemberUsecase
}

func NewProjectMemberHandler(u projectmember.ProjectMemberUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/projects/:projectId/members")
	r.POST("", h.AddProjectMember)
	r.DELETE("/:userId", h.RemoveProjectMember)
}

func (h *Handler) AddProjectMember(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)

	var input struct {
		UserID int64  `json:"userId"`
		Role   string `json:"role"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	member := &domain.ProjectMember{
		ProjectID: projectID,
		UserID:    input.UserID,
		Role:      input.Role,
	}

	if err := h.u.AddProjectMember(member); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, member)
}

func (h *Handler) RemoveProjectMember(c *gin.Context) {
	projectID, _ := strconv.ParseInt(c.Param("projectId"), 10, 64)
	userID, _ := strconv.ParseInt(c.Param("userId"), 10, 64)

	if err := h.u.RemoveProjectMember(projectID, userID); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}
