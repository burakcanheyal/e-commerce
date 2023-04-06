package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/hash"
	"attempt4/core/platform/jwt"
	"attempt4/core/platform/postgres/repository"
)

type Authentication struct {
	UserRepository repository.UserRepository
	Secret         string
	Secret2        string
}

func NewAuthentication(userRepos repository.UserRepository, secret string, secret2 string) Authentication {
	a := Authentication{userRepos, secret, secret2}
	return a
}
func (p *Authentication) Login(userDto dto.AuthDto) error {
	user, err := p.UserRepository.GetByName(userDto.Username)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return internal.UserNotFound
	}

	if user.Status == enum.UserDeletedStatus {
		return internal.DeletedUser
	}
	if user.Status == enum.UserPassiveStatus {
		return internal.PassiveUser
	}

	err = hash.CompareEncryptedPasswords(user.Password, userDto.Password)
	if err != nil {
		return err
	}

	return nil
}

func (p *Authentication) GetUserByTokenString(tokenString string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	username, err := jwt.ExtractUsernameFromToken(tokenString, p.Secret)
	if err != nil {
		return userDto, err
	}

	user, err := p.UserRepository.GetByName(username)
	if err != nil {
		return userDto, err
	}
	if user.Id == 0 {
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

func (p *Authentication) GenerateAccessToken(Username string) (string, error) {
	accessToken, err := jwt.GenerateAccessToken(Username, p.Secret)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

func (p *Authentication) GenerateRefreshToken(Username string) (string, error) {
	refreshToken, err := jwt.GenerateRefreshToken(Username, p.Secret2)
	if err != nil {
		return "", err
	}
	return refreshToken, nil
}

func (p *Authentication) ValidateAccessToken(tokenString string) error {
	err := jwt.ValidateToken(tokenString, p.Secret)
	if err != nil {
		return err
	}
	return nil
}

func (p *Authentication) ValidateRefreshToken(tokenString string) error {
	err := jwt.ValidateToken(tokenString, p.Secret2)
	if err != nil {
		return err
	}
	return nil
}
