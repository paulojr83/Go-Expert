package auction_usecase

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/bid_usecase"
)

func (au *AuctionUseCase) FindAuctionById(ctx context.Context,
	auctionId string,
) (*AuctionOutputDto, *internal_error.InternalError) {

	auction, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}
	return &AuctionOutputDto{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductionCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}, nil
}

func (au *AuctionUseCase) FindAuctions(ctx context.Context,
	status AuctionStatus,
	category, productName string) ([]AuctionOutputDto, *internal_error.InternalError) {
	auctions, err := au.auctionRepositoryInterface.FindAuctions(ctx, auction_entity.AuctionStatus(status), category, productName)
	if err != nil {
		return nil, err
	}

	var auctionsOutputDto []AuctionOutputDto
	for _, auction := range auctions {
		auctionsOutputDto = append(auctionsOutputDto, AuctionOutputDto{
			Id:          auction.Id,
			ProductName: auction.ProductName,
			Category:    auction.Category,
			Description: auction.Description,
			Condition:   ProductionCondition(auction.Condition),
			Status:      AuctionStatus(auction.Status),
			Timestamp:   auction.Timestamp,
		})
	}
	return auctionsOutputDto, nil
}

func (au *AuctionUseCase) FindWinningBidByAuctionId(ctx context.Context,
	auctionId string) (*WinningInfoOutputDto, *internal_error.InternalError) {

	auction, err := au.auctionRepositoryInterface.FindAuctionById(ctx, auctionId)
	if err != nil {
		return nil, err
	}

	auctionDto := AuctionOutputDto{
		Id:          auction.Id,
		ProductName: auction.ProductName,
		Category:    auction.Category,
		Description: auction.Description,
		Condition:   ProductionCondition(auction.Condition),
		Status:      AuctionStatus(auction.Status),
		Timestamp:   auction.Timestamp,
	}
	bidWinning, err := au.bidRepositoryInterface.FindWinningBidByActionId(ctx, auctionId)
	if err != nil {
		return &WinningInfoOutputDto{
			Auction: auctionDto,
			Bid:     nil,
		}, nil
	}

	return &WinningInfoOutputDto{
		Auction: auctionDto,
		Bid: &bid_usecase.BidOutputDto{
			Id:        bidWinning.Id,
			UserId:    bidWinning.UserId,
			AuctionId: bidWinning.AuctionId,
			Amount:    bidWinning.Amount,
			Timestamp: bidWinning.Timestamp,
		},
	}, nil
}
