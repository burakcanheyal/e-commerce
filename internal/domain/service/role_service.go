package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"math/rand"
	"time"
)

type RolService struct {
	userRepository       repository.UserRepository
	roleRepository       repository.RoleRepository
	submissionRepository repository.SubmissionRepository
	appLogService        app_log.ApplicationLogService
}

func NewRolService(
	userRepository repository.UserRepository,
	roleRepository repository.RoleRepository,
	submissionRepository repository.SubmissionRepository,
	appLogService app_log.ApplicationLogService) RolService {
	k := RolService{
		userRepository,
		roleRepository,
		submissionRepository,
		appLogService,
	}
	return k
}

func (k *RolService) SubmissionUserRole(id int32) error {
	operation, err := k.submissionRepository.GetByUserId(id)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if operation.Id != 0 {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.OperationWaiting.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.OperationWaiting)
		return internal.OperationWaiting
	}

	if operation.Status != enum.SubmissionWaiting {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.OperationResponded.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.OperationResponded)
		return internal.OperationResponded
	}

	receiverUserId := int32(1)

	operation = entity.Submission{
		SubmissionNumber: RandomString(15),
		SubmissionType:   enum.SubmissionRolChange,
		Status:           enum.SubmissionWaiting,
		AppliedUserId:    id,
		ReceiverUserId:   &receiverUserId,
		OperationDate:    time.Now(),
	}

	operation, err = k.submissionRepository.Create(operation)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if operation.Id == 0 {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.OperationNotCreated.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.OperationNotCreated)
		return internal.OperationNotCreated
	}

	return nil
}

func (k *RolService) ResultOfUpdateUserRole(ResponseDto dto.AppOperationDto, id int32) error {
	operation, err := k.submissionRepository.GetByUserId(ResponseDto.UserId)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if operation.Id == 0 {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.OperationNotFound.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.OperationNotFound)
		return internal.OperationNotFound
	}

	if operation.SubmissionNumber != ResponseDto.OperationNumber {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.OperationFailInNumber.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.OperationFailInNumber)
		return internal.OperationFailInNumber
	}

	operation.Status = ResponseDto.Response

	key, err := k.roleRepository.GetByUserId(ResponseDto.UserId)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if key.Id == 0 {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.RoleNotFound.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(internal.RoleNotFound)
		return internal.RoleNotFound
	}

	if ResponseDto.Response == enum.SubmissionApproved {
		key.Rol = enum.RoleManager
		operation.Status = enum.SubmissionApproved
	} else {
		operation.Status = enum.SubmissionRejected
	}

	currentTime := time.Now()

	operation.OperationResultDate = &currentTime

	err = k.submissionRepository.Update(operation)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	err = k.roleRepository.Update(key)
	if err != nil {
		k.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Role", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func RandomString(len int) string {

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	str := string(bytes)
	return str
}

func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}
