package tech

import (
	"bufio"
	"fmt"
	"os"
	modelsUI "surlit/internal/app/tech/models"
	"surlit/internal/app/tech/utils"
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type Handle func()

type Handler struct {
	userServ     interfaces.UserService
	projectServ  interfaces.ProjectService
	linkServ     interfaces.LinkService
	routingServ  interfaces.RoutingService
	redirectServ interfaces.RedirectService
	statsServ    interfaces.StatService
	healthServ   interfaces.HealthCheckService

	reader InputReader
	cache  map[int]any
}

func NewService(userS interfaces.UserService,
	projectS interfaces.ProjectService,
	linkS interfaces.LinkService,
	routingS interfaces.RoutingService,
	redirectS interfaces.RedirectService,
	statsS interfaces.StatService,
	healthS interfaces.HealthCheckService) *Handler {

	reader := NewInputReaderBufio(bufio.NewReader(os.Stdin))
	return &Handler{
		userServ:     userS,
		projectServ:  projectS,
		linkServ:     linkS,
		routingServ:  routingS,
		redirectServ: redirectS,
		statsServ:    statsS,
		healthServ:   healthS,
		cache:        make(map[int]any),
		reader:       reader,
	}
}
func (s *Handler) printUser(user *models.User) {
	userRole, err := utils.UserRoleLogicToUI(user.UserRole)
	if err != nil {
		fmt.Println("cant convert user role to ui")
	}
	fmt.Printf("\nID: %d\n", user.ID)
	fmt.Printf("User: %s\n", user.UserName)
	fmt.Printf("Email: %s\n", user.Email)
	fmt.Printf("Role: %s\n", userRole)
	fmt.Printf("Created: %s\n\n", user.CreatedAt)

}

func (s *Handler) loginHandler() {
	fmt.Printf("Enter login: ")
	login, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read login")
		return
	}
	fmt.Printf("Enter password: ")
	password, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read password")
		return
	}
	user, err := s.userServ.Authorization(login, password)
	if err != nil {
		fmt.Println("cant authorize user")
		return
	}
	s.printUser(user)
	s.cache[modelsUI.KeyUser] = modelsUI.UserAuthorized{UserID: user.ID}
	return
}

func (s *Handler) registrationHandler() {
	fmt.Printf("Enter login: ")
	login, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read login")
		return
	}
	fmt.Printf("Enter password: ")
	password, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read password")
		return
	}
	fmt.Printf("Enter email: ")
	email, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read email")
		return
	}
	newUser := models.User{
		UserName: login,
		Password: password,
		Email:    email,
		UserRole: models.UserRegular,
	}
	err = s.userServ.Registration(&newUser)
	if err != nil {
		fmt.Println("cant register user")
		return
	}
	return
}

func (s *Handler) createProjectHandler() {
	fmt.Printf("Enter project name: ")
	name, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read project name")
		return
	}
	fmt.Printf("Enter project description: ")
	description, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read project description")
		return
	}
	user, ok := s.cache[modelsUI.KeyUser].(modelsUI.UserAuthorized)
	if !ok {
		fmt.Println("need to authorize to create project")
		return
	}
	newProject := models.Project{
		Name:        name,
		Description: description,
	}
	err = s.projectServ.CreateProject(user.UserID, &newProject)
	if err != nil {
		fmt.Println("cant add project")
		return
	}
	fmt.Printf("Project %s created\n", name)
}

func (s *Handler) getUserProjectsHandler() {
	user, ok := s.cache[modelsUI.KeyUser].(modelsUI.UserAuthorized)
	if !ok {
		fmt.Println("need to authorize to get projects")
		return
	}
	projects, err := s.projectServ.GetUserProjects(user.UserID)
	if err != nil {
		fmt.Println("cant get projects")
		return
	}
	for _, project := range projects {
		fmt.Printf("Project ID: %d\n", project.ID)
		fmt.Printf("Project Name: %s\n", project.Name)
		fmt.Printf("Description: %s\n\n", project.Description)
	}
}

