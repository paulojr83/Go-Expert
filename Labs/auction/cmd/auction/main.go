package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/database/mongodb"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/api/web/controller/auction_cotroller"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/api/web/controller/bid_controller"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/api/web/controller/user_controller"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/database/auction"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/database/bid"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/database/user"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/auction_usecase"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/bid_usecase"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/user_usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

func main() {
	/*if err := godotenv.Load("cmd/auction/.env"); err != nil {
		log.Fatalf("Error trying to load env variables %w", err)
		return
	}*/
	ctx := context.Background()

	mongoDBConnection, err := mongodb.NewMongoDBConnection(ctx)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}

	router := gin.Default()

	userController,
		bidController,
		auctionController := initDependencies(mongoDBConnection)

	router.GET("/auctions", auctionController.FindAuctions)
	router.POST("/auctions", auctionController.CreateAuction)
	router.GET("/auctions/winner/:auctionId", auctionController.FindAuctionById)

	router.GET("/bid/:auctionId", bidController.FindAuctionById)
	router.POST("/bid", bidController.CreateAuction)

	router.GET("/users/:userId", userController.FindUserById)

	PORT := os.Getenv("HTTP_PORT")
	err = router.Run(PORT)
	log.Println("Starting server on port: ", PORT)
	if err != nil {
		panic(err)
		return
	}
}

func initDependencies(dataBase *mongo.Database) (
	userController *user_controller.UserController,
	bidController *bid_controller.BidController,
	auctionController *auction_cotroller.AuctionController,
) {
	auctionRepository := auction.NewAuctionRepository(dataBase)
	bidRepository := bid.NewBidRepository(dataBase, auctionRepository)
	userRepository := user.NewUserRepository(dataBase)

	userController = user_controller.NewUserController(user_usecase.NewUserUseCase(userRepository))
	auctionController = auction_cotroller.NewAuctionController(auction_usecase.NewAuctionUseCase(auctionRepository, bidRepository))
	bidController = bid_controller.NewBidController(bid_usecase.NewBidUseCase(bidRepository))

	return userController, bidController, auctionController
}
