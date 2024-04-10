package unit_tests_test

import (
	"testing"

	"surlit/internal/logic/controllers"
	errors2 "surlit/internal/logic/errors"
	"surlit/internal/logic/models"
	mock "surlit/internal/mock/repository"

	"go.uber.org/mock/gomock"
)

func TestUserService_Registration_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	user := &models.User{
		ID:       models.NewUUID(),
		UserName: "user abc",
		Password: "password abc",
	}

	userRepo.EXPECT().Authorize(user.UserName, user.Password).Return(nil, nil)
	userRepo.EXPECT().InsertUser(user).Return(nil)

	err := userService.Registration(user)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestUserService_Registration_Negative_AlreadyExists(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	user := &models.User{
		ID:       models.NewUUID(),
		UserName: "user abc",
		Password: "password abc",
	}

	userRepo.EXPECT().Authorize(user.UserName, user.Password).Return(user, nil)

	err := userService.Registration(user)
	errMsg := "user registration: " + errors2.ErrAlreadyExist.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestUserService_Registration_Negative_InsertError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	user := &models.User{
		ID:       models.NewUUID(),
		UserName: "user abc",
		Password: "password abc",
	}

	userRepo.EXPECT().Authorize(user.UserName, user.Password).Return(nil, nil)
	userRepo.EXPECT().InsertUser(user).Return(errors2.ErrCantInsert)

	err := userService.Registration(user)
	errMsg := "user registration: cant insert record"
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestUserService_Authorization_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	user := &models.User{
		ID:       models.NewUUID(),
		UserName: "user abc",
		Password: "password abc",
	}

	userRepo.EXPECT().Authorize(user.UserName, user.Password).Return(user, nil)

	gotUser, err := userService.Authorization(user.UserName, user.Password)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotUser == nil || gotUser.ID != user.ID {
		t.Errorf("Expected %+v; got: %+v", user, gotUser)
	}
}

func TestUserService_Authorization_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	userName := "user abc"
	password := "password abc"

	userRepo.EXPECT().Authorize(userName, password).Return(nil, errors2.ErrAuth)

	_, err := userService.Authorization(userName, password)
	errMsg := "user authorization: auth failed"
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestUserService_GetUserInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	userID := models.NewUUID()
	user := &models.User{
		ID:       userID,
		UserName: "user abc",
		Password: "password abc",
	}

	userRepo.EXPECT().GetUserInfo(userID).Return(user, nil)

	gotUser, err := userService.GetUserInfo(userID)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
	if gotUser == nil || gotUser.ID != user.ID {
		t.Errorf("Expected %+v; got: %+v", user, gotUser)
	}
}

func TestUserService_GetUserInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	userID := models.NewUUID()

	userRepo.EXPECT().GetUserInfo(userID).Return(nil, errors2.ErrCantRequest)

	_, err := userService.GetUserInfo(userID)
	errMsg := "user get info: " + errors2.ErrCantRequest.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}

func TestUserService_UpdateUserInfo_Positive(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	userID := models.NewUUID()
	userName := "user abc"
	userPassword := "password abc"
	user := &models.UpdateUser{
		UserName: &userName,
		Password: &userPassword,
	}

	userRepo.EXPECT().UpdateUserInfo(userID, user).Return(nil)

	err := userService.UpdateUserInfo(userID, user)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestUserService_UpdateUserInfo_Negative(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mock.NewMockUserRepository(ctrl)
	userService := controllers.NewUserService(userRepo)

	userID := models.NewUUID()
	userName := "user abc"
	userPassword := "password abc"
	user := &models.UpdateUser{
		UserName: &userName,
		Password: &userPassword,
	}

	userRepo.EXPECT().UpdateUserInfo(userID, user).Return(errors2.ErrCantUpdate)

	err := userService.UpdateUserInfo(userID, user)
	errMsg := "user update info: " + errors2.ErrCantUpdate.Error()
	if err == nil || err.Error() != errMsg {
		t.Errorf("Expected error: \"%s\"; got: \"%v\"", errMsg, err)
	}
}
