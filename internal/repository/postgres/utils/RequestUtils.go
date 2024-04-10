package utils

import (
	"fmt"
	"surlit/internal/logic/models"
	repoModels "surlit/internal/repository/postgres/models"
)

func RouteFilterConstruct(linkID models.UUID, filter *models.RouteFilter) (*repoModels.RequestRouteFilter, error) {
	repoLinkID, err := LinkIDLogicToRepo(linkID)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to repo: %w", err)
	}
	repoPlatform, err := PlatformLogicToRepo(filter.Platform)
	if err != nil {
		return nil, fmt.Errorf("cant convert platform to repo: %w", err)
	}
	repoAZone, err := AZoneLogicToRepo(filter.AZone)
	if err != nil {
		return nil, fmt.Errorf("cant convert azone to repo: %w", err)
	}
	return &repoModels.RequestRouteFilter{
		LinkID:     repoLinkID,
		PlatformID: repoPlatform,
		AZoneID:    repoAZone,
	}, nil
}
