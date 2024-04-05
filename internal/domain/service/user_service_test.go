package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/platform/app_log"
	"attempt4/platform/postgres"
	"attempt4/platform/postgres/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

var userService UserService
var db *gorm.DB

func init() {
	db = postgres.GetTestDb()

	userRepository := repository.NewUserRepository(db)
	roleRepository := repository.NewRoleRepository(db)
	walletRepository := repository.NewWalletRepository(db)
	appLogRepository := app_log.NewApplicationLogRepository(db)

	appLogService := app_log.NewApplicationLogService(appLogRepository)
	userService = NewUserService(userRepository, roleRepository, walletRepository, appLogService)
}
func TestCreateUserSuccess(t *testing.T) {

	newUser := dto.UserDto{
		Id:        0,
		Username:  "testUser",
		Password:  "testPass",
		Email:     "test@gmail.com",
		Name:      "Test",
		Surname:   "Test",
		Status:    0,
		Phone:     "+905555555555",
		BirthDate: time.Date(0001, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
	}

	err := userService.CreateUser(newUser)
	assert.Equal(t, nil, err)

	db.Rollback()
}
func TestCreateUserExistUserError(t *testing.T) {
	newUser := dto.UserDto{
		Id:        0,
		Username:  "testUser",
		Password:  "testPass",
		Email:     "test@gmail.com",
		Name:      "Test",
		Surname:   "Test",
		Status:    1,
		Phone:     "+905555555555",
		BirthDate: time.Date(0001, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
	}

	err := userService.CreateUser(newUser)
	assert.Equal(t, nil, err)
	err = userService.CreateUser(newUser)
	assert.Equal(t, internal.UserExist, err)

	db.Rollback()
}
func TestCreateUserDeleteSuccess(t *testing.T) {
	newUser := dto.UserDto{
		Id:        0,
		Username:  "testUser",
		Password:  "testPass",
		Email:     "test@gmail.com",
		Name:      "Test",
		Surname:   "Test",
		Status:    1,
		Phone:     "+905555555555",
		BirthDate: time.Date(0001, time.Month(1), 1, 0, 0, 0, 0, time.UTC),
	}

	err := userService.CreateUser(newUser)
	assert.Equal(t, nil, err)
	err = userService.DeleteUser(1)
	assert.Equal(t, nil, err)

	db.Rollback()
}
