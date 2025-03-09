package routes

import (
	"firebase.google.com/go/auth"
	"github.com/nunthapongnp/time-doo-api/internal/controllers"
	"github.com/nunthapongnp/time-doo-api/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupApiRoutes(
	router *gin.Engine,
	authClient *auth.Client,
	taskController *controllers.TaskController,
	subTaskController *controllers.SubTaskController,
) {
	api := router.Group("/api/v1")
	api.Use(middleware.AuthMiddleware(authClient))
	{
		SetupTaskRoutes(api, taskController)
		SetupSubTaskRoutes(api, subTaskController)
	}
}
