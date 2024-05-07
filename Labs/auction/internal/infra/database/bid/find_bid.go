package bid

import (
	"context"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/bid_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

func (bd *BidRepository) FindBidByActionId(ctx context.Context,
	actionId string,
) ([]bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"action_id": actionId}

	cursor, err := bd.Collection.Find(ctx, filter)
	var errMsg string
	if err != nil {
		errMsg = fmt.Sprintf("Error trying to find bids by Auction Id %s", actionId)
		logger.Error(errMsg, err)
		return nil, internal_error.NewInternalServerError(errMsg)
	}

	defer cursor.Close(ctx)

	var bidsEntityMongo []BidEntityMongo
	if err := cursor.All(ctx, bidsEntityMongo); err != nil {
		logger.Error(errMsg, err)
		return nil, internal_error.NewInternalServerError(errMsg)
	}

	var bids []bid_entity.Bid
	for _, bid := range bidsEntityMongo {
		bids = append(bids, bid_entity.Bid{
			Id:        bid.Id,
			UserId:    bid.UserId,
			AuctionId: bid.AuctionId,
			Amount:    bid.Amount,
			Timestamp: time.Unix(bid.Timestamp, 0),
		})
	}
	return bids, nil
}

func (bd *BidRepository) FindWinningBidByActionId(ctx context.Context,
	actionId string) (*bid_entity.Bid, *internal_error.InternalError) {

	filter := bson.M{"action_id": actionId}
	var bidEntityMongo BidEntityMongo

	opts := options.FindOne().SetSort(bson.D{{"amount", -1}})
	if err := bd.Collection.FindOne(ctx, filter, opts).Decode(bidEntityMongo); err != nil {
		errMsg := fmt.Sprintf("Error trying to find bids by Auction Id %s", actionId)
		logger.Error(errMsg, err)
		return nil, internal_error.NewInternalServerError(errMsg)
	}

	return &bid_entity.Bid{
		Id:        bidEntityMongo.Id,
		UserId:    bidEntityMongo.UserId,
		AuctionId: bidEntityMongo.AuctionId,
		Amount:    bidEntityMongo.Amount,
		Timestamp: time.Unix(bidEntityMongo.Timestamp, 0),
	}, nil
}
