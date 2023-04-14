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

type ProductService struct {
	productRepos repository.ProductRepository
	UserRepos    repository.UserRepository
}

func NewProductService(
	productRepository repository.ProductRepository,
	userRepos repository.UserRepository) ProductService {
	p := ProductService{
		productRepository,
		userRepos,
	}
	return p
}

func (p *ProductService) CreateProduct(productDto dto.ProductDto, id int32) (dto.ProductDto, error) {
	product, err := p.productRepos.GetByName(productDto.Name)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, err
	}
	if product.Id != 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductExist))
		return productDto, internal.ProductExist
	}

	user, err := p.UserRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, err
	}
	if user.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.UserNotFound))
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
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, err
	}

	return productDto, nil
}

func (p *ProductService) DeleteProduct(name string) error {
	product, err := p.productRepos.GetByName(name)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
		return internal.ProductNotFound
	}
	deletedTime := time.Now()

	product.DeletedAt = &deletedTime
	product.Status = enum.ProductDeleted

	err = p.productRepos.Delete(product)

	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	return nil
}

func (p *ProductService) GetProductByName(name string) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepos.GetByName(name)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductDeleted))
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductUnavailable))
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) GetProductById(id int32, quantity int32) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepos.GetById(id)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productDto, err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
		Name:     product.Name,
		Quantity: quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductDeleted))
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductUnavailable))
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) UpdateProduct(productDto dto.ProductUpdateDto) error {
	product, err := p.productRepos.GetByName(productDto.Name)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}
	if product.Id == 0 {
		zap.Logger.Error("Hata", zap2.Error(internal.ProductNotFound))
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
		zap.Logger.Error("Hata", zap2.Error(err))
		return err
	}

	return nil
}

func (p *ProductService) GetAllProducts(filter dto.Filter, pagination dto.Pagination) ([]dto.ProductDto, int64, error) {
	var productsDto []dto.ProductDto
	var products []entity.Product
	var totalNumber int64
	var err error

	products, totalNumber, err = p.productRepos.GetAllProducts(filter, pagination)
	if err != nil {
		zap.Logger.Error("Hata", zap2.Error(err))
		return productsDto, totalNumber, err
	}

	for i, _ := range products {
		productsDto = append(productsDto,
			dto.ProductDto{
				Name:     products[i].Name,
				Quantity: products[i].Quantity,
				Price:    products[i].Price,
			})
	}

	return productsDto, totalNumber, nil
}
