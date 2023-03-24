package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
	"strconv"
	"time"
)

type WalletService struct {
	userRepository            repository.UserRepository
	walletRepository          repository.WalletRepository
	productRepository         repository.ProductRepository
	orderRepository           repository.OrderRepository
	walletOperationRepository repository.WalletOperationRepository
}

func NewWalletService(
	userRepository repository.UserRepository,
	walletRepository repository.WalletRepository,
	productRepository repository.ProductRepository,
	orderRepository repository.OrderRepository,
	walletOperationRepository repository.WalletOperationRepository) WalletService {

	w := WalletService{
		userRepository,
		walletRepository,
		productRepository,
		orderRepository,
		walletOperationRepository,
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

func (w *WalletService) Purchase(id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if wallet.Id == 0 {
		if err != nil {
			return err
		}
		return internal.WalletNotFound
	}

	products := ""
	price := float64(0)

	orders, count, err := w.orderRepository.GetAllOrders(dto.Filter{}, dto.Pagination{}, id)
	if count == 0 {
		return internal.EmptyCart
	}

	start := w.walletRepository.Begin()

	for i, _ := range orders {
		price += orders[i].Price
		products += "-" + strconv.Itoa(int(orders[i].ProductId))

		product, err := w.productRepository.GetById(orders[i].ProductId)
		if product.Id == 0 {
			w.walletRepository.Rollback(start)
			if err != nil {
				return err
			}
			return internal.ProductNotFound
		}

		sellerWallet, err := w.walletRepository.GetByUserId(product.UserId)
		if sellerWallet.Id == 0 {
			w.walletRepository.Rollback(start)
			if err != nil {
				return err
			}
			return internal.WalletNotFound
		}

		sellerWallet.Balance = sellerWallet.Balance + orders[i].Price

		err = w.walletRepository.Update(wallet)
		if err != nil {
			w.walletRepository.Rollback(start)
			return err
		}
	}

	if price > wallet.Balance {
		w.walletRepository.Rollback(start)
		return internal.WalletInadequate
	}

	wallet.Balance = wallet.Balance - price

	err = w.walletRepository.Update(wallet)
	if err != nil {
		w.walletRepository.Rollback(start)
		return err
	}

	w.walletRepository.Commit(start)

	walletOperation := entity.WalletOperation{
		OperationNumber: "",
		Products:        products,
		Price:           price,
		UserId:          id,
		OperationDate:   time.Now(),
	}

	walletOperation, err = w.walletOperationRepository.Create(walletOperation)
	if walletOperation.Id == 0 {
		if err != nil {
			return err
		}
		return internal.FailInPurchase
	}
	return nil
}
