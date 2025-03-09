package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/controllers"
)

func SetupSubTaskRoutes(router *gin.RouterGroup, tc *controllers.SubTaskController) {
	tasks := router.Group("/tasks")
	{
		tasks.POST("/:taskId/subtasks", tc.CreateSubtask)
		tasks.GET("/:taskId/subtasks/:subtaskId", tc.GetSubtask)
		tasks.PUT("/:taskId/subtasks/:subtaskId", tc.UpdateSubtask)
		tasks.DELETE("/:taskId/subtasks/:subtaskId", tc.DeleteSubtask)
	}
}
