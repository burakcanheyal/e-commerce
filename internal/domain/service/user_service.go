package service

import (
	"attempt4/internal"
	dto2 "attempt4/internal/domain/dto"
	entity2 "attempt4/internal/domain/entity"
	enum2 "attempt4/internal/domain/enum"
	"attempt4/platform/hash"
	repository2 "attempt4/platform/postgres/repository"
	"fmt"
	"time"
)

type UserService struct {
	userRepository   repository2.UserRepository
	roleRepository   repository2.RoleRepository
	walletRepository repository2.WalletRepository
}

func NewUserService(
	userRepository repository2.UserRepository,
	roleRepository repository2.RoleRepository,
	walletRepository repository2.WalletRepository) UserService {
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

	user.Status = enum2.UserDeletedStatus

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

func (u *UserService) GetUserByUsername(username string) (dto2.UserDto, error) {
	userDto := dto2.UserDto{}
	user, err := u.userRepository.GetByName(username)
	if err != nil {
		return userDto, err
	}
	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}

	userDto = dto2.UserDto{
		Username:  user.Username,
		Email:     user.Email,
		Name:      user.Name,
		Surname:   user.Surname,
		Status:    user.Status,
		BirthDate: *user.BirthDate,
	}

	return userDto, nil
}

func (u *UserService) GetUserById(id int32) (dto2.UserDto, error) {
	userDto := dto2.UserDto{}
	user, err := u.userRepository.GetById(id)
	if err != nil {
		return userDto, err
	}
	if user.Id == 0 {
		return userDto, internal.UserNotFound
	}

	userDto = dto2.UserDto{
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

func (u *UserService) UpdateUser(id int32, userDto dto2.UserDto) error {
	user, err := u.userRepository.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return err
		}
		return internal.UserNotFound
	}
	updatedTime := time.Now()
	user = entity2.User{
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

func (u *UserService) UpdateUserPassword(id int32, userDto dto2.UserUpdatePasswordDto) error {
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

	entityUser := entity2.User{
		Id:            user.Id,
		Username:      user.Username,
		Password:      userDto.NewPassword,
		Email:         user.Email,
		Name:          user.Name,
		Surname:       user.Surname,
		Status:        enum2.UserActiveStatus,
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

func (u *UserService) CreateUser(userDto dto2.UserDto) error {
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

	user = entity2.User{
		Username:      userDto.Username,
		Password:      encryptedPassword,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        enum2.UserPassiveStatus,
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

	key := entity2.Role{
		UserId: user.Id,
		Rol:    enum2.RoleUser,
	}

	key, err = u.roleRepository.Create(key)
	if err != nil {
		return err
	}
	if key.Id == 0 {
		return internal.RoleNotCreated
	}

	wallet := entity2.Wallet{
		UserId:    user.Id,
		Balance:   0,
		Status:    enum2.WalletPassive,
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

func (u *UserService) ActivateUser(codeDto dto2.UserUpdateCodeDto) error {
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

	user.Status = enum2.UserActiveStatus
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
