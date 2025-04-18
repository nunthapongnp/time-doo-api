package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message,omitempty"`
}

type ErrorResponse struct {
	Success bool   `json:"success"`
	Error   string `json:"error"`
	Code    int    `json:"code"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, SuccessResponse{
		Success: true,
		Data:    data,
	})
}

func Message(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, SuccessResponse{
		Success: true,
		Message: msg,
	})
}

func NoContent(c *gin.Context) {
	c.JSON(http.StatusNoContent, SuccessResponse{
		Success: true,
	})
}

func Error(c *gin.Context, err error, status int) {
	c.AbortWithStatusJSON(status, ErrorResponse{
		Success: false,
		Error:   err.Error(),
		Code:    status,
	})
}
