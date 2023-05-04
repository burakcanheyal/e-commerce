package service

import (
	"attempt4/internal"
	"attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"attempt4/platform/app_log"
	"attempt4/platform/postgres/repository"
	"attempt4/platform/zap"
	"time"
)

type ProductService struct {
	productRepository repository.ProductRepository
	userRepository    repository.UserRepository
	appLogService     app_log.ApplicationLogService
}

func NewProductService(
	productRepository repository.ProductRepository,
	userRepos repository.UserRepository,
	appLogService app_log.ApplicationLogService) ProductService {
	p := ProductService{
		productRepository,
		userRepos,
		appLogService,
	}
	return p
}

func (p *ProductService) CreateProduct(productDto dto.ProductDto, id int32) (dto.ProductDto, error) {
	product, err := p.productRepository.GetByName(productDto.Name)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return productDto, err
	}
	if product.Id != 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.ProductExist.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductExist)
		return productDto, internal.ProductExist
	}

	user, err := p.userRepository.GetById(id)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return productDto, err
	}
	if user.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.UserNotFound.Error(), RelatedTable: "User", CreatedAt: time.Now()})
		zap.Logger.Error(internal.UserNotFound)
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

	product, err = p.productRepository.Create(product)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return productDto, err
	}
	if product.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: id, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductNotFound)
		return productDto, internal.ProductNotFound
	}

	return productDto, nil
}

func (p *ProductService) DeleteProduct(name string, userId int32) error {
	product, err := p.productRepository.GetByName(name)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if product.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductNotFound)
		return internal.ProductNotFound
	}
	deletedTime := time.Now()

	product.DeletedAt = &deletedTime
	product.Status = enum.ProductDeleted

	err = p.productRepository.Delete(product)

	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	return nil
}

func (p *ProductService) GetProductByName(name string, userId int32) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepository.GetByName(name)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return productDto, err
	}
	if product.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductNotFound)
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
		Name:     product.Name,
		Quantity: product.Quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		zap.Logger.Error(internal.ProductDeleted)
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductDeleted.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductUnavailable.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductUnavailable)
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) GetProductById(id int32, quantity int32, userId int32) (dto.ProductDto, error) {
	productDto := dto.ProductDto{}
	product, err := p.productRepository.GetById(id)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return productDto, err
	}
	if product.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductNotFound)
		return productDto, internal.ProductNotFound
	}

	productDto = dto.ProductDto{
		Name:     product.Name,
		Quantity: quantity,
		Price:    product.Price,
	}

	if product.Status == enum.ProductDeleted {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductDeleted.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductDeleted)
		return productDto, internal.ProductDeleted
	}
	if product.Status == enum.ProductUnAvailable {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductUnavailable.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductUnavailable)
		return productDto, internal.ProductUnavailable
	}

	return productDto, nil
}

func (p *ProductService) UpdateProduct(productDto dto.ProductUpdateDto, userId int32) error {
	product, err := p.productRepository.GetByName(productDto.Name)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}
	if product.Id == 0 {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: internal.ProductNotFound.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(internal.ProductNotFound)
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
	err = p.productRepository.Update(entityProduct)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
		return err
	}

	return nil
}

func (p *ProductService) GetAllProducts(filter dto.Filter, pagination dto.Pagination, userId int32) ([]dto.ProductDto, int64, error) {
	var productsDto []dto.ProductDto
	var products []entity.Product
	var totalNumber int64
	var err error

	products, totalNumber, err = p.productRepository.GetAllProducts(filter, pagination)
	if err != nil {
		p.appLogService.AddLog(app_log.ApplicationLogDto{UserId: userId, LogType: "Error", Content: err.Error(), RelatedTable: "Product", CreatedAt: time.Now()})
		zap.Logger.Error(err)
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
