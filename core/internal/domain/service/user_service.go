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
	userRepository   repository.UserRepository
	roleRepository   repository.RoleRepository
	walletRepository repository.WalletRepository
}

func NewUserService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	walletRepository repository.WalletRepository) UserService {
	u := UserService{
		userRepository,
		roleRepository,
		walletRepository,
	}
	return u
}

func (u *UserService) DeleteUser(id int32) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return internal.UserNotFound
	}

	user.Status = enum.UserDeletedStatus

	deletedTime := time.Now()
	user.DeletedAt = &deletedTime

	err = u.userRepository.Delete(user)
	if err != nil {
		return err
	}

	wallet, err := u.walletRepository.GetByUserId(user.Id)
	if err != nil {
		return err
	}
	if wallet.Id == 0 {

		return internal.WalletNotFound
	}

	wallet.DeletedAt = &deletedTime

	err = u.walletRepository.Delete(wallet)
	if err != nil {
		return err
	}

	role, err := u.roleRepository.GetByUserId(user.Id)
	if err != nil {
		return err
	}
	if role.Id == 0 {
		return internal.RoleNotFound
	}

	err = u.roleRepository.Delete(role)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserByUsername(username string) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	user, err := u.userRepository.GetByName(username)
	if err != nil {
		return userDto, err
	}
	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}

	userDto = dto.UserDto{
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: *user.BirthDate,
	}

	return userDto, nil
}

func (u *UserService) GetUserById(id int32) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return userDto, err
	}
	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}

	userDto = dto.UserDto{
		Username:  user.Username,
		Password:  user.Password,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: *user.BirthDate,
	}

	return userDto, nil
}

func (u *UserService) UpdateUser(id int32, userDto dto.UserDto) error {
	user, err := u.userRepository.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return err
		}
		return internal.UserNotFound
	}
	updatedTime := time.Now()
	user = entity.User{
		Id:            user.Id,
		Username:      user.Username,
		Password:      user.Password,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        userDto.Status,
		Code:          user.Code,
		CodeExpiredAt: user.CodeExpiredAt,
		BirthDate:     &userDto.BirthDate,
		UpdatedAt:     &updatedTime,
	}

	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) UpdateUserPassword(id int32, userDto dto.UserUpdatePasswordDto) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return err
	}
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
		Status:        enum.UserActiveStatus,
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
	if err != nil {
		return err
	}
	if user.Id != 0 {
		return internal.UserExist
	}

	encryptedPassword, err := hash.EncryptPassword(userDto.Password)
	if err != nil {
		return err
	}

	currentTime := time.Now()
	expiredTime := currentTime.Add(time.Second * 300)

	user = entity.User{
		Username:      userDto.Username,
		Password:      encryptedPassword,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        enum.UserPassiveStatus,
		Code:          generateCode(),
		CodeExpiredAt: &expiredTime,
		BirthDate:     &userDto.BirthDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
		DeletedAt:     nil,
	}

	user, err = u.userRepository.Create(user)
	if err != nil {
		return internal.UserNotCreated
	}

	key := entity.Role{
		UserId: user.Id,
		Rol:    enum.RoleUser,
	}

	key, err = u.roleRepository.Create(key)
	if err != nil {
		return err
	}
	if key.Id == 0 {
		return internal.RoleNotCreated
	}

	wallet := entity.Wallet{
		UserId:    user.Id,
		Balance:   0,
		Status:    enum.WalletPassive,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	wallet, err = u.walletRepository.Create(wallet)
	if err != nil {
		return err
	}
	if wallet.Id == 0 {
		return internal.WalletNotCreated
	}
	return nil
}

func (u *UserService) ActivateUser(codeDto dto.UserUpdateCodeDto) error {
	user, err := u.userRepository.GetByName(codeDto.Username)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return internal.UserNotFound
	}

	if user.CodeExpiredAt.Before(time.Now()) {
		user.Code = generateCode()
		expiredCode := time.Now().Add(time.Second * 300)
		user.CodeExpiredAt = &expiredCode

		err = u.userRepository.Update(user)
		if err != nil {
			return err
		}
		return internal.ExceedVerifyCode
	}

	if codeDto.Code != *user.Code {
		return internal.FailInVerify
	}

	user.Status = enum.UserActiveStatus
	err = u.userRepository.Update(user)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserService) GetUserRoleById(id int32) (int, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return 0, err
	}
	if user.Id == 0 {
		return 0, internal.UserNotFound
	}

	rol, err := u.roleRepository.GetByUserId(user.Id)
	if err != nil {
		return 0, err
	}

	return rol.Rol, nil
}

func generateCode() *string {
	code := fmt.Sprint(time.Now().Nanosecond())[:6]
	return &code
}
