package postgres

import (
	"fmt"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
	"surlit/internal/repository/postgres/errors"
	repoModels "surlit/internal/repository/postgres/models"
	"surlit/internal/repository/postgres/utils"

	"gorm.io/gorm"
)

type StatsRepository struct {
	db *gorm.DB
}

var _ interfaces.StatsRepository = (*StatsRepository)(nil)

func NewStatsRepository(db *gorm.DB) *StatsRepository {
	return &StatsRepository{
		db: db,
	}
}

func (s StatsRepository) InsertStat(stat *models.Stat) error {
	//repoPlatformID, err := utils.PlatformLogicToRepo(stat.UserPlatform) // TODO
	//if err != nil {
	//	return fmt.Errorf("cant convert platform id to repo: %w", err)
	//}
	//repoAZoneID, err := utils.AZoneLogicToRepo(stat.UserAZone) // TODO
	//if err != nil {
	//	return fmt.Errorf("cant convert azone id to repo: %w", err)
	//}
	repoRouteID, err := utils.RouteIDLogicToRepo(stat.RouteID)
	if err != nil {
		return fmt.Errorf("cant convert route id to repo: %w", err)
	}
	repoStat := repoModels.ClickStat{
		RouteID:       repoRouteID,
		ClickDateTime: stat.CreatedAt,
		UserAgent:     stat.UserAgent,
		UserIPAddress: stat.UserIP,
	}
	res := s.db.Table("clicks_stats").Create(&repoStat)
	if res.Error != nil {
		return fmt.Errorf("cant insert click stat: %w", res.Error)
	}
	return nil
}

func (s StatsRepository) GetRouteStats(RouteID models.UUID) ([]*models.Stat, error) {
	repoRouteID, err := utils.RouteIDLogicToRepo(RouteID)
	if err != nil {
		return nil, fmt.Errorf("cant convert route id to repo: %w", err)
	}
	var repoStats []repoModels.ClickStat
	res := s.db.Model(repoModels.ClickStat{RouteID: repoRouteID}).Find(&repoStats)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request stats: %w", res.Error)
	}
	if len(repoStats) == 0 {
		return nil, errors.ErrNoRecordFound
	}
	logicStats := make([]*models.Stat, 0, len(repoStats))
	for _, repoStat := range repoStats {
		logicStat, err := utils.StatRepoToLogic(&repoStat)
		if err != nil {
			return nil, fmt.Errorf("cant convert stat to logic: %w", err)
		}
		logicStats = append(logicStats, logicStat)
	}
	return logicStats, nil
}

func (s StatsRepository) GetStatInfo(uuid models.UUID) (*models.Stat, error) {
	repoRecordID, err := utils.ClickStatIDLogicToRepo(uuid)
	if err != nil {
		return nil, fmt.Errorf("cant convert record id to repo: %w", err)
	}
	var (
		repoStat repoModels.ClickStat
		count    int64
	)
	res := s.db.Model(repoModels.ClickStat{RecordID: repoRecordID}).Take(&repoStat).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request click stat: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicStat, err := utils.StatRepoToLogic(&repoStat)
	if err != nil {
		return nil, fmt.Errorf("cant convert stat to logic: %w", err)
	}
	return logicStat, nil
}
