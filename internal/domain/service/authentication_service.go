package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
	"attempt4/platform/hash"
	"attempt4/platform/jwt"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"time"
)

type Authentication struct {
	UserRepository repository.UserRepository
	appLogService  app_log.ApplicationLogService
	Secret         string
	Secret2        string
}

func NewAuthentication(
	userRepos repository.UserRepository,
	secret string,
	secret2 string,
	appLogService app_log.ApplicationLogService) Authentication {
	a := Authentication{userRepos, appLogService, secret, secret2}
	return a
}

func (a *Authentication) Login(userDto dto.AuthDto) (dto.Tokens, error) {
	var tokens dto.Tokens
	user, err := a.UserRepository.GetByName(userDto.Username)
	if err != nil {
		zap.Logger.Error(err)
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		return tokens, err
	}
	if user.Id == 0 {
		zap.Logger.Error(internal.UserNotFound)
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		return tokens, internal.UserNotFound
	}

	if user.Status == enum.UserDeletedStatus {
		return tokens, internal.DeletedUser
	}
	if user.Status == enum.UserPassiveStatus {
		return tokens, internal.PassiveUser
	}

	err = hash.CompareEncryptedPasswords(user.Password, userDto.Password)
	if err != nil {
		zap.Logger.Error(err)
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		return tokens, err
	}

	accessToken, err := a.GenerateAccessToken(user.Username)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		return tokens, err
	}

	refreshToken, err := a.GenerateRefreshToken(user.Username)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		return tokens, err
	}

	tokens = dto.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return tokens, nil
}

func (a *Authentication) GetUserByTokenString(tokenString string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	username, err := jwt.ExtractUsernameFromToken(tokenString, a.Secret)
	if err != nil {
		zap.Logger.Error(err)
		return userDto, err
	}

	user, err := a.UserRepository.GetByName(username)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return userDto, err
	}
	if user.Id == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: user.Id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return userDto, internal.UserNotFound
	}

	userDto = dto.UserDto{
		Id:        user.Id,
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: *user.BirthDate,
	}

	return userDto, nil
}

func (a *Authentication) GenerateAccessToken(Username string) (string, error) {
	accessToken, err := jwt.GenerateAccessToken(Username, a.Secret)
	if err != nil {
		zap.Logger.Error(err)
		return "", err
	}
	return accessToken, nil
}

func (a *Authentication) GenerateRefreshToken(Username string) (string, error) {
	refreshToken, err := jwt.GenerateRefreshToken(Username, a.Secret2)
	if err != nil {
		zap.Logger.Error(err)
		return "", err
	}
	return refreshToken, nil
}

func (a *Authentication) ValidateAccessToken(tokenString string) error {
	err := jwt.ValidateToken(tokenString, a.Secret)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	return nil
}

func (a *Authentication) ValidateRefreshToken(tokenString string) error {
	err := jwt.ValidateToken(tokenString, a.Secret2)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	return nil
}
