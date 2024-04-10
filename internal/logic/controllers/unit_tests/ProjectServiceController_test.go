package unit_tests_test

import (
	"testing"

	"surlit/internal/logic/controllers"
	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mock "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestProjectService_CreateProject_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	project := &models.Project{
		ID:          models.NewUUID(),
		Name:        "Project ABC",
		Description: "Description ABC",
		OwnerID:     models.NewUUID(),
	}

	projectRepo.EXPECT().InsertProject(project).Return(nil)

	err := projectService.CreateProject(project)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestProjectService_CreateProject_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	project := &models.Project{
		ID:          models.NewUUID(),
		Name:        "Project ABC",
		Description: "Description ABC",
		OwnerID:     models.NewUUID(),
	}

	projectRepo.EXPECT().InsertProject(project).Return(errors2.ErrCantInsert)

	err := projectService.CreateProject(project)
	errMsg := "project create: " + errors2.ErrCantInsert.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestProjectService_GetProjectInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	projectID := models.NewUUID()
	project := &models.Project{
		ID:          projectID,
		Name:        "Project ABC",
		Description: "Description ABC",
	}

	projectRepo.EXPECT().GetProjectInfo(projectID).Return(project, nil)

	gotProject, err := projectService.GetProjectInfo(projectID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotProject == nil || gotProject.ID != project.ID {
		t.Errorf("Expected project %+v; got: %+v", project, gotProject)
	}
}

func TestProjectService_GetProjectInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	projectID := models.NewUUID()

	projectRepo.EXPECT().GetProjectInfo(projectID).Return(nil, errors2.ErrCantUpdate)

	_, err := projectService.GetProjectInfo(projectID)
	errMsg := "project get info: " + errors2.ErrCantUpdate.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestProjectService_UpdateProjectInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	projectID := models.NewUUID()
	newName := "Project ABC"
	newDescription := "Description ABC"
	project := &models.UpdateProject{
		Name:        &newName,
		Description: &newDescription,
	}

	projectRepo.EXPECT().UpdateProjectInfo(projectID, project).Return(nil)

	err := projectService.UpdateProjectInfo(projectID, project)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestProjectService_UpdateProjectInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	projectRepo := mock.NewMockProjectRepository(ctrl)
	projectService := controllers.NewProjectService(projectRepo)

	projectID := models.NewUUID()
	newName := "Project ABC"
	newDescription := "Description ABC"
	project := &models.UpdateProject{
		Name:        &newName,
		Description: &newDescription,
	}

	projectRepo.EXPECT().UpdateProjectInfo(projectID, project).Return(errors2.ErrCantUpdate)

	err := projectService.UpdateProjectInfo(projectID, project)
	errMsg := "project update info: " + errors2.ErrCantUpdate.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}
