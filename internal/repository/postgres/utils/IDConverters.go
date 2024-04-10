package utils

import (
	"surlit/internal/logic/models"
)

func UserIDLogicToRepo(uuid models.UUID) (uint, error) {
	return uint(uuid), nil
}

func UserIDRepoToLogic(uuid uint) (models.UUID, error) {
	return models.UUID(uuid), nil
}

func ProjectIDLogicToRepo(uuid models.UUID) (uint, error) {
	return uint(uuid), nil
}

func ProjectIDRepoToLogic(uuid uint) (models.UUID, error) {
	return models.UUID(uuid), nil
}

func LinkIDLogicToRepo(uuid models.UUID) (uint, error) {
	return uint(uuid), nil
}

func LinkIDRepoToLogic(uuid uint) (models.UUID, error) {
	return models.UUID(uuid), nil
}

func RouteIDLogicToRepo(uuid models.UUID) (uint, error) {
	return uint(uuid), nil
}

func RouteIDRepoToLogic(uuid uint) (models.UUID, error) {
	return models.UUID(uuid), nil
}

//func PlatformIDLogicToRepo(uuid models.UUID) (uint, error) {
//	return uint(uuid), nil
//}
//
//func PlatformIDRepoToLogic(uuid uint) (models.UUID, error) {
//	return models.UUID(uuid), nil
//}
//
//func AZoneIDLogicToRepo(uuid models.UUID) (uint, error) {
//	return uint(uuid), nil
//}
//
//func AZoneIDRepoToLogic(uuid uint) (models.UUID, error) {
//	return models.UUID(uuid), nil
//}

func ClickStatIDLogicToRepo(uuid models.UUID) (uint, error) {
	return uint(uuid), nil
}

func ClickStatIDRepoToLogic(uuid uint) (models.UUID, error) {
	return models.UUID(uuid), nil
}
