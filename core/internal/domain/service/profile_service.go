package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/hash"
	"attempt4/core/platform/jwt"
	"attempt4/core/platform/postgres/repository"
	"fmt"
	"time"
)

type UserService struct {
	userRepository repository.UserRepository
	secret         string
}

func NewUserService(userRepository repository.UserRepository, secret string) UserService {
	p := UserService{userRepository, secret}
	return p
}

func (p *UserService) DeleteUser(tokenString string) error {
	username, err := jwt.ExtractUsernameFromToken(tokenString, p.secret)
	if err != nil {
		return err
	}

	user, _ := p.userRepository.GetByName(username)
	user.Status = enum.UserDeletedStatus

	err = p.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
func (p *UserService) GetUserByTokenString(tokenString string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	username, err := jwt.ExtractUsernameFromToken(tokenString, p.secret)

	if err != nil {
		return userDto, err
	}
	user, _ := p.userRepository.GetByName(username)

	userDto = dto.UserDto{
		Username:  user.Username,
		Password:  "********",
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: user.BirthDate,
	}

	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}

	return userDto, nil
}
func (p *UserService) GetUserByUsername(username string) (dto.UserDto, error) {
	user, _ := p.userRepository.GetByName(username)
	userDto := dto.UserDto{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: user.BirthDate,
	}

	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}
	return userDto, nil
}
func (p *UserService) GetUserById(id int32) (dto.UserDto, error) {
	user, _ := p.userRepository.GetById(id)

	userDto := dto.UserDto{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: user.BirthDate,
	}

	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}
	return userDto, nil
}
func (p *UserService) UpdateUser(userDto dto.UserDto) error {
	user, err := p.userRepository.GetByName(userDto.Username)
	if user.Id == 0 {
		return internal.DBNotFound
	}

	user = entity.User{
		Id:            user.Id,
		Username:      user.Username,
		Password:      user.Password,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        1,
		Code:          user.Code,
		CodeExpiredAt: user.CodeExpiredAt,
		BirthDate:     userDto.BirthDate,
	}

	err = p.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
func (p *UserService) UpdateUserPassword(userDto dto.UserUpdatePasswordDto) error {
	user, err := p.userRepository.GetByName(userDto.UserName)
	if user.Id == 0 {
		return internal.UserNotFound
	}

	err = hash.CompareEncryptedPasswords(user.Password, userDto.Password)
	if err != nil {
		return err
	}

	entityUser := entity.User{
		Id:            user.Id,
		Username:      user.Username,
		Password:      userDto.NewPassword,
		Email:         user.Email,
		Name:          user.Name,
		Surname:       user.Surname,
		Status:        1,
		Code:          user.Code,
		CodeExpiredAt: user.CodeExpiredAt,
		BirthDate:     user.BirthDate,
	}

	err = p.userRepository.Update(entityUser)
	if err != nil {
		return err
	}

	return nil
}
func (p *UserService) CreateUser(userDto dto.UserDto) error {
	user, err := p.userRepository.GetByName(userDto.Username)
	if user.Id != 0 {
		return internal.UserExist
	}

	encryptedPassword, err := hash.EncryptPassword(userDto.Password)
	if err != nil {
		return err
	}

	//Todo: Cross table dan role alınacak
	user = entity.User{
		Username:      userDto.Username,
		Password:      encryptedPassword,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Role:          enum.RoleUser,
		Status:        enum.UserPassiveStatus,
		Code:          generateCode(),
		CodeExpiredAt: time.Now().Add(time.Second * 300),
		BirthDate:     userDto.BirthDate,
	}

	user, err = p.userRepository.Create(user)
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}
func (p *UserService) ActivateUser(codeDto dto.UserUpdateCodeDto) error {
	user, err := p.userRepository.GetByName(codeDto.Username)
	if user.Id == 0 {
		return internal.UserNotFound
	}

	if user.CodeExpiredAt.Before(time.Now()) {
		user.Code = generateCode()
		user.CodeExpiredAt = time.Now().Add(time.Second * 300)
		err = p.userRepository.Update(user)

		if err != nil {
			return err
		}
		return internal.ExceedVerifyCode
	}

	if codeDto.Code != user.Code {
		return internal.FailInVerify
	}

	user.Status = enum.UserActiveStatus
	err = p.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
func generateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}
