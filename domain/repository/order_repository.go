package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type OrderRepository interface {
	CreateOrder(ctx context.Context, order *entity.Order) (entity.Order, error)
}

type OrderRepositoryImpl struct {
	DB *orm.Provider
}

func (o *OrderRepositoryImpl) CreateOrder(ctx context.Context, order *entity.Order) (entity.Order, error) {
	if err := o.DB.WithContext(ctx).Create(order).Error; err != nil {
		return entity.Order{}, err
	}

	return entity.Order{
		ID: order.ID,
	}, nil
}

func NewOrderRepository(conn *orm.Provider) OrderRepository {
	return &OrderRepositoryImpl{DB: conn}
}