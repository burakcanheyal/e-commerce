package repository

import (
	"attempt4/core/internal"
	"attempt4/core/internal/domain/dto"
	"attempt4/core/internal/domain/entity"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	o := OrderRepository{db}
	return o
}
func (o *OrderRepository) Create(order entity.Order) (entity.Order, error) {
	if err := o.db.Create(&order).Error; err != nil {
		return order, internal.DBNotCreated
	}
	return order, nil
}

func (o *OrderRepository) Delete(id int32) error {
	var order entity.Order
	if err := o.db.Where("order_id = ?", id).Delete(&order).Error; err != nil {
		return internal.DBNotDeleted
	}
	return nil
}
func (o *OrderRepository) GetById(id int32) (entity.Order, error) {
	var order entity.Order
	if err := o.db.Model(&order).Where("order_id=?", id).Scan(&order).Error; err != nil {
		return order, internal.DBNotFound
	}
	return order, nil
}
func (o *OrderRepository) Update(order entity.Order) error {
	if err := o.db.Model(&order).Where("order_id=?", order.OrderId).Updates(order).Error; err != nil {
		return internal.DBNotUpdated
	}
	return nil
}

func (o *OrderRepository) GetAllOrders(filter dto.Filter, pagination dto.Pagination, userId int32) ([]entity.Order, int64, error) {
	var orderList []entity.Order
	var total int64

	listQuery := o.db.Find(&orderList).Count(&total)

	if filter.Quantity != 0 {
		listQuery = listQuery.Where("quantity > ?", filter.Quantity)
	}

	order := "quantity" + " " + pagination.SortBy

	if err := listQuery.Count(&total).Scopes(Paginate(pagination)).Order(order).Find(&orderList).Error; err != nil {
		return orderList, 0, err
	}
	return orderList, total, nil
}