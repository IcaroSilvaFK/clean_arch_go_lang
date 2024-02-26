package usecase

import (
	"database/sql"

	"github.com/IcaroSilvaFK/fcc-cleanArch/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListsOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {

	return &ListOrdersUseCase{
		orderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]*OrderOutputDTO, error) {

	orders, err := l.OrderRepository.ListOrders()

	var output []*OrderOutputDTO

	if err == sql.ErrNoRows {
		return output, nil
	}

	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		output = append(output, &OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})

	}

	return output, nil
}
