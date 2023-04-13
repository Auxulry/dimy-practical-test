package repository

import (
	"context"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/pkg/orm"
)

type ProductRepository interface {
	GetByID(ctx context.Context, id uint) (entity.Product, error)
}

type ProductRepositoryImpl struct {
	DB *orm.Provider
}

func (p *ProductRepositoryImpl) GetByID(ctx context.Context, id uint) (entity.Product, error) {
	var product entity.Product
	if err := p.DB.WithContext(ctx).Where("id = ?", id).First(&product).Error; err != nil {
		return entity.Product{}, err
	}
	return product, nil
}

func NewProductRepository(conn *orm.Provider) ProductRepository {
	return &ProductRepositoryImpl{DB: conn}
}