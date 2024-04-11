package utils

import (
	"fmt"
	"surlit/internal/logic/models"
)

const (
	PlatformAny     = 0
	PlatformIOS     = 1
	PlatformAndroid = 2
	PlatformMacOS   = 3
	PlatformWindows = 4

	AZoneAny        = 0
	AZoneUS         = 1
	AZoneAfrica     = 2
	AZoneWestEurope = 3
	AZoneEastEurope = 4
	AZoneAsia       = 5
	AZoneAustralia  = 6
)

func PlatformLogicToRepo(platform models.PlatformType) (uint, error) {
	switch platform {
	case models.PlatformGeneral:
		return PlatformAny, nil
	case models.PlatformAndroid:
		return PlatformAndroid, nil
	case models.PlatformIOS:
		return PlatformIOS, nil
	default:
		return PlatformAny, fmt.Errorf("cant convert platform to repo")
	}
}

func PlatformRepoToLogic(platform uint) (models.PlatformType, error) {
	switch platform {
	case PlatformAny:
		return models.PlatformGeneral, nil
	case PlatformIOS:
		return models.PlatformIOS, nil
	case PlatformAndroid:
		return models.PlatformAndroid, nil
	case PlatformMacOS, PlatformWindows: // TODO expand logic
		return models.PlatformGeneral, nil
	default:
		return models.PlatformGeneral, fmt.Errorf("cant convert platform to logic")
	}
}

func AZoneLogicToRepo(zone models.GeoType) (uint, error) {
	switch zone {
	case models.GeoGeneral:
		return AZoneAny, nil
	case models.GeoAmerica:
		return AZoneUS, nil
	case models.GeoAfrica:
		return AZoneAfrica, nil
	case models.GeoEurope:
		return AZoneWestEurope, nil
	case models.GeoAsia:
		return AZoneAsia, nil
	default:
		return AZoneAny, fmt.Errorf("cant convert azone to repo")
	}
}

func AZoneRepoToLogic(zone uint) (models.GeoType, error) {
	switch zone {
	case AZoneAny:
		return models.GeoGeneral, nil
	case AZoneUS:
		return models.GeoAmerica, nil
	case AZoneAfrica:
		return models.GeoAfrica, nil
	case AZoneWestEurope, AZoneEastEurope:
		return models.GeoEurope, nil
	case AZoneAsia:
		return models.GeoAsia, nil
	case AZoneAustralia: // TODO expand logic
		return models.GeoGeneral, nil
	default:
		return models.GeoGeneral, fmt.Errorf("cant convert azone to logic")
	}
}
