package integrationTests_test

import (
	"fmt"
	"surlit/internal/config"
	"testing"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"surlit/internal/logic/controllers"
	"surlit/internal/logic/models"
	postgresRepo "surlit/internal/repository/postgres"
)

// TODO две очереди добавить, которые обсуждали

func TestUserService_UserInsertion_Positive(t *testing.T) {
	conf, err := config.GetConfigYML("./../../../../config.yml")
	if err != nil {
		t.Errorf("cant get conf")
		return
	}
	db, err := gorm.Open(postgres.Open(conf.DBConnectionString), &gorm.Config{})
	if err != nil {
		t.Errorf("cant connect to DB")
		return
	}
	defer func(db *gorm.DB) {
		dbInstance, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("cant get db instance %w", err))
		}
		err = dbInstance.Close()
		if err != nil {
			panic(fmt.Errorf("cant close db connection %w", err))
		}
	}(db)

	userRepo := postgresRepo.NewUserRepository(db)

	user := models.User{
		UserName: "user abc",
		Password: "password abc",
		Email:    "abc@abc.abc",
		UserRole: models.UserRegular,
	}

	err = userRepo.InsertUser(&user)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}

func TestUserService_Registration(t *testing.T) {
	conf, err := config.GetConfigYML("./config.yml")
	if err != nil {
		t.Errorf("cant get conf")
		return
	}
	db, err := gorm.Open(postgres.Open(conf.DBConnectionString), &gorm.Config{})
	if err != nil {
		t.Errorf("cant connect to DB")
		return
	}
	defer func(db *gorm.DB) {
		dbInstance, err := db.DB()
		if err != nil {
			panic(fmt.Errorf("cant get db instance %w", err))
		}
		err = dbInstance.Close()
		if err != nil {
			panic(fmt.Errorf("cant close db connection %w", err))
		}
	}(db)

	userRepo := postgresRepo.NewUserRepository(db)

	userService := controllers.NewUserService(userRepo)

	user := models.User{
		UserName: "user ghc",
		Password: "password abc",
		Email:    "abc@abc.abc",
		UserRole: models.UserRegular,
	}

	err = userService.Registration(&user)
	if err != nil {
		t.Errorf("Unexpected error: \"%v\"", err)
	}
}
