package auction

import (
	"context"
	"fmt"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"strconv"
	"time"
)

// AuctionEntityMongo representa a estrutura do leilão no MongoDB
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
	return &AuctionRepository{
		Collection: database.Collection("auctions"),
	}
}

func (ar *AuctionRepository) CreateAuction(
	ctx context.Context,
	auctionEntity *auction_entity.Auction) *internal_error.InternalError {
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
		logger.Error("Error trying to insert auction", err)
		return internal_error.NewInternalServerError("Error trying to insert auction")
	}

	// Iniciar goroutine para fechamento automático do leilão
	go ar.startAuctionTimer(ctx, auctionEntityMongo.Id)

	return nil
}

func (ar *AuctionRepository) startAuctionTimer(ctx context.Context, auctionID string) {
	duration := getAuctionDuration()
	time.Sleep(duration)
	ar.closeAuction(ctx, auctionID)
}

func getAuctionDuration() time.Duration {
	durationStr := os.Getenv("AUCTION_DURATION_SECONDS")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		log.Fatalf("Invalid auction duration: %s", err)
	}
	return time.Duration(duration) * time.Second
}

func (ar *AuctionRepository) closeAuction(ctx context.Context, auctionID string) {
	filter := bson.M{"_id": auctionID}
	update := bson.M{"$set": bson.M{"status": auction_entity.Completed}}

	_, err := ar.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		logger.Error("Error trying to close auction", err)
	} else {
		logger.Info(fmt.Sprintf("Auction closed successfully %s", auctionID))
	}
}
