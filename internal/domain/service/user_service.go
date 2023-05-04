package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
	"attempt4/platform/hash"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/smtp"
	"attempt4/platform/zap"
	"fmt"
	"time"
)

type UserService struct {
	userRepository   repository.UserRepository
	roleRepository   repository.RoleRepository
	walletRepository repository.WalletRepository
	appLogService    app_log.ApplicationLogService
}

func NewUserService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	walletRepository repository.WalletRepository,
	appLogService app_log.ApplicationLogService) UserService {
	u := UserService{
		userRepository,
		roleRepository,
		walletRepository,
		appLogService,
	}
	return u
}

func (u *UserService) DeleteUser(id int32) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return internal.UserNotFound
	}

	user.Status = enum.UserDeletedStatus

	deletedTime := time.Now()
	user.DeletedAt = &deletedTime

	err = u.userRepository.Delete(user)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	wallet, err := u.walletRepository.GetByUserId(user.Id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	wallet.DeletedAt = &deletedTime

	err = u.walletRepository.Delete(wallet)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	role, err := u.roleRepository.GetByUserId(user.Id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if role.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.RoleNotFound.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return internal.RoleNotFound
	}

	err = u.roleRepository.Delete(role)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (u *UserService) GetUserById(id int32) (dto.UserDto, error) {
	userDto := dto.UserDto{}
	user, err := u.userRepository.GetById(id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return userDto, err
	}
	if user.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
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
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
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
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	return nil
}

func (u *UserService) UpdateUserPassword(id int32, userDto dto.UserUpdatePasswordDto) error {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return internal.UserNotFound
	}

	err = hash.CompareEncryptedPasswords(user.Password, userDto.Password)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
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
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (u *UserService) CreateUser(userDto dto.UserDto) error {
	user, err := u.userRepository.GetByName(userDto.Username)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	if user.Id != 0 {
		zap.Logger.Error(internal.UserExist)
		return internal.UserExist
	}

	encryptedPassword, err := hash.EncryptPassword(userDto.Password)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}

	currentTime := time.Now()
	expiredTime := currentTime.Add(time.Second * 300)
	code := generateCode()

	user = entity.User{
		Username:      userDto.Username,
		Password:      encryptedPassword,
		Email:         userDto.Email,
		Name:          userDto.Name,
		Surname:       userDto.Surname,
		Status:        enum.UserPassiveStatus,
		Code:          code,
		CodeExpiredAt: &expiredTime,
		BirthDate:     &userDto.BirthDate,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
		DeletedAt:     nil,
	}

	user, err = u.userRepository.Create(user)
	if err != nil {
		zap.Logger.Error(err)
		return internal.UserNotCreated
	}

	key := entity.Role{
		UserId: user.Id,
		Rol:    enum.RoleUser,
	}

	key, err = u.roleRepository.Create(key)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	if key.Id == 0 {
		zap.Logger.Error(internal.RoleNotCreated)
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
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		zap.Logger.Error(internal.WalletNotCreated)
		return internal.WalletNotCreated
	}

	toEmail := []string{userDto.Email}
	err = smtp.SendMail(toEmail, *code)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) ActivateUser(codeDto dto.UserUpdateCodeDto) error {
	user, err := u.userRepository.GetByName(codeDto.Username)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		zap.Logger.Error(internal.UserNotFound)
		return internal.UserNotFound
	}

	if user.CodeExpiredAt.Before(time.Now()) {
		user.Code = generateCode()
		expiredCode := time.Now().Add(time.Second * 300)
		user.CodeExpiredAt = &expiredCode

		err = u.userRepository.Update(user)
		if err != nil {
			zap.Logger.Error(err)
			return err
		}
		zap.Logger.Error(internal.ExceedVerifyCode)
		return internal.ExceedVerifyCode
	}

	if codeDto.Code != *user.Code {
		zap.Logger.Error(internal.FailInVerify)
		return internal.FailInVerify
	}

	user.Status = enum.UserActiveStatus
	err = u.userRepository.Update(user)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (u *UserService) GetUserRoleById(id int32) (int, error) {
	user, err := u.userRepository.GetById(id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return 0, err
	}
	if user.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return 0, internal.UserNotFound
	}

	rol, err := u.roleRepository.GetByUserId(user.Id)
	if err != nil {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return 0, err
	}
	if rol.Id == 0 {
		u.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.RoleNotFound.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.RoleNotFound)
		return 0, internal.RoleNotFound
	}

	return rol.Rol, nil
}

func generateCode() *string {
	code := fmt.Sprint(time.Now().Nanosecond())[:6]
	return &code
}
