package app

import (
	"fmt"
	"time-doo-api/app/config"
	"time-doo-api/app/database"
	"time-doo-api/app/router"

	"github.com/gin-gonic/gin"
)

func InitializeApp() (*config.Config, *gin.Engine) {
	cfg := config.LoadConfig()
	dsn := getDatabaseDsn(&cfg.Database)
	db := database.ConnectDatabase(dsn)
	engine := router.SetupRouter(db)
	return cfg, engine
}

func getDatabaseDsn(cfg *config.Database) string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s search_path=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		cfg.Host, cfg.User, cfg.Password, cfg.Name, cfg.Path, cfg.Port,
	)
}
