package service

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/platform/postgres/repository"
)

type ProductService struct {
	productRepos repository.ProductRepository
}

func NewProductService(productRepository repository.ProductRepository) ProductService {
	p := ProductService{productRepository}
	return p
}

func (p *ProductService) CreateProduct(productDto dto.ProductDto) (dto.ProductDto, error) {
	product, err := p.productRepos.GetByName(productDto.Name)
	if product.Id != 0 {
		return productDto, internal.ProductExist
	}

	product = entity.Product{
		Id:       product.Id,
		Name:     productDto.Name,
		Quantity: productDto.Quantity,
		Price:    productDto.Price,
	}

	product, err = p.productRepos.Create(product)
	if err != nil {
		return productDto, err
	}

	return productDto, nil
}
func (p *ProductService) DeleteProduct(name string) error {
	product, _ := p.productRepos.GetByName(name)
	if product.Id == 0 {
		return internal.ProductNotFound
	}

	err := p.productRepos.Delete(product.Id)

	if err != nil {
		return err
	}
	return nil
}
func (p *ProductService) GetProductByName(name string) (dto.ProductDto, error) {
	product, _ := p.productRepos.GetByName(name)

	productDto := dto.ProductDto{
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	if product.Id == 0 {
		return productDto, internal.ProductNotFound
	}
	return productDto, nil
}
func (p *ProductService) GetProductById(id int32, quantity int32) (dto.ProductDto, error) {
	product, _ := p.productRepos.GetById(id)

	productDto := dto.ProductDto{
		Name:     product.Name,
		Quantity: quantity,
		Price:    product.Price,
	}

	if product.Id == 0 {
		return productDto, internal.ProductNotFound
	}
	return productDto, nil
}
func (p *ProductService) UpdateProduct(productDto dto.ProductUpdateDto) error {
	product, _ := p.productRepos.GetByName(productDto.Name)

	if product.Id == 0 {
		return internal.ProductNotFound
	}

	entityProduct := entity.Product{
		Id:       product.Id,
		Name:     productDto.Name,
		Quantity: productDto.Quantity,
		Price:    productDto.Price,
	}
	err := p.productRepos.Update(entityProduct)
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
