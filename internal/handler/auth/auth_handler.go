package auth

import (
	"net/http"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/auth"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		Password: hashPassword(req.Password),
		FullName: req.FullName,
	}

	err := h.u.Register(req.FullName, user)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, nil)
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

func hashPassword(password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash)
}
