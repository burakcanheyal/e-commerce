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
	keyRepository   repository.RoleRepository
	panelRepository repository.SubmissionRepository
}

func NewRolService(
	userRepository repository.UserRepository,
	keyRepository repository.RoleRepository,
	panelRepository repository.SubmissionRepository) RolService {
	k := RolService{
		userRepository,
		keyRepository,
		panelRepository,
	}
	return k
}

func (k *RolService) SubmissionUserRole(id int32) error {
	operation, err := k.panelRepository.GetByUserId(id)
	if operation.Id != 0 {
		return internal.OperationWaiting
	}

	if operation.Status != enum.SubmissionWaiting {
		return internal.OperationResponded
	}

	operation = entity.Submission{
		SubmissionNumber: RandomString(15),
		SubmissionType:   enum.SubmissionRolChange,
		Status:           enum.SubmissionWaiting,
		AppliedUserId:    id,
		ReceiverUserId:   1,
		OperationDate:    time.Now(),
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

	if operation.SubmissionNumber != ResponseDto.OperationNumber {
		return internal.OperationFailInNumber
	}

	operation.Status = ResponseDto.Response

	key, err := k.keyRepository.GetByUserId(ResponseDto.UserId)
	if key.Id == 0 {
		if err != nil {
			return err
		}
		return internal.RoleNotFound
	}

	if ResponseDto.Response == enum.SubmissionApproved {
		key.Rol = enum.RoleManager
		operation.Status = enum.SubmissionApproved
	} else {
		operation.Status = enum.SubmissionRejected
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
func RandomString(len int) string {

	bytes := make([]byte, len)

	for i := 0; i < len; i++ {
		bytes[i] = byte(randInt(97, 122))
	}

	return string(bytes)
}
func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}
