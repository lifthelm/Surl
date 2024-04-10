package interfaces

import "surlit/internal/logic/models"

type PlatformDeterminator interface {
	DeterminePlatform(userAgent string) (models.PlatformType, error)
}
