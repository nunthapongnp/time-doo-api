package tenant

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/tenant"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u tenant.TenantUsecase
}

func NewTenantHandler(u tenant.TenantUsecase) *Handler {
	return &Handler{u: u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/tenants")
	r.GET("/:tenantId/members", h.GetTenantMembers)
	r.POST("", h.AddTenant)
	r.GET("/:tenantId", h.GetTenantByID)
	r.GET("", h.GetAllTenants)
}

func (h *Handler) GetTenantMembers(c *gin.Context) {
	tenantID, _ := strconv.ParseInt(c.Param("tenantId"), 10, 64)
	members, err := h.u.GetTenantMembers(tenantID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, members)
}

func (h *Handler) AddTenant(c *gin.Context) {
	var req domain.Tenant
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}
	if err := h.u.AddTenant(&req); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, req)
}

func (h *Handler) GetTenantByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("tenantId"), 10, 64)
	tenant, err := h.u.GetTenantByID(id)
	if err != nil {
		response.Error(c, err, http.StatusNotFound)
		return
	}

	response.OK(c, tenant)
}

func (h *Handler) GetAllTenants(c *gin.Context) {
	tenants, err := h.u.GetAllTenants()
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, tenants)
}
