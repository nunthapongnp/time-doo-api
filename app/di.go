package app

import (
	"time-doo-api/app/config"
	"time-doo-api/app/database"
	"time-doo-api/app/router"

	"github.com/gin-gonic/gin"
)

func InitializeApp() (*config.Config, *gin.Engine) {
	cfg := config.LoadConfig()
	db := database.ConnectDatabase(&cfg.Database)
	engine := router.SetupRouter(db)
	return cfg, engine
}
