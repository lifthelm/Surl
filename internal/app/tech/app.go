package tech

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"surlit/internal/config"
	"surlit/internal/ip_to_geo"
	"surlit/internal/logic/controllers"
	"surlit/internal/logic/models"
	"surlit/internal/platform_determ"
	postgresRepo "surlit/internal/repository/postgres"
)

type App struct {
	db      *gorm.DB
	service *Handler
	config  *config.Config
	user    *models.User
	handler *Handler
}

func (a *App) InitService() {
	userRepo := postgresRepo.NewUserRepository(a.db)
	projectRepo := postgresRepo.NewProjectRepository(a.db)
	linkRepo := postgresRepo.NewLinkRepository(a.db)
	routingRepo := postgresRepo.NewRoutingRepository(a.db)
	statsRepo := postgresRepo.NewStatsRepository(a.db)

	ipToGeo := ip_to_geo.NewIPToGeo()
	platformDeterm := platform_determ.NewPlatformDeterm()

	userServ := controllers.NewUserService(userRepo)
	projectServ := controllers.NewProjectService(projectRepo)
	linkServ := controllers.NewLinkService(linkRepo, projectRepo)
	routingServ := controllers.NewRoutingService(routingRepo)
	statsServ := controllers.NewStatService(statsRepo)
	redirectServ := controllers.NewRedirectService(linkRepo, routingRepo, ipToGeo, platformDeterm, statsRepo)

	dbCheck := postgresRepo.NewHealthCheck(a.db)
	healthCheckServ := controllers.NewHealthCheckService().
		RegisterChecks(dbCheck).
		SetTimePeriod(a.config.HealthCheckPeriod).
		SetTimeout(a.config.HealthCheckTimeout)

	service := NewService(userServ, projectServ, linkServ, routingServ, redirectServ, statsServ, healthCheckServ)

	a.service = service
	a.user = nil
}

func NewApp(conf *config.Config) (*App, error) {
	db, err := gorm.Open(postgres.Open(conf.DBConnectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return &App{
		db:      db,
		service: nil,
		config:  conf,
		user:    nil,
		handler: nil,
	}, nil
}
