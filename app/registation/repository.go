package registation

import (
	"time-doo-api/internal/repository/column"
	"time-doo-api/internal/repository/project"
	"time-doo-api/internal/repository/projectmember"
	"time-doo-api/internal/repository/task"
	"time-doo-api/internal/repository/tenant"
	"time-doo-api/internal/repository/tenantmember"
	"time-doo-api/internal/repository/user"
	ctx "time-doo-api/pkg/context"
)

type AppRepository struct {
	userRepository          user.UserRepository
	tenantRepository        tenant.TenantRepository
	tenantMemberRepository  tenantmember.TenantMemberRepository
	projectRepository       project.ProjectRepository
	projectMemberRepository projectmember.ProjectMemberRepository
	columnRepository        column.ColumnRepository
	taskRepository          task.TaskRepository
}

func RepositoryRegistation(db *ctx.AppDbContext) *AppRepository {
	return &AppRepository{
		userRepository:          user.NewUserRepository(db),
		tenantRepository:        tenant.NewTenantRepository(db),
		tenantMemberRepository:  tenantmember.NewTenantMemberRepository(db),
		projectRepository:       project.NewProjectRepository(db),
		projectMemberRepository: projectmember.NewProjectMemberRepository(db),
		columnRepository:        column.NewColumnRepository(db),
		taskRepository:          task.NewTaskRepository(db),
	}
}
