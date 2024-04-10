package controllers

import (
	"fmt"

	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type ProjectService struct {
	projectRepo interfaces.ProjectRepository
}

var _ interfaces.ProjectService = (*ProjectService)(nil)

func NewProjectService(pr interfaces.ProjectRepository) *ProjectService {
	return &ProjectService{
		projectRepo: pr,
	}
}

func (p *ProjectService) CreateProject(userID models.UUID, project *models.Project) error {
	err := p.projectRepo.InsertProject(userID, project)
	if err != nil {
		return fmt.Errorf("project create: %w", err)
	}
	return nil
}

func (p *ProjectService) GetUserProjects(userID models.UUID) ([]*models.Project, error) {
	projects, err := p.projectRepo.GetUserProjects(userID)
	if err != nil {
		return nil, fmt.Errorf("get user projects: %w", err)
	}
	return projects, nil
}

func (p *ProjectService) GetProjectInfo(id models.UUID) (*models.Project, error) {
	project, err := p.projectRepo.GetProjectInfo(id)
	if err != nil {
		return nil, fmt.Errorf("project get info: %w", err)
	}
	return project, nil
}

func (p *ProjectService) UpdateProjectInfo(id models.UUID, project *models.UpdateProject) error {
	err := p.projectRepo.UpdateProjectInfo(id, project)
	if err != nil {
		return fmt.Errorf("project update info: %w", err)
	}
	return nil
}
