package registation

import (
	"time-doo-api/internal/handler/auth"
	"time-doo-api/internal/handler/column"
	"time-doo-api/internal/handler/project"
	"time-doo-api/internal/handler/projectmember"
	"time-doo-api/internal/handler/task"
	"time-doo-api/internal/handler/tenant"
	"time-doo-api/internal/handler/tenantmember"
	"time-doo-api/internal/handler/user"
)

type AppHandler struct {
	AuthHandler          *auth.Handler
	userHandler          *user.Handler
	tenantHandler        *tenant.Handler
	tenantMemberHandler  *tenantmember.Handler
	projectHandler       *project.Handler
	projectMemberHandler *projectmember.Handler
	columnHandler        *column.Handler
	taskHandler          *task.Handler
}

func HandlerRegistation(u *AppUsecase) *AppHandler {
	return &AppHandler{
		AuthHandler:          auth.NewAuthHandler(u.tenantUsecase, u.tenantMemberUsecase, u.userUsecase),
		userHandler:          user.NewUserHandler(u.userUsecase),
		tenantHandler:        tenant.NewTenantHandler(u.tenantUsecase),
		tenantMemberHandler:  tenantmember.NewTenantMemberHandler(u.tenantMemberUsecase),
		projectHandler:       project.NewProjectHandler(u.projectUsecase),
		projectMemberHandler: projectmember.NewProjectMemberHandler(u.projectMemberUsecase),
		columnHandler:        column.NewColumnHandler(u.columnUsecase),
		taskHandler:          task.NewTaskHandler(u.taskUsecase),
	}
}
