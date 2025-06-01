package auth

import (
	"errors"
	"net/http"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/model"
	"time-doo-api/internal/usecase/tenant"
	"time-doo-api/internal/usecase/tenantmember"
	"time-doo-api/internal/usecase/user"
	"time-doo-api/pkg/jwt"
	"time-doo-api/pkg/response"

	pwd "time-doo-api/pkg/bcrypt"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	tenantUsecase       tenant.TenantUsecase
	tenantmemberUsecase tenantmember.TenantMemberUsecase
	userUsecase         user.UserUsecase
}

func NewAuthHandler(tenantUsecase tenant.TenantUsecase, tenantmemberUsecase tenantmember.TenantMemberUsecase, userUsecase user.UserUsecase) *Handler {
	return &Handler{tenantUsecase, tenantmemberUsecase, userUsecase}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	rg.POST("/register", h.RegisterTenant)
	rg.POST("/login", h.Login)
}

func (h *Handler) RegisterTenant(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	tenant := &domain.Tenant{
		Name: req.Email,
	}

	err := h.tenantUsecase.AddTenant(tenant)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	userDto := &model.UserDTO{
		Email:    req.Email,
		Password: pwd.HashPassword(req.Password),
		FullName: req.Email,
		TenantID: tenant.ID,
		Role:     "admin",
	}

	usr, err := h.userUsecase.AddUser(userDto)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	mem, err := h.tenantmemberUsecase.FindTenantMemberByUserID(usr.ID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	token, err := jwt.GenerateToken(uint(usr.ID), uint(mem.TenantID), mem.Role)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, gin.H{"token": token})
}

func (h *Handler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}

	usr, err := h.userUsecase.GetUserByEmail(req.Email)
	if err != nil {
		response.Error(c, err, http.StatusUnauthorized)
		return
	}

	mem, err := h.tenantmemberUsecase.FindTenantMemberByUserID(usr.ID)
	if err != nil {
		response.Error(c, err, http.StatusUnauthorized)
		return
	}

	if !pwd.VerifyPassword(usr.Password, req.Password) {
		response.Error(c, errors.New("invalid password"), http.StatusUnauthorized)
		return
	}

	token, err := jwt.GenerateToken(uint(usr.ID), uint(mem.TenantID), mem.Role)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, gin.H{"token": token})
}
