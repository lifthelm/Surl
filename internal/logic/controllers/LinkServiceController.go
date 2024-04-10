package controllers

import (
	"fmt"

	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type LinkService struct {
	linkRepo    interfaces.LinkRepository
	projectRepo interfaces.ProjectRepository
}

var _ interfaces.LinkService = (*LinkService)(nil)

func NewLinkService(lr interfaces.LinkRepository, pr interfaces.ProjectRepository) *LinkService {
	return &LinkService{
		linkRepo:    lr,
		projectRepo: pr,
	}
}

func (l *LinkService) CreateLink(userID, projectID models.UUID, link *models.Link) error {
	userRole, err := l.projectRepo.GetUserRoleInProject(userID, projectID)
	if err != nil {
		return fmt.Errorf("cant get user role in project: %w", err)
	}
	if userRole != models.ProjectOwner {
		return errors2.ErrCantAddNewLinkInProject
	}
	err = l.linkRepo.InsertLink(userID, projectID, link)
	if err != nil {
		return fmt.Errorf("link create: %w", err)
	}
	return nil
}

func (l *LinkService) GetProjectLinks(projectID models.UUID) ([]*models.Link, error) {
	links, err := l.linkRepo.GetProjectLinks(projectID)
	if err != nil {
		return nil, fmt.Errorf("get project links: %w", err)
	}
	return links, nil
}

func (l *LinkService) GetUserLinks(userID models.UUID) ([]*models.Link, error) {
	links, err := l.linkRepo.GetUserLinks(userID)
	if err != nil {
		return nil, fmt.Errorf("get project links: %w", err)
	}
	return links, nil
}

func (l *LinkService) GetLinkInfo(id models.UUID) (*models.Link, error) {
	link, err := l.linkRepo.GetLinkInfo(id)
	if err != nil {
		return nil, fmt.Errorf("link get info: %w", err)
	}
	return link, nil
}

func (l *LinkService) UpdateLinkInfo(id models.UUID, link *models.UpdateLink) error {
	err := l.linkRepo.UpdateLinkInfo(id, link)
	if err != nil {
		return fmt.Errorf("link update info: %w", err)
	}
	return nil
}
