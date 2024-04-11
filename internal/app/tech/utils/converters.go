package utils

import (
	"surlit/internal/app/tech/errors"
	"surlit/internal/logic/models"
)

func PlatformUIToLogic(platform string) (models.PlatformType, error) {
	switch platform {
	case "android":
		return models.PlatformAndroid, nil
	case "ios":
		return models.PlatformIOS, nil
	case "general":

		return models.PlatformGeneral, nil
	default:
		return models.PlatformGeneral, errors.ErrUnknownPlatform
	}
}

func GeoUIToLogic(geo string) (models.GeoType, error) {
	switch geo {
	case "america":
		return models.GeoAmerica, nil
	case "asia":
		return models.GeoAsia, nil
	case "europe":

		return models.GeoEurope, nil
	case "africa":

		return models.GeoAfrica, nil
	case "general":

		return models.GeoGeneral, nil
	default:
		return models.GeoGeneral, errors.ErrUnknownGeo
	}
}

func UUIDUIToLogic(uuid int) (models.UUID, error) {
	return models.UUID(uuid), nil
}
