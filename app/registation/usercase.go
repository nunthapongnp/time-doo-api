package registation

import (
	"time-doo-api/internal/usecase/auth"
	"time-doo-api/internal/usecase/column"
	"time-doo-api/internal/usecase/project"
	"time-doo-api/internal/usecase/projectmember"
	"time-doo-api/internal/usecase/task"
	"time-doo-api/internal/usecase/tenant"
	"time-doo-api/internal/usecase/tenantmember"
	"time-doo-api/internal/usecase/user"
)

type AppUsecase struct {
	authUsecase          auth.AuthUsecase
	userUsecase          user.UserUsecase
	tenantUsecase        tenant.TenantUsecase
	tenantMemberUsecase  tenantmember.TenantMemberUsecase
	projectUsecase       project.ProjectUsecase
	projectMemberUsecase projectmember.ProjectMemberUsecase
	columnUsecase        column.ColumnUsecase
	taskUsecase          task.TaskUsecase
}

func UsecaseRegistation(repo *AppRepository) *AppUsecase {
	return &AppUsecase{
		authUsecase:          auth.NewAuthUsecase(repo.userRepository, repo.tenantRepository, repo.tenantMemberRepository),
		userUsecase:          user.NewUserUsecase(repo.userRepository, repo.tenantMemberRepository),
		tenantUsecase:        tenant.NewTenantUsecase(repo.tenantRepository, repo.tenantMemberRepository),
		tenantMemberUsecase:  tenantmember.NewTenantMemberUsecase(repo.tenantMemberRepository),
		projectUsecase:       project.NewProjectUsecase(repo.projectRepository, repo.projectMemberRepository),
		projectMemberUsecase: projectmember.NewProjectMemberUsecase(repo.projectMemberRepository),
		columnUsecase:        column.NewColumnUsecase(repo.columnRepository),
		taskUsecase:          task.NewTaskUsecase(repo.taskRepository),
	}
}
