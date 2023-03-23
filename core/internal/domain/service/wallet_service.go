package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
)

type WalletService struct {
	userRepository   repository.UserRepository
	walletRepository repository.WalletRepository
}

func NewWalletService(
	userRepository repository.UserRepository,
	walletRepository repository.WalletRepository) WalletService {

	w := WalletService{
		userRepository,
		walletRepository,
	}
	return w
}
func (w *WalletService) UpdateBalance(walletDto dto.WalletDto, id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if wallet.Id == 0 {
		if err != nil {
			return err
		}
		return internal.WalletNotFound
	}

	user, err := w.userRepository.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return err
		}
		return internal.UserNotFound
	}

	wallet = entity.Wallet{
		Id:      wallet.Id,
		UserId:  user.Id,
		Balance: wallet.Balance + walletDto.Balance,
		Status:  enum.WalletActive,
	}

	err = w.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	return nil
}
