package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"attempt4/core/platform/postgres/repository"
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
	if product.Id != 0 {
		if err != nil {
			return productDto, err
		}
		return productDto, internal.ProductExist
	}

	user, err := p.UserRepos.GetById(id)
	if user.Id == 0 {
		if err != nil {
			return productDto, err
		}
		return productDto, internal.UserNotFound
	}

	product = entity.Product{
		Id:        product.Id,
		Name:      productDto.Name,
		Quantity:  productDto.Quantity,
		Price:     productDto.Price,
		Status:    enum.ProductAvailable,
		CreatedAt: time.Now(),
		DeletedAt: time.Now(),
		UpdatedAt: time.Now(),
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
	if product.Id == 0 {
		if err != nil {
			return err
		}
		return internal.ProductNotFound
	}

	product.Status = enum.ProductDeleted

	err = p.productRepos.Delete(product)

	if err != nil {
		return err
	}
	return nil
}

func (p *ProductService) GetProductByName(name string) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepos.GetByName(name)
	if product.Id == 0 {
		if err != nil {
			return productDto, err
		}
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
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

func (p *ProductService) GetProductById(id int32, quantity int32) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepos.GetById(id)
	if product.Id == 0 {
		if err != nil {
			return productDto, err
		}
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
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

func (p *ProductService) UpdateProduct(productDto dto.ProductUpdateDto) error {
	product, err := p.productRepos.GetByName(productDto.Name)
	if product.Id == 0 {
		if err != nil {
			return err
		}
		return internal.ProductNotFound
	}

	entityProduct := entity.Product{
		Id:       product.Id,
		Name:     productDto.Name,
		Quantity: productDto.Quantity,
		Price:    productDto.Price,
	}
	err = p.productRepos.Update(entityProduct)
	if err != nil {
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
