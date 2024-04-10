package unit_tests_test

import (
	"testing"
	"time"

	"surlit/internal/logic/controllers"
	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mock "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestStatService_CreateStat_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	stat := &models.Stat{
		ID:           models.NewUUID(),
		CreatedAt:    time.Now(),
		UserPlatform: models.PlatformGeneral,
		UserAZone:    models.GeoGeneral,
		UserAgent:    "UserAgent ABC",
	}

	statsRepo.EXPECT().InsertStat(stat).Return(nil)

	err := statService.CreateStat(stat)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestStatService_CreateStat_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	stat := &models.Stat{
		ID:           models.NewUUID(),
		CreatedAt:    time.Now(),
		UserPlatform: models.PlatformGeneral,
		UserAZone:    models.GeoGeneral,
		UserAgent:    "UserAgent ABC",
	}

	tmpErr := errors2.ErrCantInsert
	statsRepo.EXPECT().InsertStat(stat).Return(tmpErr)

	err := statService.CreateStat(stat)
	if err == nil {
		t.Error("Expected nil error")
	}
}

func TestStatService_GetRouteStats_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	routeID := models.NewUUID()
	stats := []*models.Stat{
		{
			ID:           models.NewUUID(),
			CreatedAt:    time.Now(),
			UserPlatform: models.PlatformGeneral,
			UserAZone:    models.GeoGeneral,
			UserAgent:    "UserAgent ABC",
		},
		{
			ID:           models.NewUUID(),
			CreatedAt:    time.Now(),
			UserPlatform: models.PlatformGeneral,
			UserAZone:    models.GeoGeneral,
			UserAgent:    "UserAgent EFG",
		},
	}

	statsRepo.EXPECT().GetRouteStats(routeID).Return(stats, nil)

	gotStats, err := statService.GetRouteStats(routeID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if len(gotStats) != len(stats) {
		t.Errorf("Expected \"%d\"; got \"%d\"", len(stats), len(gotStats))
	}
}

func TestStatService_GetRouteStats_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	routeID := models.NewUUID()

	tmpErr := errors2.ErrCantRequest
	statsRepo.EXPECT().GetRouteStats(routeID).Return(nil, tmpErr)

	_, err := statService.GetRouteStats(routeID)
	if err == nil {
		t.Error("Expected nil error")
	}
}

func TestStatService_GetStatInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	statID := models.NewUUID()
	stat := &models.Stat{
		ID:           models.NewUUID(),
		CreatedAt:    time.Now(),
		UserPlatform: models.PlatformGeneral,
		UserAZone:    models.GeoGeneral,
		UserAgent:    "UserAgent ABC",
	}

	statsRepo.EXPECT().GetStatInfo(statID).Return(stat, nil)

	gotStat, err := statService.GetStatInfo(statID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotStat == nil || gotStat.ID != stat.ID || gotStat.UserAgent != stat.UserAgent {
		t.Errorf("Expected %+v; got %+v", stat, gotStat)
	}
}

func TestStatService_GetStatInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	statsRepo := mock.NewMockStatsRepository(ctrl)
	statService := controllers.NewStatService(statsRepo)

	statID := models.NewUUID()

	tmpErr := errors2.ErrCantRequest
	statsRepo.EXPECT().GetStatInfo(statID).Return(nil, tmpErr)

	_, err := statService.GetStatInfo(statID)
	if err == nil {
		t.Error("Expected nil error")
	}
}
