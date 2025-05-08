package auth

import (
	"net/http"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/auth"
	pwd "time-doo-api/pkg/bcrypt"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u auth.AuthUsecase
}

func NewAuthHandler(u auth.AuthUsecase) *Handler {
	return &Handler{u}
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

	user := &domain.User{
		Email:    req.Email,
		Password: pwd.HashPassword(req.Password),
		FullName: req.Email,
	}

	err := h.u.Register(req.Email, user)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	token, err := h.u.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, err, http.StatusUnauthorized)
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

	token, err := h.u.Login(req.Email, req.Password)
	if err != nil {
		response.Error(c, err, http.StatusUnauthorized)
		return
	}

	response.OK(c, gin.H{"token": token})
}
