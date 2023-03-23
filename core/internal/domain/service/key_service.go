package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
)

type KeyService struct {
	userRepository  repository.UserRepository
	keyRepository   repository.KeyRepository
	panelRepository repository.PanelRepository
}

func NewKeyService(
	userRepository repository.UserRepository,
	keyRepository repository.KeyRepository,
	panelRepository repository.PanelRepository) KeyService {
	k := KeyService{
		userRepository,
		keyRepository,
		panelRepository,
	}
	return k
}

func (k *KeyService) SendRequestToUpdateUserRole(id int32) error {
	key, err := k.keyRepository.GetByUserId(id)
	if err != nil {
		return err
	}

	if key.Status == enum.WaitingKeyStatus {
		return internal.KeyWaiting
	}

	if key.Status == enum.NonApprovedKeyStatus {
		return internal.KeyNonApproved
	}

	panel := entity.Panel{
		OperationId: enum.OperationRoleChange,
		Status:      enum.OperationWaiting,
	}

	panel, err = k.panelRepository.Create(panel)
	if err != nil {
		return err
	}

	key.Status = enum.WaitingKeyStatus

	err = k.keyRepository.Update(key)
	if err != nil {
		return err
	}

	return nil
}
func (k *KeyService) ResponseToUpdateUserRole(ResponseDto dto.PanelDto) error {
	key, err := k.keyRepository.GetByUserId(ResponseDto.UserId)
	if err != nil {
		return err
	}

	if key.Status != enum.WaitingKeyStatus {
		return internal.KeyResponded
	}

	panel, err := k.panelRepository.GetById(ResponseDto.Id)
	if err != nil {
		return err
	}

	panel.Status = ResponseDto.Response

	if ResponseDto.Response == enum.OperationApproved {
		key.Rol = enum.RoleManager
		key.Status = enum.ApprovedKeyStatus
	} else {
		key.Status = enum.NonApprovedKeyStatus
	}

	err = k.panelRepository.Update(panel)
	if err != nil {
		return err
	}

	err = k.keyRepository.Update(key)
	if err != nil {
		return err
	}

	return nil
}
