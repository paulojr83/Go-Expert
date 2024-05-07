package auction_usecase

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/bid_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/bid_usecase"
	"time"
)

type AuctionInputDto struct {
	ProductName string              `json:"product_name" binding:"required,min=1"`
	Category    string              `json:"category" binding:"required,min=2"`
	Description string              `json:"description" binding:"required,min=10,max=200"`
	Condition   ProductionCondition `json:"condition"`
}

type AuctionOutputDto struct {
	Id          string              `json:"id"`
	ProductName string              `json:"product_name"`
	Category    string              `json:"category"`
	Description string              `json:"description"`
	Condition   ProductionCondition `json:"condition"`
	Status      AuctionStatus       `json:"status"`
	Timestamp   time.Time           `json:"timestamp" time_format:"2006-01-02 15:10:00"`
}

type WinningInfoOutputDto struct {
	Auction AuctionOutputDto          `json:"auction"`
	Bid     *bid_usecase.BidOutputDto `json:"bid,omitempty"`
}

type ProductionCondition int64
type AuctionStatus int64

type AuctionUseCase struct {
	auctionRepositoryInterface auction_entity.AuctionRepositoryInterface
	bidRepositoryInterface     bid_entity.BidRepositoryInterface
}

func NewAuctionUseCase(auctionRepository auction_entity.AuctionRepositoryInterface,
	bidRepository bid_entity.BidRepositoryInterface) AuctionUseCaseInterface {
	return &AuctionUseCase{
		auctionRepositoryInterface: auctionRepository,
		bidRepositoryInterface:     bidRepository,
	}
}

type AuctionUseCaseInterface interface {
	CreateAuction(ctx context.Context, auctionInput AuctionInputDto) *internal_error.InternalError
	FindAuctionById(ctx context.Context, auctionId string) (*AuctionOutputDto, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]AuctionOutputDto, *internal_error.InternalError)
	FindWinningBidByAuctionId(ctx context.Context,
		auctionId string) (*WinningInfoOutputDto, *internal_error.InternalError)
}

func (au *AuctionUseCase) CreateAuction(ctx context.Context,
	auctionInput AuctionInputDto,
) *internal_error.InternalError {

	auction, err := auction_entity.NewAuction(
		auctionInput.ProductName, auctionInput.Category, auctionInput.Description,
		auction_entity.ProductionCondition(auctionInput.Condition))
	if err != nil {
		return err
	}

	if err := au.auctionRepositoryInterface.CreateAuction(ctx, auction); err != nil {
		return err
	}
	return nil
}
