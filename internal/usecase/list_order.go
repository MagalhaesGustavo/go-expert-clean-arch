package usecase

import (
	"errors"

	"github.com/devfullcycle/20-CleanArch/internal/entity"
)

type OrderDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrderUseCase) Execute() ([]OrderDTO, error) {
	result, err := l.OrderRepository.GetAll()
	if err != nil {
		return []OrderDTO{}, err
	}

	if len(result) == 0 {
		return []OrderDTO{}, errors.New("no orders found")
	}

	OrderListDTO := make([]OrderDTO, len(result))

	for _, order := range result {
		OrderListDTO = append(OrderListDTO, OrderDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return OrderListDTO, nil
}
