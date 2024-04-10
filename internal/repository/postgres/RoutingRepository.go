package postgres

import (
	"fmt"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
	"surlit/internal/repository/postgres/errors"
	repoModels "surlit/internal/repository/postgres/models"
	"surlit/internal/repository/postgres/utils"

	"gorm.io/gorm"
)

type RoutingRepository struct {
	db *gorm.DB
}

var _ interfaces.RoutingRepository = (*RoutingRepository)(nil)

func NewRoutingRepository(db *gorm.DB) *RoutingRepository {
	return &RoutingRepository{
		db: db,
	}
}

func (r RoutingRepository) InsertRoute(route *models.Route) error {
	repoLinkID, err := utils.LinkIDLogicToRepo(route.LinkID)
	if err != nil {
		return fmt.Errorf("cant convert link id to repo: %w", err)
	}
	repoPlatformID, err := utils.PlatformLogicToRepo(route.Platform)
	if err != nil {
		return fmt.Errorf("cant convert platform id to repo: %w", err)
	}
	repoAZoneID, err := utils.AZoneLogicToRepo(route.AZone)
	if err != nil {
		return fmt.Errorf("cant convert azone id to repo: %w", err)
	}
	repoLink := repoModels.LinkRoute{
		LinkID:     repoLinkID,
		PlatformID: repoPlatformID,
		AZoneID:    repoAZoneID,
		LongURL:    route.LongURL,
	}
	res := r.db.Table("link_routes").Create(&repoLink)
	if res.Error != nil {
		return fmt.Errorf("cant insert: %w", res.Error)
	}
	return nil
}

func (r RoutingRepository) GetLinkRoutes(linkID models.UUID) ([]*models.Route, error) {
	repoLinkID, err := utils.LinkIDLogicToRepo(linkID)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to repo: %w", err)
	}
	var repoRoutes []repoModels.LinkRoute
	res := r.db.Model(repoModels.LinkRoute{LinkID: repoLinkID}).Find(&repoRoutes)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request link route: %w", res.Error)
	}
	logicRoutes := make([]*models.Route, 0, len(repoRoutes))
	for _, repoRoute := range repoRoutes {
		logicRoute, err := utils.RouteRepoToLogic(&repoRoute)
		if err != nil {
			return nil, fmt.Errorf("cant convert route to logic: %w", err)
		}
		logicRoutes = append(logicRoutes, logicRoute)
	}
	return logicRoutes, nil
}

func (r RoutingRepository) FindRoutes(
	linkID models.UUID,
	filter *models.RouteFilter,
) ([]*models.Route, error) {
	requestRoutes, err := utils.RouteFilterConstruct(linkID, filter)
	if err != nil {
		return nil, fmt.Errorf("cant construct route filter: %w", err)
	}
	var repoRoutes []repoModels.LinkRoute
	res := r.db.Table("link_routes").Where(requestRoutes).Find(&repoRoutes)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request link routes: %w", err)
	}
	if len(repoRoutes) == 0 {
		return nil, nil
	}
	logicRoutes := make([]*models.Route, 0, len(repoRoutes))
	for _, repoRoute := range repoRoutes {
		logicRoute, err := utils.RouteRepoToLogic(&repoRoute)
		if err != nil {
			return nil, fmt.Errorf("cant convert route to logic: %w", err)
		}
		logicRoutes = append(logicRoutes, logicRoute)
	}
	return logicRoutes, nil
}

func (r RoutingRepository) GetRouteInfo(uuid models.UUID) (*models.Route, error) {
	repoRouteID, err := utils.RouteIDLogicToRepo(uuid)
	if err != nil {
		return nil, fmt.Errorf("cant convert route id to repo: %w", err)
	}
	var (
		repoRoute repoModels.LinkRoute
		count     int64
	)
	res := r.db.Model(repoModels.LinkRoute{RouteID: repoRouteID}).Take(&repoRoute).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request link route: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicRoute, err := utils.RouteRepoToLogic(&repoRoute)
	if err != nil {
		return nil, fmt.Errorf("cant convert route to logic: %w", err)
	}
	return logicRoute, nil
}

func (r RoutingRepository) UpdateRouteInfo(uuid models.UUID, route *models.UpdateRoute) error {
	repoUpdateRoute, err := utils.UpdateRouteLogicToRepo(route)
	if err != nil {
		return fmt.Errorf("cant update route convert to repo: %w", err)
	}
	repoRouteID, err := utils.RouteIDLogicToRepo(uuid)
	if err != nil {
		return fmt.Errorf("cant convert route id to repo: %w", err)
	}
	res := r.db.Model(&repoModels.LinkRoute{RouteID: repoRouteID}).Updates(repoUpdateRoute)
	if res.Error != nil {
		return fmt.Errorf("cant update: %w", res.Error)
	}
	return nil
}
