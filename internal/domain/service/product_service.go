package service

import (
	"attempt4/internal"
	dto2 "attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	repository2 "attempt4/platform/postgres/repository"
	"time"
)

type ProductService struct {
	productRepos repository2.ProductRepository
	UserRepos    repository2.UserRepository
}

func NewProductService(
	productRepository repository2.ProductRepository,
	userRepos repository2.UserRepository) ProductService {
	p := ProductService{
		productRepository,
		userRepos,
	}
	return p
}

func (p *ProductService) CreateProduct(productDto dto2.ProductDto, id int32) (dto2.ProductDto, error) {
	product, err := p.productRepos.GetByName(productDto.Name)
	if err != nil {
		return productDto, err
	}
	if product.Id != 0 {
		return productDto, internal.ProductExist
	}

	user, err := p.UserRepos.GetById(id)
	if err != nil {
		return productDto, err
	}
	if user.Id == 0 {
		return productDto, internal.UserNotFound
	}

	product = entity.Product{
		Id:        product.Id,
		Name:      productDto.Name,
		Quantity:  productDto.Quantity,
		Price:     productDto.Price,
		Status:    enum.ProductAvailable,
		CreatedAt: time.Now(),
		DeletedAt: nil,
		UpdatedAt: nil,
		UserId:    user.Id,
	}

	product, err = p.productRepos.Create(product)
	if err != nil {
		return productDto, err
	}

	return productDto, nil
}

func (p *ProductService) DeleteProduct(name string) error {
	product, err := p.productRepos.GetByName(name)
	if err != nil {
		return err
	}
	if product.Id == 0 {
		return internal.ProductNotFound
	}
	deletedTime := time.Now()

	product.DeletedAt = &deletedTime
	product.Status = enum.ProductDeleted

	err = p.productRepos.Delete(product)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetProductByName(name string) (dto2.ProductDto, error) {
	productDto := dto2.ProductDto{}
	product, err := p.productRepos.GetByName(name)
	if err != nil {
		return productDto, err
	}
	if product.Id == 0 {
		return productDto, internal.ProductNotFound
	}

	productDto = dto2.ProductDto{
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) GetProductById(id int32, quantity int32) (dto2.ProductDto, error) {
	productDto := dto2.ProductDto{}
	product, err := p.productRepos.GetById(id)
	if err != nil {
		return productDto, err
	}
	if product.Id == 0 {
		return productDto, internal.ProductNotFound
	}

	productDto = dto2.ProductDto{
		Name:     product.Name,
		Quantity: quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) UpdateProduct(productDto dto2.ProductUpdateDto) error {
	product, err := p.productRepos.GetByName(productDto.Name)
	if err != nil {
		return err
	}
	if product.Id == 0 {
		return internal.ProductNotFound
	}

	updatedTime := time.Now()

	entityProduct := entity.Product{
		Id:        product.Id,
		Name:      productDto.Name,
		Quantity:  productDto.Quantity,
		Price:     productDto.Price,
		UpdatedAt: &updatedTime,
	}
	err = p.productRepos.Update(entityProduct)
	if err != nil {
		return err
	}

	return nil
}

func (p *ProductService) GetAllProducts(filter dto2.Filter, pagination dto2.Pagination) ([]dto2.ProductDto, int64, error) {
	var productsDto []dto2.ProductDto
	var products []entity.Product
	var totalNumber int64
	var err error

	products, totalNumber, err = p.productRepos.GetAllProducts(filter, pagination)
	if err != nil {
		return productsDto, totalNumber, err
	}

	for i, _ := range products {
		productsDto = append(productsDto,
			dto2.ProductDto{
				Name:     products[i].Name,
				Quantity: products[i].Quantity,
				Price:    products[i].Price,
			})
	}

	return productsDto, totalNumber, nil
}
