package tenantmember

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/tenantmember"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u tenantmember.TenantMemberUsecase
}

func NewTenantMemberHandler(u tenantmember.TenantMemberUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/tenants/:tenantId/members")
	r.POST("", h.AddTenantMember)
	r.PUT("/:memberId", h.EditTenantMemberRole)
	r.DELETE("/:memberId", h.RemoveTenantMember)
}

func (h *Handler) AddTenantMember(c *gin.Context) {
	tenantID, _ := strconv.ParseInt(c.Param("tenantId"), 10, 64)
	var member domain.TenantMember
	if err := c.ShouldBindJSON(&member); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}
	member.TenantID = tenantID
	if err := h.u.AddTenantMember(&member); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, member)
}

func (h *Handler) EditTenantMemberRole(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("memberId"), 10, 64)
	var body struct {
		Role string `json:"role"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}
	if err := h.u.EditTenantMemberRole(id, body.Role); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) RemoveTenantMember(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("memberId"), 10, 64)
	if err := h.u.RemoveTenantMember(id); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}
