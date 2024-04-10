package unit_tests_test

import (
	"testing"
	"time"

	"surlit/internal/logic/controllers"
	logicErrors "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mock "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestLinkService_CreateLink_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	userID := models.NewUUID()
	projectID := models.NewUUID()
	link := &models.Link{
		ID:         models.NewUUID(),
		URLShort:   "abc",
		DefLongURL: "long_abc",
		CreatedAt:  time.Now(),
	}

	projectRepo.EXPECT().GetUserRoleInProject(userID, projectID).Return(models.ProjectOwner, nil)
	linkRepo.EXPECT().InsertLink(userID, projectID, link).Return(nil)

	err := linkService.CreateLink(userID, projectID, link)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestLinkService_CreateLink_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	userID := models.NewUUID()
	projectID := models.NewUUID()
	link := &models.Link{
		ID:         models.NewUUID(),
		URLShort:   "abc",
		DefLongURL: "long_abc",
		CreatedAt:  time.Now(),
	}

	projectRepo.EXPECT().GetUserRoleInProject(userID, projectID).Return(models.ProjectMember, nil)

	err := linkService.CreateLink(userID, projectID, link)
	errMsg := logicErrors.ErrCantAddNewLinkInProject.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestLinkService_GetLinkInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	linkID := models.NewUUID()
	link := &models.Link{
		ID:         linkID,
		URLShort:   "abc",
		DefLongURL: "long_abc",
		CreatedAt:  time.Now(),
	}

	linkRepo.EXPECT().GetLinkInfo(linkID).Return(link, nil)

	gotLink, err := linkService.GetLinkInfo(linkID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotLink == nil || gotLink.ID != link.ID {
		t.Errorf("Wrong link %+v, got: %+v", link, gotLink)
	}
}

func TestLinkService_GetLinkInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	linkID := models.NewUUID()

	linkRepo.EXPECT().GetLinkInfo(linkID).Return(nil, logicErrors.ErrCantRequest)

	_, err := linkService.GetLinkInfo(linkID)
	errMsg := "link get info: " + logicErrors.ErrCantRequest.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestLinkService_UpdateLinkInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	linkID := models.NewUUID()
	newShortURL := "abc"
	newLongURL := "long_abc"
	link := &models.UpdateLink{
		URLShort:   &newShortURL,
		DefLongURL: &newLongURL,
	}

	linkRepo.EXPECT().UpdateLinkInfo(linkID, link).Return(nil)

	err := linkService.UpdateLinkInfo(linkID, link)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestLinkService_UpdateLinkInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	linkRepo := mock.NewMockLinkRepository(ctrl)
	projectRepo := mock.NewMockProjectRepository(ctrl)

	linkService := controllers.NewLinkService(linkRepo, projectRepo)

	linkID := models.NewUUID()
	newShortURL := "abc"
	newLongURL := "long_abc"
	link := &models.UpdateLink{
		URLShort:   &newShortURL,
		DefLongURL: &newLongURL,
	}

	linkRepo.EXPECT().UpdateLinkInfo(linkID, link).Return(logicErrors.ErrCantUpdate)

	err := linkService.UpdateLinkInfo(linkID, link)
	errMsg := "link update info: " + logicErrors.ErrCantUpdate.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}
