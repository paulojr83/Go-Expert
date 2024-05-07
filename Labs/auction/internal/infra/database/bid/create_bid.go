package bid

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/bid_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/database/auction"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

type BidEntityMongo struct {
	Id        string  `bson:"_id"`
	UserId    string  `bson:"user_id"`
	AuctionId string  `bson:"auction_id"`
	Amount    float64 `bson:"amount"`
	Timestamp int64   `bson:"timestamp"`
}

type BidRepository struct {
	Collection        *mongo.Collection
	AuctionRepository *auction.AuctionRepository
}

func NewBidRepository(database *mongo.Database, auctionRepository *auction.AuctionRepository) *BidRepository {
	return &BidRepository{
		Collection:        database.Collection("bid"),
		AuctionRepository: auctionRepository}
}
func (bd *BidRepository) CreateBid(ctx context.Context,
	bidEntities []bid_entity.Bid,
) *internal_error.InternalError {
	var wg sync.WaitGroup

	for _, bid := range bidEntities {
		wg.Add(1)

		go func(bidValue bid_entity.Bid) {
			defer wg.Done()

			actionEntity, err := bd.AuctionRepository.FindAuctionById(ctx, bidValue.AuctionId)
			if err != nil {
				logger.Error("Error trying to find Auction", err)
				return
			}

			if actionEntity.Status != auction_entity.Active {
				return
			}

			bidEntityMongo := &BidEntityMongo{
				Id:        bidValue.Id,
				UserId:    bidValue.UserId,
				AuctionId: bidValue.AuctionId,
				Amount:    bidValue.Amount,
				Timestamp: bidValue.Timestamp.Unix(),
			}

			if _, err := bd.Collection.InsertOne(ctx, bidEntityMongo); err != nil {
				logger.Error("Error trying to insert Bid", err)
				return
			}
		}(bid)
	}

	wg.Wait()
	return nil
}
