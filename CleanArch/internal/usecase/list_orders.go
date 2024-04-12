package usecase

import (
	"fmt"
	"github.com/paulojr83/Go-Expert/CleanArch/internal/entity"
)

type ListOrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type ListOrderUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrderUseCase {
	return &ListOrderUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrderUseCase) Execute() ([]OrderOutputDTO, error) {
	orders, err := c.OrderRepository.List()
	if err != nil {
		return []OrderOutputDTO{}, err
	}

	fmt.Println("Buscando orders:")
	fmt.Println(orders)
	var ordersOutPut []OrderOutputDTO
	for _, order := range orders {
		dto := OrderOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		}
		ordersOutPut = append(ordersOutPut, dto)
	}

	return ordersOutPut, nil
}
