package utils

import (
	"surlit/internal/app/tech/errors"
	"surlit/internal/logic/models"
)

func UserRoleLogicToUI(role models.UserRole) (string, error) {
	switch role {
	case models.UserRegular:
		return "User", nil
	case models.UserAdmin:
		return "Admin", nil
	default:
		return "-", errors.ErrUnknownRole
	}
}
