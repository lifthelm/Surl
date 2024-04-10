package utils

import (
	"surlit/internal/logic/models"
	repoModels "surlit/internal/repository/postgres/models"
)

func UpdateUserLogicToRepo(user *models.UpdateUser) (*repoModels.User, error) {
	var repoModel repoModels.User
	if user.UserName != nil {
		repoModel.UserName = *user.UserName
	}
	if user.Password != nil {
		repoModel.Password = *user.Password
	}
	if user.Email != nil {
		repoModel.Email = *user.Email
	}
	if user.UserRole != nil {
		repoRole, err := UserRoleLogicToRepo(*user.UserRole)
		if err != nil {
			return nil, err
		}
		repoModel.UserRole = repoRole
	}
	return &repoModel, nil
}

func UpdateProjectLogicToRepo(project *models.UpdateProject) (*repoModels.Project, error) {
	var repoModel repoModels.Project
	if project.Name != nil {
		repoModel.Name = *project.Name
	}
	if project.Description != nil {
		repoModel.Description = *project.Description
	}
	return &repoModel, nil
}

func UpdateLinkLogicToRepo(project *models.UpdateLink) (*repoModels.Link, error) {
	var repoModel repoModels.Link
	if project.URLShort != nil {
		repoModel.ShortURL = *project.URLShort
	}
	if project.DefLongURL != nil {
		repoModel.DefLongURL = *project.DefLongURL
	}
	return &repoModel, nil
}

func UpdateRouteLogicToRepo(route *models.UpdateRoute) (*repoModels.LinkRoute, error) {
	var repoModel repoModels.LinkRoute
	if route.LongURL != nil {
		repoModel.LongURL = *route.LongURL
	}
	if route.AZone != nil {
		repoAZone, err := AZoneLogicToRepo(*route.AZone)
		if err != nil {
			return nil, err
		}
		repoModel.AZoneID = repoAZone
	}
	if route.Platform != nil {
		repoPlatform, err := PlatformLogicToRepo(*route.Platform)
		if err != nil {
			return nil, err
		}
		repoModel.PlatformID = repoPlatform
	}
	return &repoModel, nil
}
