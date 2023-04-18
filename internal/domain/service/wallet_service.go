package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/wkhtmltopdf"
	"attempt4/platform/zap"
	"fmt"
	"github.com/hoisie/mustache"
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
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	user, err := w.userRepository.GetById(id)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		zap.Logger.Error(internal.UserNotFound)
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
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (w *WalletService) Purchase(id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if err != nil {
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	price := float32(0)

	orders, count, err := w.orderRepository.GetAllOrders(dto.Filter{}, dto.Pagination{}, id)
	if count == 0 {
		zap.Logger.Error(internal.EmptyCart)
		return internal.EmptyCart
	}

	startWalletRepository := w.walletRepository.Begin()
	startOrderRepository := w.orderRepository.Begin()

	for i, _ := range orders {
		price += orders[i].Price

		product, err := w.productRepository.GetById(orders[i].ProductId)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(err)
			return err
		}
		if product.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(internal.ProductNotFound)
			return internal.ProductNotFound
		}

		sellerWallet, err := w.walletRepository.GetByUserId(product.UserId)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(err)
			return err
		}
		if sellerWallet.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(internal.WalletNotFound)
			return internal.WalletNotFound
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
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(err)
			return err
		}
		if walletOperation.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(internal.FailInPurchase)
			return internal.FailInPurchase
		}

		orders[i].Status = enum.OrderCompleted
		err = w.orderRepository.Update(orders[i])
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(err)
			return err
		}

		err = w.walletRepository.Update(wallet)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			zap.Logger.Error(err)
			return err
		}
	}

	if price > wallet.Balance {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		zap.Logger.Error(internal.WalletInadequate)
		return internal.WalletInadequate
	}

	balance := wallet.Balance - price
	wallet.Balance = balance

	err = w.walletRepository.Update(wallet)
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		zap.Logger.Error(err)
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
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		zap.Logger.Error(err)
		return err
	}
	if walletOperation.Id == 0 {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		zap.Logger.Error(internal.FailInPurchase)
		return internal.FailInPurchase
	}

	w.orderRepository.Commit(startOrderRepository)
	w.walletRepository.Commit(startWalletRepository)
	return nil
}

func (w *WalletService) GetAllTransactions(id int32, transactionType int8) ([]dto.TransactionDto, int64, error) {
	transactions, total, err := w.walletOperationRepository.GetAllTransactionsWithJoinTable(id, transactionType)
	var list []dto.TransactionDto
	if err != nil {
		zap.Logger.Error(err)
		return list, total, err
	}
	if total == 0 {
		zap.Logger.Error(internal.TransactionNotFound)
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

func (w *WalletService) ShowStatistics(id int32) ([]byte, error) {
	items, _, err := w.GetAllTransactions(id, enum.WalletSellType)
	if err != nil {
		zap.Logger.Error(err)
		return nil, err
	}

	type transactionList struct {
		Transactions []dto.TransactionDto
	}

	dir, err := os.Getwd()
	if err != nil {
		zap.Logger.Error(err)
		return nil, err
	}

	list := transactionList{items}

	filename := fmt.Sprintf("%s/internal/application/template/transactions.html", dir)
	result := mustache.RenderFile(filename, list)

	pdfPath := fmt.Sprintf("%s/internal/application/template/pdf/statistic.pdf", dir)
	/*
			client := pdfcrowd.NewHtmlToPdfClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")
			err = client.ConvertStringToFile(result, pdfPath)

			if err != nil {
		zap.Logger.Error(err)
				return nil, err
			}
	*/

	preparedPdfData := wkhtmltopdf.RequestPdf{Body: result}

	pdf, err := preparedPdfData.GeneratePDF(pdfPath)
	if err != nil {
		return nil, err
	}
	if pdf == nil {
		return nil, internal.PdfNotCreated
	}

	return []byte(result), nil
}
