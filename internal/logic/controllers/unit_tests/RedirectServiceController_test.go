package unit_tests_test

import (
	"testing"
	"time"

	"surlit/internal/logic/controllers"
	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mockip2geo "surlit/internal/mock/ip_to_geo"
	mockplatformdeterm "surlit/internal/mock/platform_determ"
	mockrepo "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestRedirectService_FindRedirect_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mockrepo.NewMockLinkRepository(ctrl)
	routeRepo := mockrepo.NewMockRoutingRepository(ctrl)
	ipToGeo := mockip2geo.NewMockIPToGeo(ctrl)
	platformDeterm := mockplatformdeterm.NewMockPlatformDeterminator(ctrl)
	redirectService := controllers.NewRedirectService(linkRepo, routeRepo, ipToGeo, platformDeterm)

	shortURL := "abc"
	correctRedirectURL := "long_abc"

	recData := models.RequestData{
		UserAgent: "dummy agent",
		IP:        "0.0.0.0",
	}
	link := models.Link{
		ID:         models.NewUUID(),
		URLShort:   shortURL,
		DefLongURL: correctRedirectURL,
		CreatedAt:  time.Time{},
	}
	filter := models.RouteFilter{
		Platform: models.PlatformGeneral,
		AZone:    models.GeoGeneral,
	}
	routes := []*models.Route{
		{
			ID:       models.NewUUID(),
			LongURL:  correctRedirectURL,
			Platform: models.PlatformGeneral,
			AZone:    models.GeoGeneral,
		},
	}

	linkRepo.EXPECT().FindLinkByToken(shortURL).Return(&link, nil)
	ipToGeo.EXPECT().DetermineUserGeo(recData.IP).Return(filter.AZone, nil)
	platformDeterm.EXPECT().DeterminePlatform(recData.UserAgent).Return(filter.Platform, nil)
	routeRepo.EXPECT().FindRoutes(link.ID, &filter).Return(routes, nil)

	gotURL, err := redirectService.FindRedirect(shortURL, &recData)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotURL != correctRedirectURL {
		t.Errorf("Expected URL: \"%s\"; got: \"%q\"", correctRedirectURL, gotURL)
	}
}

func TestRedirectService_FindRedirect_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mockrepo.NewMockLinkRepository(ctrl)
	routeRepo := mockrepo.NewMockRoutingRepository(ctrl)
	ipToGeo := mockip2geo.NewMockIPToGeo(ctrl)
	platformDeterm := mockplatformdeterm.NewMockPlatformDeterminator(ctrl)
	redirectService := controllers.NewRedirectService(linkRepo, routeRepo, ipToGeo, platformDeterm)

	shortURL := "abc"

	recData := models.RequestData{
		UserAgent: "dummy agent",
		IP:        "0.0.0.0",
	}

	linkRepo.EXPECT().FindLinkByToken(shortURL).Return(nil, errors2.ErrCantFind)

	gotURL, err := redirectService.FindRedirect(shortURL, &recData)
	expectedErrMsg := "redirect find: " + errors2.ErrCantFind.Error()
	if err == nil || err.Error() != expectedErrMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", expectedErrMsg, err)
	}
	if gotURL != "" {
		t.Errorf("Expected \"\" redirect URL, got: \"%s\"", gotURL)
	}
}
