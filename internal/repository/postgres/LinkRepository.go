package postgres

import (
	"fmt"
	"time"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
	"surlit/internal/repository/postgres/errors"
	repoModels "surlit/internal/repository/postgres/models"
	"surlit/internal/repository/postgres/utils"

	"gorm.io/gorm"
)

type LinkRepository struct {
	db *gorm.DB
}

var _ interfaces.LinkRepository = (*LinkRepository)(nil)

func NewLinkRepository(db *gorm.DB) *LinkRepository {
	return &LinkRepository{
		db: db,
	}
}

func (l LinkRepository) InsertLink(userID, projectID models.UUID, link *models.Link) error {
	repoUserID, err := utils.UserIDLogicToRepo(userID)
	if err != nil {
		return fmt.Errorf("cant convert user id to repo: %w", err)
	}
	repoProjectID, err := utils.ProjectIDLogicToRepo(projectID)
	if err != nil {
		return fmt.Errorf("cant convert project id to repo: %w", err)
	}
	err = l.db.Transaction(func(tx *gorm.DB) error {
		repoLink := repoModels.Link{
			DefLongURL: link.DefLongURL,
			ShortURL:   link.URLShort,
			CreateDate: time.Now(),
		}
		res := tx.Table("links").Create(&repoLink)
		if res.Error != nil {
			return fmt.Errorf("cant create link in links: %w", res.Error)
		}
		repoUserLinkRel := repoModels.UserLinkRelation{
			UserID:   repoUserID,
			LinkID:   repoLink.LinkID,
			UserRole: repoModels.LinkOwner,
		}
		res = tx.Table("user_link_relations").Create(&repoUserLinkRel)
		if res.Error != nil {
			return fmt.Errorf("cant create user link relation: %w", res.Error)
		}
		repoProjectLinkRel := repoModels.ProjectLinkRelation{
			ProjectID: repoProjectID,
			LinkID:    repoLink.LinkID,
		}
		res = tx.Table("project_link_relations").Create(&repoProjectLinkRel)
		if res.Error != nil {
			return fmt.Errorf("cant create project link relation: %w", res.Error)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("cant insert project, transaction: %w", err)
	}
	return nil
}

func (l LinkRepository) GetProjectLinks(projectID models.UUID) ([]*models.Link, error) {
	repoProjectID, err := utils.ProjectIDLogicToRepo(projectID)
	if err != nil {
		return nil, fmt.Errorf("cant convert project id to repo: %w", err)
	}
	var repoLinks []repoModels.Link
	queryProjectLinksIDS := l.db.Table("project_link_relations plr").
		Select("plr.link_id").
		Where(repoModels.ProjectLinkRelation{ProjectID: repoProjectID})
	res := l.db.Table("links l").
		Select("l.*").
		Joins("left join (?) plr on plr.link_id = l.link_id", queryProjectLinksIDS).
		Find(&repoLinks)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request user project relation: %w", res.Error)
	}
	logicLinks := make([]*models.Link, 0, len(repoLinks))
	for _, repoLink := range repoLinks {
		logicLink, err := utils.LinkRepoToLogic(&repoLink)
		if err != nil {
			return nil, fmt.Errorf("cant convert link to logic: %w", err)
		}
		logicLinks = append(logicLinks, logicLink)
	}
	return logicLinks, nil
}

func (l LinkRepository) GetUserLinks(userID models.UUID) ([]*models.Link, error) {
	repoUserID, err := utils.LinkIDLogicToRepo(userID)
	if err != nil {
		return nil, fmt.Errorf("cant convert user id to repo: %w", err)
	}
	var repoLinks []repoModels.Link
	queryProjectLinksIDS := l.db.Table("user_link_relations ulr").
		Select("ulr.link_id").
		Where(repoModels.UserLinkRelation{UserID: repoUserID})
	res := l.db.Table("links l").
		Select("l.*").
		Joins("inner join (?) ulr on ulr.link_id = l.link_id", queryProjectLinksIDS).
		Find(&repoLinks)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request link project relation: %w", res.Error)
	}
	logicLinks := make([]*models.Link, 0, len(repoLinks))
	for _, repoLink := range repoLinks {
		logicLink, err := utils.LinkRepoToLogic(&repoLink)
		if err != nil {
			return nil, fmt.Errorf("cant convert link to logic: %w", err)
		}
		logicLinks = append(logicLinks, logicLink)
	}
	return logicLinks, nil
}

func (l LinkRepository) GetLinkInfo(uuid models.UUID) (*models.Link, error) {
	repoLinkID, err := utils.LinkIDLogicToRepo(uuid)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to repo: %w", err)
	}
	var (
		repoLink repoModels.Link
		count    int64
	)
	res := l.db.Model(&repoModels.Link{LinkID: repoLinkID}).Take(&repoLink).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("can request link: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicLink := models.Link{
		ID:         uuid,
		DefLongURL: repoLink.DefLongURL,
		URLShort:   repoLink.ShortURL,
		CreatedAt:  repoLink.CreateDate,
	}
	return &logicLink, nil
}

func (l LinkRepository) FindLinkByToken(token string) (*models.Link, error) {
	var (
		repoLink repoModels.Link
		count    int64
	)
	res := l.db.Table("links").Where(&repoModels.Link{ShortURL: token}).Take(&repoLink).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("can request link: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicLinkID, err := utils.LinkIDRepoToLogic(repoLink.LinkID)
	if err != nil {
		return nil, fmt.Errorf("cant convert link id to logic: %w", err)
	}
	logicLink := models.Link{
		ID:         logicLinkID,
		DefLongURL: repoLink.DefLongURL,
		URLShort:   repoLink.ShortURL,
		CreatedAt:  repoLink.CreateDate,
	}
	return &logicLink, nil
}

func (l LinkRepository) UpdateLinkInfo(uuid models.UUID, Link *models.UpdateLink) error {
	repoUpdateLink, err := utils.UpdateLinkLogicToRepo(Link)
	if err != nil {
		return fmt.Errorf("cant update link convert to repo: %w", err)
	}
	repoLinkID, err := utils.LinkIDLogicToRepo(uuid)
	if err != nil {
		return fmt.Errorf("cant convert link_id to repo: %w", err)
	}
	res := l.db.Model(&repoModels.Link{LinkID: repoLinkID}).Updates(repoUpdateLink)
	if res.Error != nil {
		return fmt.Errorf("cant update: %w", res.Error)
	}
	return nil
}
