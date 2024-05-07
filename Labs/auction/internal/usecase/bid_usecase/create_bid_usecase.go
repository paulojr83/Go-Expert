package bid_usecase

import (
	"context"
	"github.com/paulojr83/Go-Expert/Labs/auction/configuration/logger"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/entity/bid_entity"
	"github.com/paulojr83/Go-Expert/Labs/auction/internal/internal_error"
	"os"
	"strconv"
	"time"
)

type BidOutputDto struct {
	Id        string    `json:"id"`
	UserId    string    `json:"user_id"`
	AuctionId string    `json:"auction_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp" time_format:"2006-01-02 15:10:00"`
}

type BidInputDto struct {
	UserId    string  `json:"user_id"`
	AuctionId string  `json:"auction_id"`
	Amount    float64 `json:"amount"`
}

type BidUseCase struct {
	BidRepository       bid_entity.BidRepositoryInterface
	timer               *time.Timer
	maxBatchSize        int
	batchInsertInterval time.Duration
	bidChannel          chan bid_entity.Bid
}

func NewBidUseCase(bidRepository bid_entity.BidRepositoryInterface) BidRepositoryInterface {
	maxSizeInterval := getMaxBatchSizeInterval()
	timer := time.NewTimer(maxSizeInterval)
	maxBatchSize := getMaxBatchSize()
	bidUseCase := &BidUseCase{
		BidRepository:       bidRepository,
		timer:               timer,
		maxBatchSize:        maxBatchSize,
		batchInsertInterval: maxSizeInterval,
		bidChannel:          make(chan bid_entity.Bid, maxBatchSize),
	}

	bidUseCase.triggerCreateRouting(context.Background())

	return bidUseCase
}

type BidRepositoryInterface interface {
	CreateBid(ctx context.Context, bidEntity BidInputDto) *internal_error.InternalError
	FindBidByActionId(ctx context.Context, actionId string) ([]BidOutputDto, *internal_error.InternalError)
	FindWinningBidByActionId(ctx context.Context, actionId string) (*BidOutputDto, *internal_error.InternalError)
}

var bidBatch []bid_entity.Bid

func (bu *BidUseCase) triggerCreateRouting(ctx context.Context) {
	go func() {
		defer close(bu.bidChannel)

		for {
			select {
			case bidEntity, ok := <-bu.bidChannel:
				if !ok {
					if len(bidBatch) > 0 {
						if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
							logger.Error("Error trying to process bid batch list", err)
						}
					}
					return
				}

				bidBatch = append(bidBatch, bidEntity)

				if len(bidBatch) >= bu.maxBatchSize {
					if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
						logger.Error("Error trying to process bid batch list", err)
					}

					bidBatch = nil
					bu.timer.Reset(bu.batchInsertInterval)
				}
			case <-bu.timer.C:
				if err := bu.BidRepository.CreateBid(ctx, bidBatch); err != nil {
					logger.Error("Error trying to process bid batch list", err)
				}
				bidBatch = nil
				bu.timer.Reset(bu.batchInsertInterval)
			}
		}
	}()
}
func (bu *BidUseCase) CreateBid(ctx context.Context,
	bidInputDto BidInputDto) *internal_error.InternalError {

	bidEntity, err := bid_entity.NewBid(bidInputDto.UserId, bidInputDto.AuctionId, bidInputDto.Amount)
	if err != nil {
		return err
	}
	bu.bidChannel <- *bidEntity

	return nil
}

func getMaxBatchSizeInterval() time.Duration {
	batchInsertInterval := os.Getenv("BATCH_INSERT_INTERVAL")
	duration, err := time.ParseDuration(batchInsertInterval)
	if err != nil {
		return 3 * time.Minute
	}
	return duration
}

func getMaxBatchSize() int {
	value, err := strconv.Atoi(os.Getenv("MAX_BATCH_SIZE"))
	if err != nil {
		return 5
	}
	return value
}
