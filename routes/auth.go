package routes

import (
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/nunthapongnp/time-doo-api/internal/controllers"
)

func SetupPublicRoutes(router *gin.Engine, authClient *auth.Client, authController *controllers.AuthController) {
	public := router.Group("/api/v1/auth")
	{
		public.POST("/get-id-token", authController.GetFirebaseIDToken)
	}
}
