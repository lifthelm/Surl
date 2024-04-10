package interfaces

import "surlit/internal/logic/models"

type IPToGeo interface {
	DetermineUserGeo(ip string) (models.GeoType, error)
}
