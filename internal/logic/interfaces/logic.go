package interfaces

import (
	"time"

	"surlit/internal/logic/models"
)

// TODO add context to all requests

type UserService interface { // TODO add link users, add project users
	Registration(user *models.User) error // TODO think about race conditions for username registration
	Authorization(login, password string) (*models.User, error)
	GetUserInfo(ID models.UUID) (*models.User, error)
	UpdateUserInfo(ID models.UUID, user *models.UpdateUser) error // TODO think about race conditions for username update
}

type ProjectService interface { // TODO add links project
	CreateProject(userID models.UUID, project *models.Project) error
	GetUserProjects(userID models.UUID) ([]*models.Project, error)
	GetProjectInfo(ID models.UUID) (*models.Project, error)
	UpdateProjectInfo(ID models.UUID, project *models.UpdateProject) error
}

type LinkService interface { // TODO add link users, add project users, generate new link short tag
	CreateLink(userID, projectID models.UUID, link *models.Link) error
	GetProjectLinks(projectID models.UUID) ([]*models.Link, error)
	GetUserLinks(userID models.UUID) ([]*models.Link, error)
	GetLinkInfo(ID models.UUID) (*models.Link, error)
	UpdateLinkInfo(ID models.UUID, link *models.UpdateLink) error
}

type RoutingService interface { // todo get route link
	CreateRoute(route *models.Route) error
	GetLinkRoutes(LinkID models.UUID) ([]*models.Route, error) // TODO add error if duplicate rule, and if all general
	GetRouteInfo(ID models.UUID) (*models.Route, error)
	UpdateRouteInfo(ID models.UUID, route *models.UpdateRoute) error
}

type StatService interface { // todo get route stat
	CreateStat(stat *models.Stat) error
	GetRouteStats(RouteID models.UUID) ([]*models.Stat, error)
	GetStatInfo(ID models.UUID) (*models.Stat, error)
}

type RedirectService interface {
	FindRedirect(ShortURL string, rec *models.RequestData) (string, error)
}

type HealthCheckService interface {
	RegisterChecks(check ...HealthCheck) HealthCheckService
	SetTimeout(timeout time.Duration) HealthCheckService
	SetTimePeriod(timePeriod time.Duration) HealthCheckService
	StartService()
}
