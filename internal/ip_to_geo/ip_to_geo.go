package ip_to_geo

import (
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type IPToGeo struct {
}

var _ interfaces.IPToGeo = (*IPToGeo)(nil)

func NewIPToGeo() *IPToGeo {
	return &IPToGeo{}
}

func (I IPToGeo) DetermineUserGeo(ip string) (models.GeoType, error) {
	// TODO proper implement
	return models.GeoEurope, nil
}
