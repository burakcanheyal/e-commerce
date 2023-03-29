package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
	"time"
)

type WalletService struct {
	userRepository            repository.UserRepository
	walletRepository          repository.WalletRepository
	productRepository         repository.ProductRepository
	orderRepository           repository.OrderRepository
	walletOperationRepository repository.WalletOperationRepository
	roleRepository            repository.RoleRepository
}

func NewWalletService(
	userRepository repository.UserRepository,
	walletRepository repository.WalletRepository,
	productRepository repository.ProductRepository,
	orderRepository repository.OrderRepository,
	walletOperationRepository repository.WalletOperationRepository,
	roleRepository repository.RoleRepository) WalletService {

	w := WalletService{
		userRepository,
		walletRepository,
		productRepository,
		orderRepository,
		walletOperationRepository,
		roleRepository,
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

	price := float64(0)

	orders, count, err := w.orderRepository.GetAllOrders(dto.Filter{}, dto.Pagination{}, id)
	if count == 0 {
		return internal.EmptyCart
	}

	startWalletRepository := w.walletRepository.Begin()
	startOrderRepository := w.orderRepository.Begin()

	for i, _ := range orders {
		price += orders[i].Price

		product, err := w.productRepository.GetById(orders[i].ProductId)
		if product.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return internal.ProductNotFound
		}
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return err
		}

		sellerWallet, err := w.walletRepository.GetByUserId(product.UserId)
		if sellerWallet.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return internal.WalletNotFound
		}
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return err
		}

		sellerWallet.Balance = sellerWallet.Balance + orders[i].Price

		walletOperation := entity.WalletOperation{
			OperationNumber: RandomString(8),
			Type:            enum.WalletSellType,
			Balance:         orders[i].Price,
			UserId:          product.UserId,
			OrderId:         orders[i].Id,
			ProductId:       product.Id,
			OperationDate:   time.Now(),
		}

		walletOperation, err = w.walletOperationRepository.Create(walletOperation)
		if walletOperation.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return internal.FailInPurchase
		}

		orders[i].Status = enum.OrderCompleted
		err = w.orderRepository.Update(orders[i])
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return err
		}

		err = w.walletRepository.Update(wallet)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			return err
		}
	}

	if price > wallet.Balance {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		return internal.WalletInadequate
	}

	wallet.Balance = wallet.Balance - price

	err = w.walletRepository.Update(wallet)
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		return err
	}

	//Todo: Satıcı ve alıcı için ayrı entityler
	walletOperation := entity.WalletOperation{
		OperationNumber: RandomString(8),
		Type:            enum.WalletBuyType,
		Balance:         price,
		UserId:          id,
		OrderId:         1,
		ProductId:       1,
		OperationDate:   time.Now(),
	}

	walletOperation, err = w.walletOperationRepository.Create(walletOperation)
	if walletOperation.Id == 0 {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		return internal.FailInPurchase
	}
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		return err
	}

	w.orderRepository.Commit(startOrderRepository)
	w.walletRepository.Commit(startWalletRepository)
	return nil
}

func (w *WalletService) GetAllTransactions(id int32) (error, []dto.TransactionDto) {
	transactions, total, err := w.walletOperationRepository.GetAllTransactions(id)
	var list []dto.TransactionDto
	if total == 0 {
		if err != nil {
			return err, list
		}
		return internal.WalletNotFound, list
	}

	for i, _ := range transactions {
		//Todo:Reposunda ayarla
		if transactions[i].Type == enum.WalletSellType {
			continue
		}
		//Todo: Preload
		product, err := w.productRepository.GetById(transactions[i].ProductId)
		if product.Id == 0 {
			return internal.ProductNotFound, list
		}
		if err != nil {
			return err, list
		}

		order, err := w.orderRepository.GetById(transactions[i].OrderId)
		if order.Id == 0 {
			return internal.OrderNotFound, list
		}
		if err != nil {
			return err, list
		}

		seller, err := w.userRepository.GetById(product.UserId)
		if seller.Id == 0 {
			return internal.UserNotFound, list
		}
		if err != nil {
			return err, list
		}

		//Todo: obje dön id ve name
		l := dto.TransactionDto{
			OperationNumber: transactions[i].OperationNumber,
			Balance:         transactions[i].Balance,
			OrderId:         transactions[i].OrderId,
			ProductName:     product.Name,
			OrderQuantity:   order.Quantity,
			SellerName:      seller.Name,
			OperationDate:   transactions[i].OperationDate,
		}
		list = append(list, l)
	}
	return nil, list
}

func (w *WalletService) ShowStatistics(id int32) ([]dto.SellStaticsDto, error) {
	transactions, total, err := w.walletOperationRepository.GetAllTransactions(id)
	var list []dto.SellStaticsDto
	if total == 0 {
		//Todo: Custom error
		return list, internal.WalletNotFound
	}
	if err != nil {
		return list, err
	}

	for i, _ := range transactions {
		if transactions[i].Type == enum.WalletSellType {
			continue
		}
		product, err := w.productRepository.GetById(transactions[i].ProductId)
		if product.Id == 0 {
			return list, internal.ProductNotFound
		}
		if err != nil {
			return list, err
		}

		order, err := w.orderRepository.GetById(transactions[i].OrderId)
		if order.Id == 0 {
			return list, internal.OrderNotFound
		}
		if err != nil {
			return list, err
		}

		buyer, err := w.userRepository.GetById(order.UserId)
		if buyer.Id == 0 {
			return list, internal.UserNotFound
		}
		if err != nil {
			return list, err
		}

		l := dto.SellStaticsDto{
			OperationNumber: transactions[i].OperationNumber,
			Balance:         transactions[i].Balance,
			OrderId:         transactions[i].OrderId,
			ProductName:     product.Name,
			OrderQuantity:   order.Quantity,
			BuyerName:       buyer.Name,
			OperationDate:   transactions[i].OperationDate,
		}
		list = append(list, l)
	}
	return list, nil
}
