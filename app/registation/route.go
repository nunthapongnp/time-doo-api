package registation

import "github.com/gin-gonic/gin"

func RoutesRegistation(router *gin.RouterGroup, handler *AppHandler) {
	handler.userHandler.Register(router)
	handler.tenantHandler.Register(router)
	handler.tenantMemberHandler.Register(router)
	handler.projectHandler.Register(router)
	handler.projectMemberHandler.Register(router)
	handler.columnHandler.Register(router)
	handler.taskHandler.Register(router)
}
