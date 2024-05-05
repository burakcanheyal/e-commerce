package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"errors"
	"time"
)

type AdminService struct {
	userRepository   repository.UserRepository
	roleRepository   repository.RoleRepository
	walletRepository repository.WalletRepository
	appLogService    app_log.ApplicationLogService
	tripRepository   repository.TripRepository
}

func NewAdminService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	walletRepository repository.WalletRepository,
	appLogService app_log.ApplicationLogService,
	tripRepository repository.TripRepository) AdminService {
	u := AdminService{
		userRepository,
		roleRepository,
		walletRepository,
		appLogService,
		tripRepository,
	}
	return u
}
func (a *AdminService) DeleteUser(id int32) error {
	user, err := a.userRepository.GetById(id)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return internal.UserNotFound
	}

	user.Status = enum.UserDeletedStatus

	deletedTime := time.Now()
	user.DeletedAt = &deletedTime

	err = a.userRepository.Delete(user)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	wallet, err := a.walletRepository.GetByUserId(user.Id)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	wallet.DeletedAt = &deletedTime

	err = a.walletRepository.Delete(wallet)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	role, err := a.roleRepository.GetByUserId(user.Id)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if role.Id == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.RoleNotFound.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return internal.RoleNotFound
	}

	err = a.roleRepository.Delete(role)
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}
func (a *AdminService) GetAllUsers() ([]dto.UserDto, error) {
	var userDto []dto.UserDto
	user, total, err := a.userRepository.GetAllByUser()
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: 0, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return userDto, err
	}
	if total == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: 0, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return userDto, internal.UserNotFound
	}

	for _, userT := range user {
		userDtoTemp := dto.UserDto{
			Id:        userT.Id,
			Username:  userT.Username,
			Password:  userT.Password,
			Email:     userT.Email,
			Name:      userT.Name,
			Surname:   userT.Surname,
			Phone:     userT.Phone,
			Status:    userT.Status,
			BirthDate: *userT.BirthDate,
		}
		userDto = append(userDto, userDtoTemp)
	}

	return userDto, nil
}
func (a *AdminService) GetAllTrips() ([]dto.TripDto, error) {
	var tripDto []dto.TripDto
	trip, total, err := a.tripRepository.GetAllTrips()
	if err != nil {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: 0, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return tripDto, err
	}
	if total == 0 {
		a.appLogService.AddLog(app_log.ApplicationLogDto{UserId: 0, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
		return tripDto, internal.UserNotFound
	}

	for _, tripT := range trip {
		tripDtoTemp := dto.TripDto{
			Id:          tripT.Id,
			Description: tripT.Description,
			Name:        tripT.Name,
			Status:      tripT.Status,
			Lat:         tripT.Lat,
			Lng:         tripT.Lng,
		}
		tripDto = append(tripDto, tripDtoTemp)
	}

	return tripDto, nil
}
func (a *AdminService) DeleteTrip(id int32) error {
	trip, err := a.tripRepository.GetById(id)
	if trip.Id == 0 {
		return errors.New("trip not found")
	}
	if err != nil {
		return err
	}

	err = a.tripRepository.Delete(trip)
	if err != nil {
		return err
	}
	return nil
}
func (a *AdminService) AddTrip(tripDto dto.TripAddDto) error {
	trip := entity.Trip{
		Id:              0,
		Description:     tripDto.Description,
		Name:            tripDto.Name,
		Status:          enum.TripActive,
		Lat:             tripDto.Lat,
		Lng:             tripDto.Lng,
		HistoricalPoint: tripDto.HistoricalPoint,
		NaturePoint:     tripDto.NaturePoint,
		AdventurePoint:  tripDto.AdventurePoint,
		CulturalPoint:   tripDto.CulturalPoint,
		RelaxPoint:      tripDto.RelaxPoint,
	}
	temp, err := a.tripRepository.Create(trip)
	if temp.Id == 0 {
		return errors.New("trip not created")
	}
	if err != nil {
		return err
	}
	return nil
}
