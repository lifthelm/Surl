package controllers

import (
	"fmt"

	"surlit/internal/logic/errors"
	"surlit/internal/logic/interfaces"
	"surlit/internal/logic/models"
)

type UserService struct {
	userRep interfaces.UserRepository
}

var _ interfaces.UserService = (*UserService)(nil)

func NewUserService(ur interfaces.UserRepository) *UserService {
	return &UserService{
		userRep: ur,
	}
}

func (u *UserService) Registration(user *models.User) error {
	registeredUser, err := u.userRep.Authorize(user.UserName, user.Password)
	if err != nil {
		return fmt.Errorf("user registration: %w", err)
	}
	if registeredUser != nil {
		return fmt.Errorf("user registration: %w", errors.ErrAlreadyExist)
	}
	err = u.userRep.InsertUser(user)
	if err != nil {
		return fmt.Errorf("user registration: %w", errors.ErrCantInsert)
	}
	return nil
}

func (u *UserService) Authorization(login, password string) (*models.User, error) {
	registeredUser, err := u.userRep.Authorize(login, password)
	if err != nil {
		return nil, fmt.Errorf("user authorization: %w", err)
	}
	return registeredUser, nil
}

func (u *UserService) GetUserInfo(id models.UUID) (*models.User, error) {
	user, err := u.userRep.GetUserInfo(id)
	if err != nil {
		return nil, fmt.Errorf("user get info: %w", err)
	}
	return user, nil
}

func (u *UserService) UpdateUserInfo(id models.UUID, user *models.UpdateUser) error {
	err := u.userRep.UpdateUserInfo(id, user)
	if err != nil {
		return fmt.Errorf("user update info: %w", err)
	}
	return nil
}
