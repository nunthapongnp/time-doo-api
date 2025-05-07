package user

import (
	"net/http"
	"strconv"
	"time-doo-api/internal/domain"
	"time-doo-api/internal/usecase/user"
	"time-doo-api/pkg/response"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	u user.UserUsecase
}

func NewUserHandler(u user.UserUsecase) *Handler {
	return &Handler{u}
}

func (h *Handler) Register(rg *gin.RouterGroup) {
	r := rg.Group("/users")
	r.GET("", h.GetAllUsers)
	r.GET("/:userId", h.GetUserByID)
	r.POST("", h.AddUser)
	r.PUT("/:userId", h.EditUser)
	r.DELETE("/:userId", h.RemoveUser)
	r.GET("/tenants/:tenantId", h.GetUserByTenant)
}

func (h *Handler) GetAllUsers(c *gin.Context) {
	users, err := h.u.GetAllUsers()
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, users)
}

func (h *Handler) GetUserByID(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	user, err := h.u.GetUserByID(id)
	if err != nil {
		response.Error(c, err, http.StatusNotFound)
		return
	}

	response.OK(c, user)
}

func (h *Handler) AddUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}
	if err := h.u.AddUser(&user); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.Created(c, user)
}

func (h *Handler) EditUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.Error(c, err, http.StatusBadRequest)
		return
	}
	user.ID = id
	if err := h.u.EditUser(&user); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, user)
}

func (h *Handler) RemoveUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("userId"), 10, 64)
	if err := h.u.RemoveUser(id); err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.NoContent(c)
}

func (h *Handler) GetUserByTenant(c *gin.Context) {
	tenantID, _ := strconv.ParseInt(c.Param("tenantId"), 10, 64)
	users, err := h.u.GetUserByTenant(tenantID)
	if err != nil {
		response.Error(c, err, http.StatusInternalServerError)
		return
	}

	response.OK(c, users)
}
