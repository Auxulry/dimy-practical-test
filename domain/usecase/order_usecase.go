package usecase

import (
	"context"
	"fmt"
	"github.com/MochamadAkbar/dimy-practical-test/common"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/domain/repository"
	"github.com/MochamadAkbar/dimy-practical-test/exception"
	"net/http"
)

type OrderUsecase interface {
	CreateOrder(ctx context.Context, order *entity.OrderRequest) error
	CreateItems(ctx context.Context, order *entity.Order, items []*entity.Items) error
}

type OrderUsecaseImpl struct {
	OrderRepository           repository.OrderRepository
	ProductRepository         repository.ProductRepository
	PaymentRepository         repository.PaymentRepository
	CustomerRepository        repository.CustomerRepository
	CustomerAddressRepository repository.CustomerAddressRepository
	OrderItemRepository       repository.OrderItemRepository
}

func (o *OrderUsecaseImpl) CreateOrder(ctx context.Context, order *entity.OrderRequest) error {
	var total float64

	customer, err := o.CustomerRepository.GetByID(ctx, order.CustomerID)
	if err != nil {
		return common.NewErrHTTP(http.StatusNotFound, fmt.Sprintf("Customer with id: %v Not Found", order.CustomerID))
	}

	address, err := o.CustomerAddressRepository.GetByID(ctx, order.CustomerAddressID)
	if err != nil {
		return common.NewErrHTTP(http.StatusNotFound, fmt.Sprintf("Customer Address with id: %v Not Found", order.CustomerAddressID))
	}

	for _, item := range order.OrderItems {
		product, err := o.ProductRepository.GetByID(ctx, item.ProductID)
		if err != nil {
			return common.NewErrHTTP(http.StatusNotFound, fmt.Sprintf("Product with id: %v Not Found", 1))
		}

		total = total + product.Price
	}

	payload := &entity.Order{
		CustomerID:        customer.ID,
		CustomerAddressID: address.ID,
		Total:             total,
	}

	result, err := o.OrderRepository.CreateOrder(ctx, payload)
	if err != nil {
		return exception.NewException(err.Error())
	}

	err = o.CreateItems(ctx, &result, order.OrderItems)
	if err != nil {
		return err
	}
	return nil
}

func (o *OrderUsecaseImpl) CreateItems(ctx context.Context, order *entity.Order, items []*entity.Items) error {
	for _, item := range items {
		product, err := o.ProductRepository.GetByID(ctx, item.ProductID)
		if err != nil {
			return common.NewErrHTTP(http.StatusNotFound, fmt.Sprintf("Product with id: %v Not Found", 1))
		}

		payment, err := o.PaymentRepository.GetByID(ctx, item.PaymentMethodID)
		if err != nil {
			return common.NewErrHTTP(http.StatusNotFound, fmt.Sprintf("Payment Method with id: %v Not Found", 1))
		}

		payload := &entity.OrderItem{
			OrderID:         order.ID,
			ProductID:       product.ID,
			PaymentMethodID: payment.ID,
			Total:           product.Price,
		}

		_, err = o.OrderItemRepository.CreateItem(ctx, payload)
		if err != nil {
			return exception.NewException(err.Error())
		}
	}
	return nil
}

func NewOrderUsecase(
	order repository.OrderRepository,
	product repository.ProductRepository,
	payment repository.PaymentRepository,
	customer repository.CustomerRepository,
	address repository.CustomerAddressRepository,
	items repository.OrderItemRepository,
) OrderUsecase {
	return &OrderUsecaseImpl{
		OrderRepository:           order,
		ProductRepository:         product,
		PaymentRepository:         payment,
		CustomerRepository:        customer,
		CustomerAddressRepository: address,
		OrderItemRepository:       items,
	}
}