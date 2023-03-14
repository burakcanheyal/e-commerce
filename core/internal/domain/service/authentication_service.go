package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/hash"
	"attempt4/core/platform/jwt"
)

type Authentication struct {
	UserService UserService
	Secret      string
	Secret2     string
}

func NewAuthentication(userService UserService, secret string, secret2 string) Authentication {
	a := Authentication{userService, secret, secret2}
	return a
}
func (p *Authentication) Login(userDto dto.AuthDto) error {
	user, err := p.UserService.userRepository.GetByName(userDto.Username)
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
