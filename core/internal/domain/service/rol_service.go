package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
	"math/rand"
	"time"
)

type RolService struct {
	userRepository  repository.UserRepository
	keyRepository   repository.KeyRepository
	panelRepository repository.AppOperationRepository
}

func NewRolService(
	userRepository repository.UserRepository,
	keyRepository repository.KeyRepository,
	panelRepository repository.AppOperationRepository) RolService {
	k := RolService{
		userRepository,
		keyRepository,
		panelRepository,
	}
	return k
}

func (k *RolService) AppOperationToUpdateUserRole(id int32) error {
	operation, err := k.panelRepository.GetByUserId(id)
	if operation.Id != 0 {
		return internal.OperationWaiting
	}

	if operation.Status != enum.OperationWaiting {
		return internal.OperationResponded
	}

	operation = entity.AppOperation{
		OperationNumber: randomString(15),
		OperationId:     enum.OperationRoleChange,
		Status:          enum.OperationWaiting,
		AppliedUserId:   id,
		ReceiverUserId:  1,
		OperationDate:   time.Now(),
	}

	operation, err = k.panelRepository.Create(operation)
	if operation.Id == 0 {
		if err != nil {
			return err
		}
		return internal.OperationNotCreated
	}

	return nil
}

func (k *RolService) ResultOfUpdateUserRole(ResponseDto dto.AppOperationDto, id int32) error {
	operation, err := k.panelRepository.GetByUserId(ResponseDto.UserId)
	if operation.Id == 0 {
		if err != nil {
			return err
		}
		return internal.OperationNotFound
	}

	if operation.OperationNumber != ResponseDto.OperationNumber {
		return internal.OperationFailInNumber
	}

	operation.Status = ResponseDto.Response

	key, err := k.keyRepository.GetByUserId(ResponseDto.UserId)
	if key.KeyId == 0 {
		if err != nil {
			return err
		}
		return internal.KeyNotFound
	}

	if ResponseDto.Response == enum.OperationApproved {
		key.Rol = enum.RoleManager
		operation.Status = enum.OperationApproved
	} else {
		operation.Status = enum.OperationReject
	}

	operation.OperationResultDate = time.Now()

	err = k.panelRepository.Update(operation)
	if err != nil {
		return err
	}

	err = k.keyRepository.Update(key)
	if err != nil {
		return err
	}

	return nil
}
func randomString(len int) string {

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	return string(bytes)
}
func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}
