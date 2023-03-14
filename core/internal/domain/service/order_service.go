package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/platform/jwt"
	"attempt4/core/platform/postgres/repository"
)

type OrderService struct {
	orderRepos   repository.OrderRepository
	productRepos repository.ProductRepository
	userRepos    repository.UserRepository
	Secret       string
}

func NewOrderService(orderRepos repository.OrderRepository, productRepos repository.ProductRepository,
	userRepos repository.UserRepository, secret string) OrderService {
	o := OrderService{orderRepos, productRepos, userRepos, secret}
	return o
}

func (o *OrderService) CreateOrder(orderDto dto.OrderDto, tokenString string) (dto.OrderDescriptionDto, error) {
	orderDescription := dto.OrderDescriptionDto{}

	order, err := o.orderRepos.GetById(orderDto.Id)
	if order.OrderId != 0 {
		return orderDescription, err
	}

	username, err := jwt.ExtractUsernameFromToken(tokenString, o.Secret)
	if err != nil {
		return orderDescription, err
	}

	user, err := o.userRepos.GetByName(username)
	if err != nil {
		return orderDescription, err
	}

	product, err := o.productRepos.GetById(orderDto.ProductId)
	if product.Id == 0 {
		return orderDescription, err
	}

	order = entity.Order{
		UserId:    user.Id,
		ProductId: orderDto.ProductId,
		Quantity:  orderDto.Quantity,
	}

	order, err = o.orderRepos.Create(order)
	if err != nil {
		return orderDescription, err
	}

	if orderDto.Quantity > product.Quantity {
		return orderDescription, internal.ExceedOrder
	}

	product.Quantity = product.Quantity - orderDto.Quantity
	err = o.productRepos.Update(product)
	if err != nil {
		return orderDescription, err
	}

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
	order, _ := o.orderRepos.GetById(id)
	if order.OrderId == 0 {
		return internal.OrderNotFound
	}

	err := o.orderRepos.Delete(order.OrderId)

	if err != nil {
		return err
	}
	return nil
}
func (o *OrderService) GetOrderById(id int32) (dto.OrderDto, error) {
	order, _ := o.orderRepos.GetById(id)

	OrderDto := dto.OrderDto{
		Id:        0,
		ProductId: order.ProductId,
		Quantity:  order.Quantity,
	}

	if order.OrderId == 0 {
		return OrderDto, internal.OrderNotFound
	}

	return OrderDto, nil
}
func (o *OrderService) UpdateOrder(orderDto dto.OrderDto, tokenString string) error {
	order, _ := o.orderRepos.GetById(orderDto.Id)
	if order.OrderId == 0 {
		return internal.OrderNotFound
	}

	username, err := jwt.ExtractUsernameFromToken(tokenString, o.Secret)
	if err != nil {
		return err
	}

	user, err := o.userRepos.GetByName(username)
	if err != nil {
		return err
	}

	order = entity.Order{
		OrderId:   order.OrderId,
		UserId:    user.Id,
		ProductId: orderDto.ProductId,
		Quantity:  orderDto.Quantity,
	}

	err = o.orderRepos.Update(order)
	if err != nil {
		return err
	}

	return nil
}
func (o *OrderService) GetAllOrders(tokenString string, filter dto.Filter, pagination dto.Pagination) ([]dto.ProductDto, int64, error) {
	var productDto []dto.ProductDto
	var order []entity.Order
	var totalNumber int64
	var err error

	username, err := jwt.ExtractUsernameFromToken(tokenString, o.Secret)
	if err != nil {
		return productDto, totalNumber, err
	}

	user, err := o.userRepos.GetByName(username)
	if err != nil {
		return productDto, totalNumber, err
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
