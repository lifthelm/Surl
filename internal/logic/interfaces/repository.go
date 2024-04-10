package interfaces

import "surlit/internal/logic/models"

// TODO add context to all requests

type UserRepository interface { // TODO add link users, add project users
	InsertUser(user *models.User) error
	Authorize(login, password string) (*models.User, error)
	GetUserInfo(uuid models.UUID) (*models.User, error)
	UpdateUserInfo(uuid models.UUID, user *models.UpdateUser) error
}

type ProjectRepository interface { // TODO add links project
	InsertProject(userID models.UUID, project *models.Project) error
	GetUserProjects(userID models.UUID) ([]*models.Project, error)
	GetProjectInfo(uuid models.UUID) (*models.Project, error)
	GetUserRoleInProject(userID, projectID models.UUID) (models.UserProjectRole, error)
	UpdateProjectInfo(uuid models.UUID, project *models.UpdateProject) error
}

type LinkRepository interface { // TODO get route link
	InsertLink(userID, projectID models.UUID, link *models.Link) error
	GetProjectLinks(projectID models.UUID) ([]*models.Link, error)
	GetUserLinks(userID models.UUID) ([]*models.Link, error)
	GetLinkInfo(uuid models.UUID) (*models.Link, error)
	FindLinkByToken(token string) (*models.Link, error)
	UpdateLinkInfo(uuid models.UUID, Link *models.UpdateLink) error
}

type RoutingRepository interface { // TODO add get stat Route
	InsertRoute(route *models.Route) error
	GetLinkRoutes(linkID models.UUID) ([]*models.Route, error)
	FindRoutes(linkID models.UUID, filter *models.RouteFilter) ([]*models.Route, error)
	GetRouteInfo(uuid models.UUID) (*models.Route, error)
	UpdateRouteInfo(uuid models.UUID, route *models.UpdateRoute) error
}

type StatsRepository interface {
	InsertStat(stat *models.Stat) error
	GetRouteStats(RouteID models.UUID) ([]*models.Stat, error)
	GetStatInfo(uuid models.UUID) (*models.Stat, error)
}
