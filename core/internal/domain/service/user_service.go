package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/hash"
	"attempt4/core/platform/postgres/repository"
	"fmt"
	"time"
)

type UserService struct {
	userRepository repository.UserRepository
	keyRepository  repository.KeyRepository
}

func NewUserService(userRepository repository.UserRepository, keyRepository repository.KeyRepository) UserService {
	u := UserService{
		userRepository,
		keyRepository,
	}
	return u
}

func (u *UserService) DeleteUser(username string) error {
	user, err := u.userRepository.GetByName(username)
	if user.Id == 0 {
		if err != nil {
			return err
		}
		return internal.UserNotFound
	}

	user.Status = enum.UserDeletedStatus

	err = u.userRepository.Delete(user)
	if err != nil {
		return err
	}

	return nil
}
func (u *UserService) GetUserByUsername(username string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	user, err := u.userRepository.GetByName(username)
	if user.Id == 0 {
		if err != nil {
			return userDto, err
		}
		return userDto, internal.UserNotFound
	}

	userDto = dto.UserDto{
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: user.BirthDate,
	}

	return userDto, nil
}
func (u *UserService) GetUserById(id int32) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	user, err := u.userRepository.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return userDto, err
		}
		return userDto, internal.UserNotFound
	}

	userDto = dto.UserDto{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: user.BirthDate,
	}

	return userDto, nil
}
func (u *UserService) UpdateUser(userDto dto.UserDto) error {
	user, err := u.userRepository.GetByName(userDto.Username)
	if user.Id == 0 {
		if err != nil {
			return err
		}
		return internal.UserNotFound
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
		if err != nil {
			return err
		}
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
		if err != nil {
			return err
		}
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
		if err != nil {
			return err
		}
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
func (u *UserService) GetUserRoleByUsername(username string) (int, error) {
	user, err := u.userRepository.GetByName(username)
	if user.Id == 0 {
		if err != nil {
			return 0, err
		}
		return 0, internal.UserNotFound
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
