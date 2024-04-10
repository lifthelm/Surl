package utils

import (
	"surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	repoModels "surlit/internal/repository/postgres/models"
)

func UserRoleLogicToRepo(role models.UserRole) (string, error) {
	switch role {
	case models.UserRegular:
		return repoModels.UserRegular, nil
	case models.UserAdmin:
		return repoModels.UserAdmin, nil
	default:
		return "", errors.ErrUnknownUserRole
	}
}

func UserRoleRepoToLogic(role string) (models.UserRole, error) {
	switch role {
	case repoModels.UserRegular:
		return models.UserRegular, nil
	case repoModels.UserAdmin:
		return models.UserAdmin, nil
	default:
		return 0, errors.ErrCantConvertToLogic
	}
}

func UserProjectRoleLogicToRepo(role models.UserProjectRole) (string, error) {
	switch role {
	case models.ProjectOwner:
		return repoModels.ProjectOwner, nil
	case models.ProjectMember:
		return repoModels.ProjectMember, nil
	default:
		return "", errors.ErrUnknownUserRole
	}
}

func UserProjectRoleRepoToLogic(role string) (models.UserProjectRole, error) {
	switch role {
	case repoModels.ProjectOwner:
		return models.ProjectOwner, nil
	case repoModels.ProjectMember:
		return models.ProjectMember, nil
	default:
		return 0, errors.ErrCantConvertToLogic
	}
}
