package controllers

import (
	"errors"
	"fmt"

	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type RedirectService struct {
	linksRepo      interfaces.LinkRepository
	routesRepo     interfaces.RoutingRepository
	ipToGeo        interfaces.IPToGeo
	platformDeterm interfaces.PlatformDeterminator
	stats          interfaces.StatsRepository
}

var _ interfaces.RedirectService = (*RedirectService)(nil)

func NewRedirectService(
	lr interfaces.LinkRepository,
	rr interfaces.RoutingRepository,
	ip2geo interfaces.IPToGeo,
	determ interfaces.PlatformDeterminator,
	stats interfaces.StatsRepository,
) *RedirectService {
	return &RedirectService{
		linksRepo:      lr,
		routesRepo:     rr,
		ipToGeo:        ip2geo,
		platformDeterm: determ,
		stats:          stats,
	}
}

func (r RedirectService) BuildFilter(rec *models.RequestData) (*models.RouteFilter, error) {
	var (
		platform models.PlatformType
		aZone    models.GeoType
		err      error
	)

	platform, err = r.platformDeterm.DeterminePlatform(rec.UserAgent)
	if err != nil {
		return nil, err
	}

	aZone, err = r.ipToGeo.DetermineUserGeo(rec.IP)
	if err != nil {
		return nil, err
	}

	return &models.RouteFilter{
		Platform: platform,
		AZone:    aZone,
	}, nil
}

func (r RedirectService) ChooseRoute(
	routes []*models.Route,
	rec *models.RouteFilter,
) (*models.Route, error) {
	var (
		chosenRoute              *models.Route
		chosenWeight, currWeight int
	)
	for i := range len(routes) {
		currWeight = 0
		if routes[i].Platform == rec.Platform {
			currWeight += 2
		} else if routes[i].Platform != models.PlatformGeneral {
			continue
		}
		if routes[i].AZone == rec.AZone {
			currWeight++
		} else if routes[i].AZone != models.GeoGeneral {
			continue
		}
		if currWeight > chosenWeight {
			chosenWeight = currWeight
			chosenRoute = routes[i]
		}
	}
	if chosenWeight == 0 {
		return nil, errors2.ErrNoRoute
	}
	return chosenRoute, nil
}

func (r RedirectService) FindRedirect(shortURL string, rec *models.RequestData) (string, error) {
	// TODO add stat insert in db, maybe non blocking
	link, err := r.linksRepo.FindLinkByToken(shortURL)
	if err != nil {
		return "", fmt.Errorf("redirect find: %w", err)
	}
	filter, err := r.BuildFilter(rec)
	if err != nil {
		return link.DefLongURL, fmt.Errorf("redirect build filter: %w", err)
	}
	routes, err := r.routesRepo.FindRoutes(link.ID, filter)
	if err != nil {
		return link.DefLongURL, fmt.Errorf("redirect find route: %w", err)
	}
	if len(routes) == 0 {
		return link.DefLongURL, nil
	}
	route, err := r.ChooseRoute(routes, filter)
	if err != nil {
		if errors.Is(err, errors2.ErrNoRoute) {
			return link.DefLongURL, nil
		}
		return link.DefLongURL, fmt.Errorf("cant match route: %w", err)
	}
	return route.LongURL, nil
}
