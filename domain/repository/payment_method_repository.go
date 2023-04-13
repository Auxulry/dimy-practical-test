package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type PaymentRepository interface {
	GetByID(ctx context.Context, id uint) (entity.PaymentMethod, error)
}

type PaymentRepositoryImpl struct {
	DB *orm.Provider
}

func (p *PaymentRepositoryImpl) GetByID(ctx context.Context, id uint) (entity.PaymentMethod, error) {
	var payment entity.PaymentMethod
	if err := p.DB.WithContext(ctx).Where("id = ?", id).First(&payment).Error; err != nil {
		return entity.PaymentMethod{}, err
	}
	return payment, nil
}

func NewPaymentRepository(conn *orm.Provider) PaymentRepository {
	return &PaymentRepositoryImpl{DB: conn}
}