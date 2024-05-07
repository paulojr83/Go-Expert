package bid_controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/rest_err"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/infra/api/web/validation"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/usecase/bid_usecase"
	"net/http"
)

type BidController struct {
	bidUseCase bid_usecase.BidRepositoryInterface
}

func NewBidController(bidUseCase bid_usecase.BidRepositoryInterface) *BidController {
	return &BidController{bidUseCase}
}

func (bid *BidController) CreateAuction(c *gin.Context) {
	var bidInput bid_usecase.BidInputDto

	if err := c.ShouldBindJSON(&bidInput); err != nil {
		errRest := validation.ValidateErr(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	err := bid.bidUseCase.CreateBid(context.Background(), bidInput)
	if err != nil {
		errRest := rest_err.ConvertError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	c.Status(http.StatusCreated)
}
