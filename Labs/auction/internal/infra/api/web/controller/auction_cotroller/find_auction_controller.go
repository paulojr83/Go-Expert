package auction_cotroller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/rest_err"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/auction_usecase"
	"net/http"
	"strconv"
)

func (ac *AuctionController) FindWinningBidByAuctionId(c *gin.Context) {
	auctionId := c.Query("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := ac.auctionUseCase.FindWinningBidByAuctionId(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}
func (ac *AuctionController) FindAuctions(c *gin.Context) {
	status := c.Query("status")
	category := c.Query("category")
	productName := c.Query("productName")

	statusNumber, errConv := strconv.Atoi(status)
	if errConv != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "status",
			Message: "Error trying to validate action status param",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctions, err := ac.auctionUseCase.FindAuctions(context.Background(),
		auction_usecase.AuctionStatus(statusNumber),
		category,
		productName)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctions)
}
func (ac *AuctionController) FindAuctionById(c *gin.Context) {
	auctionId := c.Query("auctionId")

	if err := uuid.Validate(auctionId); err != nil {
		errRest := rest_err.NewBadRequestError("Invalid fields", rest_err.Causes{
			Field:   "auctionId",
			Message: "Invalid UUID value",
		})
		c.JSON(errRest.Code, errRest)
		return
	}

	auctionData, err := ac.auctionUseCase.FindAuctionById(context.Background(), auctionId)
	if err != nil {
		errRest := rest_err.ConvertError(err)
		c.JSON(errRest.Code, errRest)
		return
	}

	c.JSON(http.StatusOK, auctionData)
}
