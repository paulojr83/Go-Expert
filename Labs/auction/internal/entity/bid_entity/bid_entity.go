package bid_entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"time"
)

type Bid struct {
	Id        string
	UserId    string
	AuctionId string
	Amount    float64
	Timestamp time.Time
}

func NewBid(userId, auctionId string,
	amount float64) (*Bid, *internal_error.InternalError) {
	bid := &Bid{
		Id:        uuid.New().String(),
		UserId:    userId,
		AuctionId: auctionId,
		Amount:    amount,
		Timestamp: time.Now(),
	}
	if err := bid.Validate(); err != nil {
		return nil, err
	}
	return bid, nil
}

func (b *Bid) Validate() *internal_error.InternalError {
	if err := uuid.Validate(b.UserId); err != nil {
		return internal_error.NewBadRequestError("UserId is not valid id")
	} else if err := uuid.Validate(b.AuctionId); err != nil {
		return internal_error.NewBadRequestError("AuctionId is not valid id")
	} else if b.Amount <= 0 {
		return internal_error.NewBadRequestError("Amount is not a valid value")
	}
	return nil
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntities []Bid) *internal_error.InternalError
	FindBidByActionId(ctx context.Context, actionId string) ([]Bid, *internal_error.InternalError)
	FindWinningBidByActionId(ctx context.Context, actionId string) (*Bid, *internal_error.InternalError)
}