func (s *Handler) getUserLinksHandler() {
	user, ok := s.cache[modelsUI.KeyUser].(modelsUI.UserAuthorized)
	if !ok {
		fmt.Println("need to authorize to get links")
		return
	}
	links, err := s.linkServ.GetUserLinks(user.UserID)
	if err != nil {
		fmt.Println("cant get links")
		return
	}
	for _, link := range links {
		fmt.Printf("Link ID: %d\n", link.ID)
		fmt.Printf("Link Created: %s\n", link.CreatedAt)
		fmt.Printf("Link URL: %s\n", link.DefLongURL)
		fmt.Printf("Short tag: %s\n\n", link.URLShort)
	}
}

func (s *Handler) getProjectHandler() {
	fmt.Printf("Enter project ID: ")
	projectID, err := s.reader.ReadUUID()
	if err != nil {
		fmt.Println("cant read project ID")
		return
	}
	project, err := s.projectServ.GetProjectInfo(projectID)
	if err != nil {
		fmt.Printf("cant get project: %s\n", err)
		return
	}
	fmt.Printf("Project ID: %d\n", project.ID)
	fmt.Printf("Project Name: %s\n", project.Name)
	fmt.Printf("Description: %s\n\n", project.Description)
}

func (s *Handler) getLinkHandler() {
	fmt.Printf("Enter link ID: ")
	linkID, err := s.reader.ReadUUID()
	if err != nil {
		fmt.Println("cant read link ID")
		return
	}
	link, err := s.linkServ.GetLinkInfo(linkID)
	if err != nil {
		fmt.Printf("cant get link: %s\n", err)
		return
	}
	fmt.Printf("Link ID: %d\n", link.ID)
	fmt.Printf("Link URL: %s\n", link.DefLongURL)
	fmt.Printf("Link Tag: %s\n", link.URLShort)
	fmt.Printf("Link Craeted: %s\n\n", link.CreatedAt)
}

func (s *Handler) createRoute() {
	fmt.Printf("Enter link ID: ")
	linkID, err := s.reader.ReadUUID()
	if err != nil {
		fmt.Println("cant read link ID")
		return
	}
	fmt.Printf("Enter URL: ")
	url, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read URL")
		return
	}
	fmt.Printf("Enter platform: ")
	platform, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read platform")
		return
	}
	fmt.Printf("Enter availability zone: ")
	azone, err := s.reader.ReadString()
	if err != nil {
		fmt.Println("cant read availability zone")
		return
	}
	logicPlatform, err := utils.PlatformUIToLogic(platform)
	if err != nil {
		fmt.Println("cant convert platform")
		return
	}
	logicGeo, err := utils.GeoUIToLogic(azone)
	if err != nil {
		fmt.Println("cant convert geo")
		return
	}
	route := models.Route{
		LinkID:   linkID,
		LongURL:  url,
		Platform: logicPlatform,
		AZone:    logicGeo,
	}
	err = s.routingServ.CreateRoute(&route)
	if err != nil {
		fmt.Println("cant create route")
		return
	}
	fmt.Printf("Route ID: %d\n", route.ID)
}

func (s *Handler) getLinkRoutesHandler() {
	var linkID models.UUID
	fmt.Printf("Enter link ID: ")
	_, err := fmt.Scanln(&linkID)
	if err != nil {
		fmt.Println("cant read link ID")
		return
	}
	routes, err := s.routingServ.GetLinkRoutes(linkID)
	if err != nil {
		fmt.Println("cant get routes")
		return
	}
	for _, route := range routes {
		fmt.Printf("Route ID: %d\n", route.ID)
		fmt.Printf("URL: %s\n", route.LongURL)
		fmt.Printf("Platform: %s\n", route.Platform)
		fmt.Printf("Availability zone: %s\n\n", route.AZone)
	}
}

func (s *Handler) getRouteStatsHandler() {
	var routeID models.UUID
	fmt.Printf("Enter route ID: ")
	_, err := fmt.Scanln(&routeID)
	if err != nil {
		fmt.Println("cant read route ID")
		return
	}
	stats, err := s.statsServ.GetRouteStats(routeID)
	if err != nil {
		fmt.Println("cant get stats")
		return
	}
	for _, stat := range stats {
		fmt.Printf("Stat ID: %d\n", stat.ID)
		fmt.Printf("Stat UserAgent: %s\n", stat.UserAgent)
		fmt.Printf("Stat IP: %s\n\n", stat.UserIP)
	}
}
