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
	keyRepository  repository.KeyRepository
	secret         string
}

func NewUserService(userRepository repository.UserRepository, keyRepository repository.KeyRepository, secret string) UserService {
	u := UserService{userRepository, keyRepository, secret}
	return u
}

func (u *UserService) DeleteUser(tokenString string) error {
	username, err := jwt.ExtractUsernameFromToken(tokenString, u.secret)
	if err != nil {
		return err
	}

	user, _ := u.userRepository.GetByName(username)
	user.Status = enum.UserDeletedStatus

	err = u.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserService) GetUserByTokenString(tokenString string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	username, err := jwt.ExtractUsernameFromToken(tokenString, u.secret)

	if err != nil {
		return userDto, err
	}
	user, _ := u.userRepository.GetByName(username)

	userDto = dto.UserDto{
		Username:  user.Username,
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
func (u *UserService) GetUserByUsername(username string) (dto.UserDto, error) {
	user, _ := u.userRepository.GetByName(username)
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
func (u *UserService) GetUserById(id int32) (dto.UserDto, error) {
	user, _ := u.userRepository.GetById(id)

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
func (u *UserService) UpdateUser(userDto dto.UserDto) error {
	user, err := u.userRepository.GetByName(userDto.Username)
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

	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}
func (u *UserService) UpdateUserPassword(userDto dto.UserUpdatePasswordDto) error {
	user, err := u.userRepository.GetByName(userDto.UserName)
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

	err = u.userRepository.Update(entityUser)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserService) CreateUser(userDto dto.UserDto) error {
	user, err := u.userRepository.GetByName(userDto.Username)
	if user.Id != 0 {
		return internal.UserExist
	}

	encryptedPassword, err := hash.EncryptPassword(userDto.Password)
	if err != nil {
		return err
	}

	user = entity.User{
		Username:      userDto.Username,
		Password:      encryptedPassword,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        enum.UserPassiveStatus,
		Code:          generateCode(),
		CodeExpiredAt: time.Now().Add(time.Second * 300),
		BirthDate:     userDto.BirthDate,
	}

	user, err = u.userRepository.Create(user)
	if err != nil {
		return internal.UserNotCreated
	}

	key := entity.Key{
		UserId: user.Id,
		Rol:    enum.RoleUser,
	}

	key, err = u.keyRepository.Create(key)
	if key.KeyId == 0 {
		return internal.KeyNotCreated
	}

	return nil
}
func (u *UserService) ActivateUser(codeDto dto.UserUpdateCodeDto) error {
	user, err := u.userRepository.GetByName(codeDto.Username)
	if user.Id == 0 {
		return internal.UserNotFound
	}

	if user.CodeExpiredAt.Before(time.Now()) {
		user.Code = generateCode()
		user.CodeExpiredAt = time.Now().Add(time.Second * 300)
		err = u.userRepository.Update(user)

		if err != nil {
			return err
		}
		return internal.ExceedVerifyCode
	}

	if codeDto.Code != user.Code {
		return internal.FailInVerify
	}

	user.Status = enum.UserActiveStatus
	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserService) GetUserRoleByTokenString(tokenString string) (int, error) {
	username, err := jwt.ExtractUsernameFromToken(tokenString, u.secret)
	if err != nil {
		return 0, err
	}

	user, err := u.userRepository.GetByName(username)
	if err != nil {
		return 0, err
	}

	if user.Status != enum.UserActiveStatus {
		return 0, internal.UserUnactivated
	}
	rol, err := u.keyRepository.GetByUserId(user.Id)
	if err != nil {
		return 0, err
	}

	return rol.Rol, nil
}
func generateCode() string {
	return fmt.Sprint(time.Now().Nanosecond())[:6]
}