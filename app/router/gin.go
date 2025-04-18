package router

import (
	"net/http"
	"time-doo-api/app/registation"
	"time-doo-api/internal/middleware"
	ctx "time-doo-api/pkg/context"
	"time-doo-api/pkg/logger"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	logger := logger.NewZapLogger()
	defer logger.Sync()

	r := gin.Default()
	r.Use(middleware.RecoveryMiddleware())
	r.Use(middleware.RecoveryWithLogger(logger))
	r.Use(middleware.RequestLogger(logger))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "OK"})
	})

	v1 := r.Group("/api/v1")
	v1.Use(middleware.JWTAuthMiddleware())

	appDbContext := ctx.NewAppDbContext(db)
	appRepository := registation.RepositoryRegistation(appDbContext)
	appUsecase := registation.UsecaseRegistation(appRepository)
	appHandler := registation.HandlerRegistation(appUsecase)
	registation.RoutesRegistation(v1, appHandler)

	auth := r.Group("/api/auth")
	appHandler.AuthHandler.Register(auth)

	return r
}
