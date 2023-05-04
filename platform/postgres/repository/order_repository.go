package repository

import (
	"attempt4/internal"
	dto "attempt4/internal/domain/dto"
	"attempt4/internal/domain/entity"
	"attempt4/internal/domain/enum"
	"gorm.io/gorm"
)

type OrderRepository struct {
	Db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	o := OrderRepository{db}
	return o
}

func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {
	if err := o.Db.Create(&order).Error; err != nil {
		return order, internal.DBNotCreated
	}
	return order, nil
}

func (o *OrderRepository) Delete(order entity.Order) error {
	if err := o.Db.Model(&order).Where("id = ?", order.Id).Update("status", enum.OrderCancel).Error; err != nil {
		return internal.DBNotDeleted
	}
	if err := o.Db.Model(&order).Where("id = ?", order.Id).Update("deleted_at", order.DeletedAt).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}

func (o *OrderRepository) GetById(id int32) (entity.Order, error) {
	var order entity.Order
	if err := o.Db.Model(&order).Where("id=?", id).Scan(&order).Error; err != nil {
		return order, internal.DBNotFound
	}
	return order, nil
}

func (o *OrderRepository) Update(order entity.Order) error {
	if err := o.Db.Model(&order).Where("status != ?", enum.OrderCancel).Where("id=?", order.Id).Updates(
		entity.Order{
			ProductId: order.ProductId,
			Quantity:  order.Quantity,
			Status:    order.Status,
			Price:     order.Price,
			UpdatedAt: order.UpdatedAt,
		}).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}

func (o *OrderRepository) GetAllOrders(filter dto.Filter, pagination dto.Pagination, userId int32) ([]entity.Order, int64, error) {
	var orderList []entity.Order
	var total int64

	listQuery := o.Db.Find(&orderList).Where("status = ?", enum.OrderActive).Where("user_id = ?", userId)

	if filter.Quantity != 0 {
		listQuery = listQuery.Where("quantity > ?", filter.Quantity)
	}

	order := "quantity" + " " + pagination.SortBy

	if pagination.Page == 0 {
		if err := listQuery.Count(&total).Order(order).Find(&orderList).Error; err != nil {
			return orderList, 0, err
		}
	} else {
		if err := listQuery.Count(&total).Scopes(Paginate(pagination)).Order(order).Find(&orderList).Error; err != nil {
			return orderList, 0, err
		}
	}

	return orderList, total, nil
}

func (o *OrderRepository) Begin() *gorm.DB {
	return o.Db.Begin()
}

func (o *OrderRepository) Rollback(rollback *gorm.DB) {
	rollback.Rollback()
}

func (o *OrderRepository) Commit(commit *gorm.DB) {
	commit.Commit()
}
