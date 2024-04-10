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

type UserRepository struct {
	db *gorm.DB
}

var _ interfaces.UserRepository = (*UserRepository)(nil)

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u UserRepository) InsertUser(user *models.User) error {
	userRoleRepo, err := utils.UserRoleLogicToRepo(user.UserRole)
	if err != nil {
		return fmt.Errorf("cant convert model to repo: %w", err)
	}
	userRepo := repoModels.User{
		UserName:         user.UserName,
		Email:            user.Email,
		Password:         user.Password,
		UserRole:         userRoleRepo,
		RegistrationDate: time.Now(),
	}
	res := u.db.Table("users").Create(&userRepo)
	if res.Error != nil {
		return fmt.Errorf("cant insert: %w", res.Error)
	}
	return nil
}

func (u UserRepository) Authorize(login, password string) (*models.User, error) {
	var (
		userRepo repoModels.User
		count    int64
	)
	res := u.db.Table("users").
		Where("username = ? AND password = ?", login, password).Take(&userRepo).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("cant run request: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicRole, err := utils.UserRoleRepoToLogic(userRepo.UserRole)
	if err != nil {
		return nil, fmt.Errorf("cant convert user role: %w", err)
	}
	logicID, err := utils.UserIDRepoToLogic(userRepo.UserID)
	if err != nil {
		return nil, fmt.Errorf("cant convert user id: %w", err)
	}
	user := models.User{
		ID:        logicID,
		UserName:  userRepo.UserName,
		Email:     userRepo.Email,
		Password:  userRepo.Password,
		UserRole:  logicRole,
		CreatedAt: userRepo.RegistrationDate,
	}
	return &user, nil
}

func (u UserRepository) GetUserInfo(uuid models.UUID) (*models.User, error) {
	var (
		userRepo repoModels.User
		count    int64
	)
	repoID, err := utils.UserIDLogicToRepo(uuid)
	if err != nil {
		return nil, fmt.Errorf("cant convert user id: %w", err)
	}
	res := u.db.Table("users").Where("user_id = ?", repoID).Take(&userRepo).Count(&count)
	if res.Error != nil {
		return nil, fmt.Errorf("cant run request: %w", res.Error)
	}
	if count != 1 {
		return nil, errors.ErrCantMatchRecord
	}
	logicRole, err := utils.UserRoleRepoToLogic(userRepo.UserRole)
	if err != nil {
		return nil, fmt.Errorf("cant convert user role: %w", err)
	}
	logicUser := models.User{
		ID:        uuid,
		UserName:  userRepo.UserName,
		Email:     userRepo.Email,
		Password:  userRepo.Password,
		UserRole:  logicRole,
		CreatedAt: userRepo.RegistrationDate,
	}
	return &logicUser, nil
}

func (u UserRepository) UpdateUserInfo(uuid models.UUID, user *models.UpdateUser) error {
	repoUpdateUser, err := utils.UpdateUserLogicToRepo(user)
	if err != nil {
		return fmt.Errorf("cant update user convert to repo: %w", err)
	}
	repoUserID, err := utils.UserIDLogicToRepo(uuid)
	if err != nil {
		return fmt.Errorf("cant convert userid to repo: %w", err)
	}
	res := u.db.Model(&repoModels.User{UserID: repoUserID}).Updates(repoUpdateUser)
	if res.Error != nil {
		return fmt.Errorf("cant update: %w", res.Error)
	}
	return nil
}
