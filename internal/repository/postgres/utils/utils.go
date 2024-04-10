package utils

import (
	"fmt"
	"surlit/internal/logic/models"
	repoModels "surlit/internal/repository/postgres/models"
)

func RouteRepoToLogic(route *repoModels.LinkRoute) (*models.Route, error) {
	logicRouteID, err := RouteIDRepoToLogic(route.RouteID)
	if err != nil {
		return nil, fmt.Errorf("cant convert route id to logic: %w", err)
	}
	logicLinkID, err := LinkIDRepoToLogic(route.LinkID)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to logic: %w", err)
	}
	logicPlatform, err := PlatformRepoToLogic(route.PlatformID)
	if err != nil {
		return nil, fmt.Errorf("cant convert platform id to logic: %w", err)
	}
	logicAZone, err := AZoneRepoToLogic(route.AZoneID)
	if err != nil {
		return nil, fmt.Errorf("cant convert azone id to logic: %w", err)
	}
	return &models.Route{
		LinkID:   logicLinkID,
		ID:       logicRouteID,
		Platform: logicPlatform,
		AZone:    logicAZone,
		LongURL:  route.LongURL,
	}, nil
}

func StatRepoToLogic(stat *repoModels.ClickStat) (*models.Stat, error) {
	logicRecordID, err := ClickStatIDRepoToLogic(stat.RecordID)
	if err != nil {
		return nil, fmt.Errorf("cant convert record id to logic: %w", err)
	}
	logicRouteID, err := RouteIDRepoToLogic(stat.RouteID)
	if err != nil {
		return nil, fmt.Errorf("cant convert route id to logic: %w", err)
	}
	// TODO User country + platform
	return &models.Stat{
		ID:        logicRecordID,
		RouteID:   logicRouteID,
		UserIP:    stat.UserIPAddress,
		UserAgent: stat.UserAgent,
		CreatedAt: stat.ClickDateTime,
	}, nil
}

func ProjectRepoToLogic(project *repoModels.Project) (*models.Project, error) {
	logicProjectID, err := ProjectIDRepoToLogic(project.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("cant convert project id to logic: %w", err)
	}
	return &models.Project{
		Name:        project.Name,
		Description: project.Description,
		ID:          logicProjectID,
	}, nil
}

func LinkRepoToLogic(link *repoModels.Link) (*models.Link, error) {
	logicLinkID, err := LinkIDRepoToLogic(link.LinkID)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to logic: %w", err)
	}
	return &models.Link{
		ID:         logicLinkID,
		URLShort:   link.ShortURL,
		DefLongURL: link.DefLongURL,
		CreatedAt:  link.CreateDate,
	}, nil
}
