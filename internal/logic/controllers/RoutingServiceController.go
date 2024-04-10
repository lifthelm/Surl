package controllers

import (
	"fmt"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type RoutingService struct {
	routeRepo interfaces.RoutingRepository
}

var _ interfaces.RoutingService = (*RoutingService)(nil)

func NewRoutingService(rr interfaces.RoutingRepository) *RoutingService {
	return &RoutingService{
		routeRepo: rr,
	}
}

func (r *RoutingService) CreateRoute(route *models.Route) error {
	err := r.routeRepo.InsertRoute(route)
	if err != nil {
		return fmt.Errorf("route create: %w", err)
	}
	return nil
}

func (r *RoutingService) GetLinkRoutes(linkID models.UUID) ([]*models.Route, error) {
	routes, err := r.routeRepo.GetLinkRoutes(linkID)
	if err != nil {
		return nil, fmt.Errorf("route get link routes: %w", err)
	}
	return routes, nil
}

func (r *RoutingService) GetRouteInfo(id models.UUID) (*models.Route, error) {
	route, err := r.routeRepo.GetRouteInfo(id)
	if err != nil {
		return nil, fmt.Errorf("route get route info: %w", err)
	}
	return route, nil
}

func (r *RoutingService) UpdateRouteInfo(id models.UUID, route *models.UpdateRoute) error {
	err := r.routeRepo.UpdateRouteInfo(id, route)
	if err != nil {
		return fmt.Errorf("route update info: %w", err)
	}
	return nil
}
