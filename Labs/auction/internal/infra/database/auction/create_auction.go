package auction

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuctionEntityMongo struct {
	Id          string                             `bson:"_id"`
	ProductName string                             `bson:"product_name"`
	Category    string                             `bson:"category"`
	Description string                             `bson:"description"`
	Condition   auction_entity.ProductionCondition `bson:"condition"`
	Status      auction_entity.AuctionStatus       `bson:"status"`
	Timestamp   int64                              `bson:"timestamp"`
}

type AuctionRepository struct {
	Collection *mongo.Collection
}

func NewAuctionRepository(database *mongo.Database) *AuctionRepository {
	return &AuctionRepository{Collection: database.Collection("auctions")}
}

func (ar AuctionRepository) CreateAuction(ctx context.Context,
	auctionEntity *auction_entity.Auction,
) *internal_error.InternalError {

	auctionEntityMongo := &AuctionEntityMongo{
		Id:          auctionEntity.Id,
		ProductName: auctionEntity.ProductName,
		Category:    auctionEntity.Category,
		Description: auctionEntity.Description,
		Condition:   auctionEntity.Condition,
		Status:      auctionEntity.Status,
		Timestamp:   auctionEntity.Timestamp.Unix(),
	}

	_, err := ar.Collection.InsertOne(ctx, auctionEntityMongo)
	if err != nil {
		logger.Error("Error trying to insert action", err)
		return internal_error.NewInternalServerError("Error trying to insert action")
	}
	return nil
}
