package auction_entity

import (
	"context"
	"github.com/google/uuid"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"time"
)

func NewAuction(productName, category, description string,
	condition ProductionCondition) (*Auction, *internal_error.InternalError) {
	auction := &Auction{
		Id:          uuid.New().String(),
		ProductName: productName,
		Category:    category,
		Description: description,
		Condition:   condition,
		Status:      Active,
		Timestamp:   time.Now(),
	}
	if err := auction.Validate(); err != nil {
		return nil, err
	}
	return auction, nil
}

func (a *Auction) Validate() *internal_error.InternalError {
	if len(a.ProductName) <= 1 ||
		len(a.Category) <= 2 ||
		len(a.Description) <= 10 {
		return internal_error.NewBadRequestError("Invalid object")
	}
	return nil
}

type Auction struct {
	Id          string
	ProductName string
	Category    string
	Description string
	Condition   ProductionCondition
	Status      AuctionStatus
	Timestamp   time.Time
}

type ProductionCondition int
type AuctionStatus int

const (
	New ProductionCondition = iota
	Used
	Refurbished
)

const (
	Active AuctionStatus = iota
	Completed
)

type AuctionRepositoryInterface interface {
	CreateAuction(ctx context.Context, auction *Auction) *internal_error.InternalError
	FindAuctionById(ctx context.Context, auctionId string) (*Auction, *internal_error.InternalError)
	FindAuctions(ctx context.Context, status AuctionStatus, category, productName string) ([]Auction, *internal_error.InternalError)
}
