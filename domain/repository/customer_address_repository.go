package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type CustomerAddressRepository interface {
	GetByID(ctx context.Context, id uint) (entity.CustomerAddress, error)
}

type CustomerAddressRepositoryImpl struct {
	DB *orm.Provider
}

func (p *CustomerAddressRepositoryImpl) GetByID(ctx context.Context, id uint) (entity.CustomerAddress, error) {
	var address entity.CustomerAddress
	if err := p.DB.WithContext(ctx).Where("id = ?", id).First(&address).Error; err != nil {
		return entity.CustomerAddress{}, err
	}
	return address, nil
}

func NewCustomerAddressRepository(conn *orm.Provider) CustomerAddressRepository {
	return &CustomerAddressRepositoryImpl{DB: conn}
}