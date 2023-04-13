package handler

import (
	"github.com/MochamadAkbar/dimy-practical-test/common"
	"github.com/MochamadAkbar/dimy-practical-test/domain/entity"
	"github.com/MochamadAkbar/dimy-practical-test/domain/usecase"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type OrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

type OrderHandlerImpl struct {
	OrderUsecase usecase.OrderUsecase
}

func (o *OrderHandlerImpl) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req entity.OrderRequest
	common.SerializeRequest(r, &req)

	err := o.OrderUsecase.CreateOrder(r.Context(), &req)
	if err != nil {
		result := common.ErrResponse{
			Code:    common.GetHTTPCode(err),
			Status:  http.StatusText(common.GetHTTPCode(err)),
			Message: err.Error(),
		}

		common.SerializeWriter(w, result.Code, result)
		return
	}

	result := common.Response{
		Code:   http.StatusCreated,
		Status: http.StatusText(http.StatusCreated),
		Data:   "Transaction Successfully Created",
	}

	common.SerializeWriter(w, result.Code, result)
}

func NewOrderHandler(usecase usecase.OrderUsecase, router chi.Router) error {
	handler := &OrderHandlerImpl{
		OrderUsecase: usecase,
	}

	router.Post("/order/create", handler.CreateOrder)

	return nil
}