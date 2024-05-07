package auction_cotroller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/rest_err"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/api/web/validation"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/auction_usecase"
	"net/http"
)

type AuctionController struct {
	auctionUseCase auction_usecase.AuctionUseCaseInterface
}

func NewAuctionController(auctionUseCase auction_usecase.AuctionUseCaseInterface) *AuctionController {
	return &AuctionController{auctionUseCase}
}

func (ac *AuctionController) CreateAuction(c *gin.Context) {
	var auctionInput auction_usecase.AuctionInputDto

	if err := c.ShouldBindJSON(&auctionInput); err != nil {
		errRest := validation.ValidateErr(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	err := ac.auctionUseCase.CreateAuction(context.Background(), auctionInput)
	if err != nil {
		errRest := rest_err.ConvertError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	c.Status(http.StatusCreated)
}
