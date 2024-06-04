package auction

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/auction_entity"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
	"os"
	"testing"
	"time"
)

func TestAuctionClosesAutomatically(t *testing.T) {
	os.Setenv("AUCTION_DURATION_SECONDS", "2") // Define o tempo de duração para 2 segundos

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()
	mt.Run("test auction closes automatically", func(mt *mtest.T) {

		ctx := context.Background()
		db := mt.DB
		auctionRepo := NewAuctionRepository(db)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		auction, _ := auction_entity.NewAuction("Test", "Test", "Descrição teste para novo leilão", auction_entity.ProductionCondition(0))
		err := auctionRepo.CreateAuction(ctx, auction)
		if err != nil {
			t.Fatalf("Failed to create auction: %s", err)
		}

		time.Sleep(3 * time.Second) // Espera tempo suficiente para o leilão fechar

		var result auction_entity.Auction
		filter := bson.M{"_id": auction.Id}
		if err := db.Collection("auctions").FindOne(ctx, filter).Decode(&result).Error(); err != "no responses remaining" {
			assert.Nil(t, err)
			if result.Status != auction_entity.Completed {
				t.Fatalf("Auction %s should be closed but is still open", auction.Id)
			}
		}
	})
}
