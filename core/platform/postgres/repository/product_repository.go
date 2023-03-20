package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"attempt4/core/internal/domain/enum"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) ProductRepository {
	p := ProductRepository{db}
	return p
}

func (p *ProductRepository) Create(product entity.Product) (entity.Product, error) {
	if err := p.db.Create(&product).Error; err != nil {
		return product, internal.DBNotCreated
	}

	return product, nil
}
func (p *ProductRepository) Delete(product entity.Product) error {
	if err := p.db.Where("status != ?", enum.DeletedProduct).Where("id = ?", product.Id).Updates(product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return internal.DBNotFound
		}
		return err
	}

	return nil
}
func (p *ProductRepository) GetByName(name string) (entity.Product, error) {
	var product entity.Product
	if err := p.db.Model(&product).Where("name=?", name).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, nil
		}
		return product, err
	}

	return product, nil
}
func (p *ProductRepository) GetById(id int32) (entity.Product, error) {
	var product entity.Product
	if err := p.db.Model(&product).Where("id=?", id).Scan(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return product, nil
		}
		return product, err
	}

	return product, nil
}

func (p *ProductRepository) GetAllProducts(filter dto.Filter, pagination dto.Pagination) ([]entity.Product, int64, error) {
	var productList []entity.Product
	var total int64
	var order string
	listQuery := p.db.Find(&productList).Where("status != ?", enum.DeletedProduct)

	if filter.Name != "" {
		listQuery = listQuery.Where("name ilike ?", "%"+filter.Name+"%")
	}

	if filter.Quantity != 0 {
		listQuery = listQuery.Where("quantity > ?", filter.Quantity)
	}

	if filter.Price != 0 {
		listQuery = listQuery.Where("price > ?", filter.Price)
	}

	if pagination.OrderBy != "" {
		order = pagination.OrderBy
		if pagination.SortBy != "" {
			order = order + " " + pagination.SortBy + " "
		}
	}

	if err := listQuery.Count(&total).Scopes(Paginate(pagination)).Order(order).Find(&productList).Error; err != nil {
		return productList, 0, err
	}
	return productList, total, nil
}
func (p *ProductRepository) Update(product entity.Product) error {
	if err := p.db.Model(&product).Where("id=?", product.Id).Updates(product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return internal.DBNotFound
		}
		return internal.DBNotUpdated
	}

	return nil
}
func (p *ProductRepository) GetIdByName(name string) (int32, error) {
	var product entity.Product
	if err := p.db.Model(&product).Where("name=?", name).Scan(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return 0, internal.DBNotFound
		}
		return 0, err
	}

	return product.Id, nil
}
