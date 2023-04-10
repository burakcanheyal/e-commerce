package service

import (
	"attempt4/internal"
	dto2 "attempt4/internal/domain/dto"
	entity2 "attempt4/internal/domain/entity"
	enum2 "attempt4/internal/domain/enum"
	repository2 "attempt4/platform/postgres/repository"
	"fmt"
	"github.com/hoisie/mustache"
	"github.com/pdfcrowd/pdfcrowd-go"
	"os"
	"time"
)

type WalletService struct {
	userRepository            repository2.UserRepository
	walletRepository          repository2.WalletRepository
	productRepository         repository2.ProductRepository
	orderRepository           repository2.OrderRepository
	walletOperationRepository repository2.WalletOperationRepository
	roleRepository            repository2.RoleRepository
}

func NewWalletService(
	userRepository repository2.UserRepository,
	walletRepository repository2.WalletRepository,
	productRepository repository2.ProductRepository,
	orderRepository repository2.OrderRepository,
	walletOperationRepository repository2.WalletOperationRepository,
	roleRepository repository2.RoleRepository) WalletService {

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
func (w *WalletService) UpdateBalance(walletDto dto2.WalletDto, id int32) error {
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

	wallet = entity2.Wallet{
		Id:        wallet.Id,
		UserId:    user.Id,
		Balance:   balance,
		Status:    enum2.WalletActive,
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

	orders, count, err := w.orderRepository.GetAllOrders(dto2.Filter{}, dto2.Pagination{}, id)
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

		walletOperation := entity2.WalletOperation{
			OperationNumber: RandomString(8),
			Type:            enum2.WalletSellType,
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

		orders[i].Status = enum2.OrderCompleted
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

	walletOperation := entity2.WalletOperation{
		OperationNumber: RandomString(8),
		Type:            enum2.WalletBuyType,
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

func (w *WalletService) GetAllTransactions(id int32, transactionType int8) ([]dto2.TransactionDto, int64, error) {
	transactions, total, err := w.walletOperationRepository.GetAllTransactionsWithJoinTable(id, transactionType)
	var list []dto2.TransactionDto
	if err != nil {
		return list, total, err
	}
	if total == 0 {
		return list, total, internal.TransactionNotFound
	}

	for i, _ := range transactions {
		var l dto2.TransactionDto
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
	items, _, err := w.GetAllTransactions(id, enum2.WalletSellType)
	if err != nil {
		return err
	}

	type transactionList struct {
		Transactions []dto2.TransactionDto
	}

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	list := transactionList{items}

	filename := fmt.Sprintf("%s/internal/application/template/transactions.html", dir)
	result := mustache.RenderFile(filename, list)

	pdfPath := fmt.Sprintf("%s/internal/application/template/pdf/statistic.pdf", dir)

	client := pdfcrowd.NewHtmlToPdfClient("demo", "ce544b6ea52a5621fb9d55f8b542d14d")
	err = client.ConvertStringToFile(result, pdfPath)
	if err != nil {
		return err
	}
	/*
		preparedPdfData := wkhtmltopdf.RequestPdf{Body: result}

		bool, err := preparedPdfData.GeneratePDF(pdfPath)
		if !bool {
			return err
		}
	*/

	return nil
}
