package integrationTests_test

import (
	"fmt"
	"testing"

	"surlit/internal/config"
	"surlit/internal/logic/controllers"
	"surlit/internal/logic/models"
	mockinterfaces "surlit/internal/mock/ip_to_geo"
	mockinterfaces2 "surlit/internal/mock/platform_determ"
	postgresRepo "surlit/internal/repository/postgres"

	"go.uber.org/mock/gomock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestRedirect_DefURL(t *testing.T) {
	conf, err := config.GetConfigYML("./../../../../config.yml")
	if err != nil {
		t.Errorf("cant get conf")
		return
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db, err := gorm.Open(postgres.Open(conf.DBConnectionString), &gorm.Config{})
	if err != nil {
		t.Errorf("Unexpected db connection error: \"%v\"", err)
	}
	defer func(db *gorm.DB) {
		dbInstance, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("cant get db instance %w", err))
		}
		err = dbInstance.Close()
		if err != nil {
			panic(fmt.Errorf("cant close db connection %w", err))
		}
	}(db)

	linkRepo := postgresRepo.NewLinkRepository(db)
	routeRepo := postgresRepo.NewRoutingRepository(db)
	statsRepo := postgresRepo.NewStatsRepository(db)

	mockIPToGeo := mockinterfaces.NewMockIPToGeo(ctrl)
	mockPlatformDeterm := mockinterfaces2.NewMockPlatformDeterminator(ctrl)

	mockPlatformDeterm.EXPECT().DeterminePlatform("none").
		Return(models.PlatformGeneral, nil)
	mockIPToGeo.EXPECT().DetermineUserGeo("0.0.0.0").Return(models.GeoGeneral, nil)

	redirectService := controllers.NewRedirectService(linkRepo, routeRepo, mockIPToGeo, mockPlatformDeterm, statsRepo)

	ansLong := "https://www.chandler.info/"

	rec := models.RequestData{
		UserAgent: "none",
		IP:        "0.0.0.0",
	}

	long, err := redirectService.FindRedirect("0VZHNT", &rec)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if long != ansLong {
		t.Errorf("Long url mismatch: \"%s\" != \"%s\"(target)", long, ansLong)
	}
}

func TestRedirect_Route(t *testing.T) {
	conf, err := config.GetConfigYML("./config.yml")
	if err != nil {
		t.Errorf("cant get conf")
		return
	}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	db, err := gorm.Open(postgres.Open(conf.DBConnectionString), &gorm.Config{})
	if err != nil {
		t.Errorf("Unexpected db connection error: \"%v\"", err)
	}
	defer func(db *gorm.DB) {
		dbInstance, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("cant get db instance %w", err))
		}
		err = dbInstance.Close()
		if err != nil {
			panic(fmt.Errorf("cant close db connection %w", err))
		}
	}(db)

	linkRepo := postgresRepo.NewLinkRepository(db)
	routeRepo := postgresRepo.NewRoutingRepository(db)
	statsRepo := postgresRepo.NewStatsRepository(db)

	mockIPToGeo := mockinterfaces.NewMockIPToGeo(ctrl)
	mockPlatformDeterm := mockinterfaces2.NewMockPlatformDeterminator(ctrl)

	mockPlatformDeterm.EXPECT().DeterminePlatform("none").
		Return(models.PlatformAndroid, nil)
	mockIPToGeo.EXPECT().DetermineUserGeo("0.0.0.0").Return(models.GeoAsia, nil)

	redirectService := controllers.NewRedirectService(linkRepo, routeRepo, mockIPToGeo, mockPlatformDeterm, statsRepo)

	ansLong := "http://boyd.com/"

	rec := models.RequestData{
		UserAgent: "none",
		IP:        "0.0.0.0",
	}

	long, err := redirectService.FindRedirect("GAAFO9", &rec)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if long != ansLong {
		t.Errorf("Long url mismatch: \"%s\" != \"%s\"(target)", long, ansLong)
	}
}
