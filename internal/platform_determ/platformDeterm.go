package platform_determ

import (
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
	"surlit/internal/repository/postgres/utils"
)

type PlatformDeterm struct {
}

var _ interfaces.PlatformDeterminator = (*PlatformDeterm)(nil)

func NewPlatformDeterm() *PlatformDeterm {
	return &PlatformDeterm{}
}

func (p PlatformDeterm) DeterminePlatform(userAgent string) (models.PlatformType, error) {
	//TODO implement me
	return utils.PlatformAndroid, nil
}
