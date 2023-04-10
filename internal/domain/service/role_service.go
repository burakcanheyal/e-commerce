package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	enum2 "attempt4/internal/domain/enum"
	repository2 "attempt4/platform/postgres/repository"
	"math/rand"
	"time"
)

type RolService struct {
	userRepository       repository2.UserRepository
	keyRepository        repository2.RoleRepository
	submissionRepository repository2.SubmissionRepository
}

func NewRolService(
	userRepository repository2.UserRepository,
	keyRepository repository2.RoleRepository,
	submissionRepository repository2.SubmissionRepository) RolService {
	k := RolService{
		userRepository,
		keyRepository,
		submissionRepository,
	}
	return k
}

func (k *RolService) SubmissionUserRole(id int32) error {
	operation, err := k.submissionRepository.GetByUserId(id)
	if err != nil {
		return err
	}
	if operation.Id != 0 {
		return internal.OperationWaiting
	}

	if operation.Status != enum2.SubmissionWaiting {
		return internal.OperationResponded
	}

	receiverUserId := int32(1)

	operation = entity.Submission{
		SubmissionNumber: RandomString(15),
		SubmissionType:   enum2.SubmissionRolChange,
		Status:           enum2.SubmissionWaiting,
		AppliedUserId:    id,
		ReceiverUserId:   &receiverUserId,
		OperationDate:    time.Now(),
	}

	operation, err = k.submissionRepository.Create(operation)
	if err != nil {
		return err
	}
	if operation.Id == 0 {
		return internal.OperationNotCreated
	}

	return nil
}

func (k *RolService) ResultOfUpdateUserRole(ResponseDto dto.AppOperationDto, id int32) error {
	operation, err := k.submissionRepository.GetByUserId(ResponseDto.UserId)
	if err != nil {
		return err
	}
	if operation.Id == 0 {
		return internal.OperationNotFound
	}

	if operation.SubmissionNumber != ResponseDto.OperationNumber {
		return internal.OperationFailInNumber
	}

	operation.Status = ResponseDto.Response

	key, err := k.keyRepository.GetByUserId(ResponseDto.UserId)
	if err != nil {
		return err
	}
	if key.Id == 0 {
		return internal.RoleNotFound
	}

	if ResponseDto.Response == enum2.SubmissionApproved {
		key.Rol = enum2.RoleManager
		operation.Status = enum2.SubmissionApproved
	} else {
		operation.Status = enum2.SubmissionRejected
	}

	currentTime := time.Now()

	operation.OperationResultDate = &currentTime

	err = k.submissionRepository.Update(operation)
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

	str := string(bytes)
	return str
}
func randInt(min int, max int) int {

	return min + rand.Intn(max-min)
}
