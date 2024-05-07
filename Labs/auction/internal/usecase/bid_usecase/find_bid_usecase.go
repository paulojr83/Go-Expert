package bid_usecase

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
)

func (bd BidUseCase) FindBidByActionId(ctx context.Context,
	actionId string) ([]BidOutputDto, *internal_error.InternalError) {

	bids, err := bd.BidRepository.FindBidByActionId(ctx, actionId)
	if err != nil {
		return nil, err
	}

	var bidsOutput []BidOutputDto
	for _, bid := range bids {
		bidsOutput = append(bidsOutput, BidOutputDto{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: bid.Timestamp,
		})
	}
	return bidsOutput, nil
}
func (bd BidUseCase) FindWinningBidByActionId(ctx context.Context,
	actionId string) (*BidOutputDto, *internal_error.InternalError) {
	bid, err := bd.BidRepository.FindWinningBidByActionId(ctx, actionId)
	if err != nil {
		return nil, err
	}

	return &BidOutputDto{
		Id:        bid.Id,
		UserId:    bid.UserId,
		AuctionId: bid.AuctionId,
		Amount:    bid.Amount,
		Timestamp: bid.Timestamp,
	}, nil
}
