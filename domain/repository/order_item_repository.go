package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type OrderItemRepository interface {
	GetByID(ctx context.Context, id uint) (entity.OrderItem, error)
	CreateItem(ctx context.Context, order *entity.OrderItem) (entity.OrderItem, error)
}

type OrderItemRepositoryImpl struct {
	DB *orm.Provider
}

func (p *OrderItemRepositoryImpl) GetByID(ctx context.Context, id uint) (entity.OrderItem, error) {
	var item entity.OrderItem
	if err := p.DB.WithContext(ctx).Where("id = ?", id).First(&item).Error; err != nil {
		return entity.OrderItem{}, err
	}
	return item, nil
}

func (o *OrderItemRepositoryImpl) CreateItem(ctx context.Context, order *entity.OrderItem) (entity.OrderItem, error) {
	if err := o.DB.WithContext(ctx).Create(order).Error; err != nil {
		return entity.OrderItem{}, err
	}

	return entity.OrderItem{
		ID: order.ID,
	}, nil
}

func NewOrderItemRepository(conn *orm.Provider) OrderItemRepository {
	return &OrderItemRepositoryImpl{DB: conn}
}