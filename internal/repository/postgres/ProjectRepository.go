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

type ProjectRepository struct {
	db *gorm.DB
}

var _ interfaces.ProjectRepository = (*ProjectRepository)(nil)

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{
		db: db,
	}
}

func (p ProjectRepository) InsertProject(userID models.UUID, project *models.Project) error {
	ownerID, err := utils.UserIDLogicToRepo(userID)
	if err != nil {
		return fmt.Errorf("cant convert user id to repo: %w", err)
	}
	err = p.db.Transaction(func(tx *gorm.DB) error {
		repoProject := repoModels.Project{
			Name:        project.Name,
			Description: project.Description,
		}
		res := tx.Table("projects").Create(&repoProject)
		if res.Error != nil {
			return fmt.Errorf("cant create project in projects: %w", res.Error)
		}
		repoUserProjectRel := repoModels.UserProjectRelation{
			UserID:    ownerID,
			ProjectID: repoProject.ProjectID,
			UserRole:  repoModels.ProjectOwner,
		}
		res = tx.Table("user_project_relations").Create(&repoUserProjectRel)
		if res.Error != nil {
			return fmt.Errorf("cant create user project relation: %w", res.Error)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("cant insert project, transaction: %w", err)
	}
	return nil
}

func (p ProjectRepository) GetUserProjects(userID models.UUID) ([]*models.Project, error) {
	repoUserID, err := utils.UserIDLogicToRepo(userID)
	if err != nil {
		return nil, fmt.Errorf("cant convert userid to repo: %w", err)
	}
	var repoProjects []repoModels.Project
	queryUserProjectIDS := p.db.Table("user_project_relations upr").
		Select("upr.project_id").
		Where(repoModels.UserProjectRelation{UserID: repoUserID})
	res := p.db.Table("projects p").
		Select("p.*").
		Joins("left join (?) upr on upr.project_id = p.project_id", queryUserProjectIDS).
		Find(&repoProjects)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request user project relation: %w", res.Error)
	}
	logicProjects := make([]*models.Project, 0, len(repoProjects))
	for _, repoProject := range repoProjects {
		logicProject, err := utils.ProjectRepoToLogic(&repoProject)
		if err != nil {
			return nil, fmt.Errorf("cant convert project to logic: %w", err)
		}
		logicProjects = append(logicProjects, logicProject)
	}
	return logicProjects, nil
}

func (p ProjectRepository) GetProjectInfo(uuid models.UUID) (*models.Project, error) {
	repoID, err := utils.ProjectIDLogicToRepo(uuid)
	if err != nil {
		return nil, fmt.Errorf("cant convert project id to repo: %w", err)
	}
	var (
		repoProject repoModels.Project
		count       int64
	)
	res := p.db.Model(&repoModels.Project{ProjectID: repoID}).
		Select("projects.project_id, projects.name, projects.description").Take(&repoProject).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("cant request project: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicID, err := utils.ProjectIDRepoToLogic(repoProject.ProjectID)
	if err != nil {
		return nil, fmt.Errorf("cant convert project id to logic: %w", err)
	}
	logicProject := models.Project{
		ID:          logicID,
		Name:        repoProject.Name,
		Description: repoProject.Description,
	}
	return &logicProject, nil
}

func (p ProjectRepository) GetUserRoleInProject(
	userID,
	projectID models.UUID,
) (models.UserProjectRole, error) {
	repoUserID, err := utils.UserIDLogicToRepo(userID)
	if err != nil {
		return 0, fmt.Errorf("cant convert user id to repo: %w", err)
	}
	repoProjectID, err := utils.ProjectIDLogicToRepo(projectID)
	if err != nil {
		return 0, fmt.Errorf("cant convert project id to repo: %w", err)
	}
	var (
		repoProject repoModels.RequestProjectRole
		count       int64
	)
	res := p.db.Model(&repoModels.Project{ProjectID: repoProjectID}).
		Select("projects.project_id, projects.name, projects.description, upr.user_role").
		Joins("join user_project_relations upr on upr.project_id = projects.project_id").
		Where("user_id = ?", repoUserID).Take(&repoProject).Count(&count)
	if res.Error != nil {
		return 0, fmt.Errorf("cant request project: %w", res.Error)
	}
	if count != 1 {
		return 0, errors.ErrCantMatchRecord
	}
	userRole, err := utils.UserProjectRoleRepoToLogic(repoProject.UserRole)
	if err != nil {
		return 0, fmt.Errorf("cant convert user role to logic: %w", err)
	}
	return userRole, nil
}

func (p ProjectRepository) UpdateProjectInfo(
	uuid models.UUID,
	project *models.UpdateProject,
) error {
	repoUpdateProject, err := utils.UpdateProjectLogicToRepo(project)
	if err != nil {
		return fmt.Errorf("cant update project convert to repo: %w", err)
	}
	repoProjectID, err := utils.ProjectIDLogicToRepo(uuid)
	if err != nil {
		return fmt.Errorf("cant convert project_id to repo: %w", err)
	}
	res := p.db.Model(&repoModels.Project{ProjectID: repoProjectID}).Updates(repoUpdateProject)
	if res.Error != nil {
		return fmt.Errorf("cant update: %w", res.Error)
	}
	return nil
}
