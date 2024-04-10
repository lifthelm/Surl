package unit_tests_test

import (
	"fmt"
	"testing"

	"surlit/internal/logic/controllers"
	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mock "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestRoutingService_CreateRoute_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	Route := &models.Route{}

	RoutingRepo.EXPECT().InsertRoute(Route).Return(nil)

	err := routingService.CreateRoute(Route)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestRoutingService_CreateRoute_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	Route := &models.Route{}

	tmpErr := errors2.ErrCantInsert
	RoutingRepo.EXPECT().InsertRoute(Route).Return(tmpErr)

	err := routingService.CreateRoute(Route)
	errMsg := fmt.Sprintf("route create: %v", tmpErr)
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestRoutingService_GetLinkRoutes_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	mockLinkID := models.NewUUID()

	mockRoutes := make([]*models.Route, 4)

	RoutingRepo.EXPECT().GetLinkRoutes(mockLinkID).Return(mockRoutes, nil)

	routes, err := routingService.GetLinkRoutes(mockLinkID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if len(routes) != len(mockRoutes) {
		t.Errorf("Expected routes: %d; got: %d", len(mockRoutes), len(routes))
	}
}

func TestRoutingService_GetLinkRoutes_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	mockLinkID := models.NewUUID()

	tmpErr := errors2.ErrCantRequest
	RoutingRepo.EXPECT().GetLinkRoutes(mockLinkID).Return(nil, tmpErr)

	routes, err := routingService.GetLinkRoutes(mockLinkID)
	errMsg := fmt.Sprintf("route get link routes: %v", tmpErr)
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
	if routes != nil {
		t.Errorf("Expected nil; got: \"%v\"", routes)
	}
}

func TestRoutingService_GetRouteInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	routeID := models.NewUUID()

	Route := &models.Route{}

	RoutingRepo.EXPECT().GetRouteInfo(routeID).Return(Route, nil)

	route, err := routingService.GetRouteInfo(routeID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if route != Route {
		t.Errorf("Expected route: \"%v\"; got: \"%v\"", Route, route)
	}
}

func TestRoutingService_GetRouteInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	routeID := models.NewUUID()

	tmpErr := errors2.ErrCantRequest
	RoutingRepo.EXPECT().GetRouteInfo(routeID).Return(nil, tmpErr)

	route, err := routingService.GetRouteInfo(routeID)
	errMsg := fmt.Sprintf("route get route info: %v", tmpErr)
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
	if route != nil {
		t.Errorf("Expected nil, got: \"%v\"", route)
	}
}

func TestRoutingService_UpdateRouteInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	routeID := models.NewUUID()
	newLongURL := "long abc"
	Route := &models.UpdateRoute{
		LongURL: &newLongURL,
	}

	RoutingRepo.EXPECT().UpdateRouteInfo(routeID, Route).Return(nil)

	err := routingService.UpdateRouteInfo(routeID, Route)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestRoutingService_UpdateRouteInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	RoutingRepo := mock.NewMockRoutingRepository(ctrl)
	routingService := controllers.NewRoutingService(RoutingRepo)

	routeID := models.NewUUID()
	newLongURL := "long abc"
	Route := &models.UpdateRoute{
		LongURL: &newLongURL,
	}

	tmpErr := errors2.ErrCantUpdate
	RoutingRepo.EXPECT().UpdateRouteInfo(routeID, Route).Return(tmpErr)

	err := routingService.UpdateRouteInfo(routeID, Route)
	errMsg := fmt.Sprintf("route update info: %v", tmpErr)
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}
