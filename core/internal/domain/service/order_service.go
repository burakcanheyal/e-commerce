package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
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
	if order.OrderId != 0 {
		if err != nil {
			return orderDescription, err
		}
		return orderDescription, internal.OrderExist
	}

	user, err := o.userRepos.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return orderDescription, err
		}
		return orderDescription, internal.UserNotFound
	}

	product, err := o.productRepos.GetById(orderDto.ProductId)
	if product.Id == 0 {
		if err != nil {
			return orderDescription, err
		}
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
		Price:     float64(orderDto.Quantity) * float64(product.Price),
	}

	order, err = o.orderRepos.Create(order)
	if err != nil {
		return orderDescription, err
	}

	product.Quantity = product.Quantity - orderDto.Quantity
	if product.Quantity == 0 {
		product.Status = enum.ProductUnAvailable
	}

	err = o.productRepos.Update(product)
	if err != nil {
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
	if order.OrderId == 0 {
		if err != nil {
			return err
		}
		return internal.OrderNotFound
	}

	order.Status = enum.OrderCancel

	err = o.orderRepos.Delete(order)

	if err != nil {
		return err
	}
	return nil
}

func (o *OrderService) GetOrderById(id int32) (dto.OrderDto, error) {
	orderDto := dto.OrderDto{}
	order, err := o.orderRepos.GetById(id)
	if order.OrderId == 0 {
		if err != nil {
			return orderDto, err
		}
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
	if order.OrderId == 0 {
		if err != nil {
			return err
		}
		return internal.OrderNotFound
	}

	user, err := o.userRepos.GetById(id)
	if err != nil {
		return err
	}

	product, err := o.productRepos.GetById(orderDto.ProductId)
	if product.Id == 0 {
		if err != nil {
			return err
		}
		return internal.ProductNotFound
	}

	order = entity.Order{
		OrderId:   order.OrderId,
		UserId:    user.Id,
		ProductId: orderDto.ProductId,
		Quantity:  orderDto.Quantity,
		Price:     float64(orderDto.Quantity) * float64(product.Price),
	}

	err = o.orderRepos.Update(order)
	if err != nil {
		return err
	}

	return nil
}

func (o *OrderService) GetAllOrders(id int32, filter dto.Filter, pagination dto.Pagination) ([]dto.ProductDto, int64, error) {
	var productDto []dto.ProductDto
	var order []entity.Order
	var totalNumber int64
	var err error

	user, err := o.userRepos.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return productDto, totalNumber, err
		}
		return productDto, totalNumber, internal.UserNotFound
	}

	order, totalNumber, err = o.orderRepos.GetAllOrders(filter, pagination, user.Id)
	if err != nil {
		return productDto, totalNumber, err
	}

	for i, _ := range order {
		product, err := o.productRepos.GetById(order[i].ProductId)
		if err != nil {
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
