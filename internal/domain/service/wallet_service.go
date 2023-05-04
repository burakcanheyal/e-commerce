package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
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
	appLogService             app_log.ApplicationLogService
}

func NewWalletService(
	userRepository repository.UserRepository,
	walletRepository repository.WalletRepository,
	productRepository repository.ProductRepository,
	orderRepository repository.OrderRepository,
	walletOperationRepository repository.WalletOperationRepository,
	roleRepository repository.RoleRepository,
	appLogService app_log.ApplicationLogService) WalletService {

	w := WalletService{
		userRepository,
		walletRepository,
		productRepository,
		orderRepository,
		walletOperationRepository,
		roleRepository,
		appLogService,
	}
	return w
}

func (w *WalletService) UpdateBalance(walletDto dto.WalletDto, id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if err != nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	user, err := w.userRepository.GetById(id)
	if err != nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if user.Id == 0 {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (w *WalletService) Purchase(id int32) error {
	wallet, err := w.walletRepository.GetByUserId(id)
	if err != nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if wallet.Id == 0 {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(internal.WalletNotFound)
		return internal.WalletNotFound
	}

	price := float32(0)

	orders, count, err := w.orderRepository.GetAllOrders(dto.Filter{}, dto.Pagination{}, id)
	if count == 0 {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.EmptyCart.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
			zap.Logger.Error(err)
			return err
		}
		if product.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
			zap.Logger.Error(internal.ProductNotFound)
			return internal.ProductNotFound
		}

		sellerWallet, err := w.walletRepository.GetByUserId(product.UserId)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
			zap.Logger.Error(err)
			return err
		}
		if sellerWallet.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
			zap.Logger.Error(err)
			return err
		}
		if walletOperation.Id == 0 {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.FailInPurchase.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
			zap.Logger.Error(internal.FailInPurchase)
			return internal.FailInPurchase
		}

		orders[i].Status = enum.OrderCompleted
		err = w.orderRepository.Update(orders[i])
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Order", CreatedAt: time.Now()})
			zap.Logger.Error(err)
			return err
		}

		err = w.walletRepository.Update(wallet)
		if err != nil {
			w.orderRepository.Rollback(startOrderRepository)
			w.walletRepository.Rollback(startWalletRepository)
			w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
			zap.Logger.Error(err)
			return err
		}
	}

	if price > wallet.Balance {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.WalletInadequate.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(internal.WalletInadequate)
		return internal.WalletInadequate
	}

	balance := wallet.Balance - price
	wallet.Balance = balance

	err = w.walletRepository.Update(wallet)
	if err != nil {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if walletOperation.Id == 0 {
		w.orderRepository.Rollback(startOrderRepository)
		w.walletRepository.Rollback(startWalletRepository)
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.FailInPurchase.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return list, total, err
	}
	if total == 0 {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.TransactionNotFound.Error(), RelatedTable: "Wallet", CreatedAt: time.Now()})
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
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Transactions", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return nil, err
	}

	transactionList := dto.Transaction{Transactions: items}

	ChartData := dto.PieChartData{}

	var total float32
	for i, _ := range items {
		total += items[i].Balance
	}

	for i, _ := range items {
		ChartData.PieChartData = append(
			ChartData.PieChartData,
			dto.PieItemDto{
				OperationNumber: items[i].OperationNumber,
				Ratio:           fmt.Sprintf("%0.2f", items[i].Balance/total),
			})
	}

	dir, err := os.Getwd()
	if err != nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Transactions", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return nil, err
	}

	filename := fmt.Sprintf("%s/internal/application/template/transactions.html", dir)
	result := mustache.RenderFile(filename, transactionList, ChartData)

	preparedPdfData := wkhtmltopdf.RequestPdf{Body: result}

	pdfPath := fmt.Sprintf("%s/internal/application/template/pdf/statistic.pdf", dir)
	pdf, err := preparedPdfData.GeneratePDF(pdfPath)
	if err != nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Transactions", CreatedAt: time.Now()})
		return nil, err
	}
	if pdf == nil {
		w.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.PdfNotCreated.Error(), RelatedTable: "Transactions", CreatedAt: time.Now()})
		return nil, internal.PdfNotCreated
	}

	return pdf, nil
}
