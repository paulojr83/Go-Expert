package auction

import (
	"context"
	"errors"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

func (ar AuctionRepository) FindAuctionById(ctx context.Context,
	auctionId string,
) (*auction_entity.Auction, *internal_error.InternalError) {

	var auctionEntityMongo AuctionEntityMongo

	filter := bson.M{"_id": auctionId}
	err := ar.Collection.FindOne(ctx, filter).Decode(auctionEntityMongo)

	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			logger.Error(fmt.Sprintf("Auction not found with this id =%s", auctionId), err)
			return nil, internal_error.NewNotFoundError(
				fmt.Sprintf("Auction not found with this id =%s", auctionId))
		}

		logger.Error("Error trying to find Auction by Id", err)
		return nil, internal_error.NewInternalServerError("Error trying to find Auction by Id")
	}

	return &auction_entity.Auction{
		Id:          auctionEntityMongo.Id,
		ProductName: auctionEntityMongo.ProductName,
		Category:    auctionEntityMongo.Category,
		Description: auctionEntityMongo.Description,
		Condition:   auctionEntityMongo.Condition,
		Status:      auctionEntityMongo.Status,
		Timestamp:   time.Unix(auctionEntityMongo.Timestamp, 0),
	}, nil
}

func (ar AuctionRepository) FindAuctions(ctx context.Context,
	status auction_entity.AuctionStatus, category, productName string,
) ([]auction_entity.Auction, *internal_error.InternalError) {

	filter := bson.M{}
	if status != 0 {
		filter["status"] = status
	}

	if category != "" {
		filter["category"] = primitive.Regex{
			Pattern: category,
			Options: "i",
		}
	}
	if productName != "" {
		filter["product_name"] = primitive.Regex{
			Pattern: productName,
			Options: "i",
		}
	}
	cursor, err := ar.Collection.Find(ctx, filter)

	if err != nil {
		logger.Error("Error trying to find Auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find Auctions")

	}
	defer cursor.Close(ctx)

	var auctionsEntityMongo []AuctionEntityMongo
	if err := cursor.All(ctx, auctionsEntityMongo); err != nil {
		logger.Error("Error trying to find Auctions", err)
		return nil, internal_error.NewInternalServerError("Error trying to find Auctions")
	}

	var auctions []auction_entity.Auction

	for _, entityMongo := range auctionsEntityMongo {
		auctions = append(auctions, auction_entity.Auction{
			Id:          entityMongo.Id,
			ProductName: entityMongo.ProductName,
			Category:    entityMongo.Category,
			Description: entityMongo.Description,
			Condition:   entityMongo.Condition,
			Status:      entityMongo.Status,
			Timestamp:   time.Unix(entityMongo.Timestamp, 0),
		})
	}
	return auctions, nil
}
