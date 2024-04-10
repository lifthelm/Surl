package controllers

import (
	"fmt"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type StatService struct {
	statRepo interfaces.StatsRepository
}

var _ interfaces.StatService = (*StatService)(nil)

func NewStatService(sr interfaces.StatsRepository) *StatService {
	return &StatService{
		statRepo: sr,
	}
}

func (s *StatService) CreateStat(stat *models.Stat) error {
	err := s.statRepo.InsertStat(stat)
	if err != nil {
		return fmt.Errorf("stat create: %w", err)
	}
	return nil
}

func (s *StatService) GetRouteStats(routeID models.UUID) ([]*models.Stat, error) {
	stats, err := s.statRepo.GetRouteStats(routeID)
	if err != nil {
		return nil, fmt.Errorf("stat get route stats: %w", err)
	}
	return stats, nil
}

func (s *StatService) GetStatInfo(id models.UUID) (*models.Stat, error) {
	stat, err := s.statRepo.GetStatInfo(id)
	if err != nil {
		return nil, fmt.Errorf("stat get stat info: %w", err)
	}
	return stat, nil
}
