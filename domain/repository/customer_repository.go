package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type CustomerRepository interface {
	GetByID(ctx context.Context, id uint) (entity.Customer, error)
}

type CustomerRepositoryImpl struct {
	DB *orm.Provider
}

func (p *CustomerRepositoryImpl) GetByID(ctx context.Context, id uint) (entity.Customer, error) {
	var customer entity.Customer
	if err := p.DB.WithContext(ctx).Where("id = ?", id).First(&customer).Error; err != nil {
		return entity.Customer{}, err
	}
	return customer, nil
}

func NewCustomerRepository(conn *orm.Provider) CustomerRepository {
	return &CustomerRepositoryImpl{DB: conn}
}