package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
	"fmt"
	"github.com/hoisie/mustache"
	"github.com/pdfcrowd/pdfcrowd-go"
	"os"
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
	if err != nil {
		return err
	}
	if wallet.Id == 0 {
		return internal.WalletNotFound
	}

	user, err := w.userRepository.GetById(id)
	if err != nil {
		return err
	}
	if user.Id == 0 {
		return internal.UserNotFound
	}

	balance := wallet.Balance + walletDto.Balance

	updatedTime := time.Now()

	wallet = entity.Wallet{
		Id:        wallet.Id,
		UserId:    user.Id,
		Balance:   balance,
		Status:    enum.WalletActive,
		UpdatedAt: &updatedTime,
	}

	err = w.walletRepository.Update(wallet)
	if err != nil {
		return err
	}

	return nil
}

func (w *WalletService) Purchase(id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if err != nil {
		return err
	}
	if wallet.Id == 0 {
		return internal.WalletNotFound
	}

	price := float32(0)

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

		balance := sellerWallet.Balance + orders[i].Price
		sellerWallet.Balance = balance

		currentTime := time.Now()

		walletOperation := entity.WalletOperation{
			OperationNumber: RandomString(8),
			Type:            enum.WalletSellType,
			Balance:         orders[i].Price,
			UserId:          &product.UserId,
			OrderId:         &orders[i].Id,
			ProductId:       &product.Id,
			OperationDate:   currentTime,
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

	balance := wallet.Balance - price
	wallet.Balance = balance

	err = w.walletRepository.Update(wallet)
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		return err
	}

	currentTime := time.Now()

	walletOperation := entity.WalletOperation{
		OperationNumber: RandomString(8),
		Type:            enum.WalletBuyType,
		Balance:         price,
		UserId:          &id,
		OrderId:         nil,
		ProductId:       nil,
		OperationDate:   currentTime,
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

func (w *WalletService) GetAllTransactions(id int32, transactionType int8) ([]dto.TransactionDto, int64, error) {
	transactions, total, err := w.walletOperationRepository.GetAllTransactionsWithJoinTable(id, transactionType)
	var list []dto.TransactionDto
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, internal.TransactionNotFound
	}

	for i, _ := range transactions {
		var l dto.TransactionDto
		if transactions[i].OrderId == nil {
			l.OrderId = 0
			l.OrderQuantity = 0
		} else {
			l.OrderId = *transactions[i].OrderId
			l.OrderQuantity = transactions[i].Order.Quantity
		}

		if transactions[i].ProductId == nil {
			l.ProductName = ""
		} else {
			l.ProductName = transactions[i].Product.Name
			l.SellerName = transactions[i].Product.User.Name
		}

		l.OperationNumber = transactions[i].OperationNumber
		l.Balance = transactions[i].Balance
		l.OperationDate = transactions[i].OperationDate

		list = append(list, l)
	}
	return list, total, nil
}

func (w *WalletService) ShowStatistics(id int32) error {
	items, _, err := w.GetAllTransactions(id, enum.WalletSellType)
	if err != nil {
		return err
	}

	type transactionList struct {
		Transactions []dto.TransactionDto
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	list := transactionList{items}

	filename := fmt.Sprintf("%s/core/internal/application/template/transactions.html", dir)
	result := mustache.RenderFile(filename, list)
	client := pdfcrowd.NewHtmlToPdfClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")

	pdf := fmt.Sprintf("%s/core/internal/application/template/pdf/statistic.pdf", dir)
	err = client.ConvertStringToFile(result, pdf)

	if err != nil {
		return err
	}

	return nil
}
