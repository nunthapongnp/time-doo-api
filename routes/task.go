package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/controllers"
)

func SetupTaskRoutes(router *gin.RouterGroup, tc *controllers.TaskController) {
	tasks := router.Group("/tasks")
	{
		tasks.POST("/", tc.CreateTask)
		tasks.GET("/:taskId", tc.GetTask)
		tasks.PUT("/:taskId", tc.UpdateTask)
		tasks.DELETE("/:taskId", tc.DeleteTask)
	}
}
