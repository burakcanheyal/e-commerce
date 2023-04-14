package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	zap2 "go.uber.org/zap"
	"time"
)

type OrderService struct {
	orderRepos   repository.OrderRepository
	productRepos repository.ProductRepository
	userRepos    repository.UserRepository
}

func NewOrderService(
	orderRepos repository.OrderRepository,
	productRepos repository.ProductRepository,
	userRepos repository.UserRepository) OrderService {

	o := OrderService{
		orderRepos,
		productRepos,
		userRepos,
	}
	return o
}

func (o *OrderService) CreateOrder(orderDto dto.OrderDto, id int32) (dto.OrderDescriptionDto, error) {
	orderDescription := dto.OrderDescriptionDto{}

	start := o.orderRepos.Db.Begin()

	order, err := o.orderRepos.GetById(orderDto.Id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDescription, err
	}
	if order.Id != 0 {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDescription, internal.OrderExist
	}

	user, err := o.userRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDescription, err
	}
	if user.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.UserNotFound))
		return orderDescription, internal.UserNotFound
	}

	product, err := o.productRepos.GetById(orderDto.ProductId)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDescription, err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
		return orderDescription, internal.ProductNotFound
	}

	if orderDto.Quantity > product.Quantity {
		return orderDescription, internal.ExceedOrder
	}

	order = entity.Order{
		UserId:    user.Id,
		ProductId: orderDto.ProductId,
		Quantity:  orderDto.Quantity,
		Status:    enum.OrderActive,
		Price:     float32(orderDto.Quantity) * product.Price,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
		DeletedAt: nil,
	}

	order, err = o.orderRepos.Create(order)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDescription, err
	}

	product.Quantity = product.Quantity - orderDto.Quantity
	if product.Quantity == 0 {
		product.Status = enum.ProductUnAvailable
	}

	err = o.productRepos.Update(product)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		start.Rollback()
		return orderDescription, err
	}

	start.Commit()

	productDto := dto.ProductDto{
		Name:     product.Name,
		Quantity: orderDto.Quantity,
		Price:    product.Price,
	}

	orderDescription = dto.OrderDescriptionDto{
		Username: user.Username,
		Products: productDto,
	}

	return orderDescription, nil
}

func (o *OrderService) DeleteOrder(id int32) error {
	order, err := o.orderRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if order.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.OrderNotFound))
		return internal.OrderNotFound
	}

	order.Status = enum.OrderCancel

	deletedTime := time.Now()
	order.DeletedAt = &deletedTime

	err = o.orderRepos.Delete(order)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	return nil
}

func (o *OrderService) GetOrderById(id int32) (dto.OrderDto, error) {
	orderDto := dto.OrderDto{}
	order, err := o.orderRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return orderDto, err
	}
	if order.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.OrderNotFound))
		return orderDto, internal.OrderNotFound
	}

	orderDto = dto.OrderDto{
		Id:        0,
		ProductId: order.ProductId,
		Quantity:  order.Quantity,
	}

	return orderDto, nil
}

func (o *OrderService) UpdateOrder(orderDto dto.OrderDto, id int32) error {
	order, err := o.orderRepos.GetById(orderDto.Id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if order.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.OrderNotFound))
		return internal.OrderNotFound
	}

	user, err := o.userRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if user.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.UserNotFound))
		return internal.UserNotFound
	}

	product, err := o.productRepos.GetById(orderDto.ProductId)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
		return internal.ProductNotFound
	}

	updatedTime := time.Now()

	order = entity.Order{
		Id:        order.Id,
		UserId:    user.Id,
		ProductId: orderDto.ProductId,
		Quantity:  orderDto.Quantity,
		Price:     float32(orderDto.Quantity) * product.Price,
		UpdatedAt: &updatedTime,
	}

	err = o.orderRepos.Update(order)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}

	return nil
}

func (o *OrderService) GetAllOrders(id int32, filter dto.Filter, pagination dto.Pagination) ([]dto.ProductDto, int64, error) {
	var productDto []dto.ProductDto
	var order []entity.Order
	var totalNumber int64

	user, err := o.userRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, totalNumber, err
	}
	if user.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.UserNotFound))
		return productDto, totalNumber, internal.UserNotFound
	}

	order, totalNumber, err = o.orderRepos.GetAllOrders(filter, pagination, user.Id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, totalNumber, err
	}

	for i, _ := range order {
		product, err := o.productRepos.GetById(order[i].ProductId)
		if err != nil {
			zap.Logger.Error("Hata", zap2.Error(err))
			return productDto, totalNumber, err
		}
		productDto = append(productDto, dto.ProductDto{
			Name:     product.Name,
			Quantity: product.Quantity,
			Price:    product.Price,
		})
	}

	return productDto, totalNumber, nil
}