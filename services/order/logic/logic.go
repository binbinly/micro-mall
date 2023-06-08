package logic

import (
	"context"
	"github.com/binbinly/pkg/cache"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	"gorm.io/gorm"
	"order/model"
	"pkg/constvar"
	"pkg/proto/cart"
	"pkg/proto/core"
	"pkg/proto/market"
	"pkg/proto/member"
	"pkg/proto/product"
	"pkg/proto/warehouse"

	"github.com/redis/go-redis/v9"
	"order/repository"
)

var _ Logic = (*logic)(nil)

// Logic 订单服务接口
type Logic interface {
	OrderSubmit(ctx context.Context, memberID, addressID, couponID int64, skuIds []int64, note string) error
	SubmitSkuOrder(ctx context.Context, memberID, skuID, addressID, couponID int64, num int, note string) error
	SubmitKillOrder(ctx context.Context, memberID, skuID, addressID int64, price, num int, orderNo string) error
	OrderDetail(ctx context.Context, mid, id int64) (*model.OrderModel, error)
	OrderCancel(ctx context.Context, mid, id int64) error
	MyOrderList(ctx context.Context, mid int64, status int, page int32) ([]*model.OrderModel, error)
	OrderPayNotify(ctx context.Context, mid int64, amount int, pType int64, orderNo, tradeNo, transHash string) error
	OrderConfirmReceipt(ctx context.Context, mid, id int64) error
	OrderRefund(ctx context.Context, mid, id int64, content string) error
	OrderComment(ctx context.Context, mid, id int64, skuIds []int64, star int32, content, resources string) error
	OrderInfo(ctx context.Context, id int64) (*model.OrderModel, error)

	Close() error
}

// logic 营销服务
type logic struct {
	repo             repository.IRepo
	rdb              *redis.Client
	releaseEvent     micro.Event
	createEvent      micro.Event
	cartService      cart.CartService
	memberService    member.MemberService
	marketService    market.MarketService
	warehouseService warehouse.WarehouseService
	productService   product.ProductService
	thirdService     core.ThirdPartyService
}

// New init logic
func New(db *gorm.DB, rdb *redis.Client, c client.Client) Logic {
	return &logic{
		rdb:              rdb,
		repo:             repository.New(db, cache.NewRedisCache(rdb)),
		releaseEvent:     micro.NewEvent(constvar.TopicOrderRelease, c),
		createEvent:      micro.NewEvent(constvar.KeyOrderCreate, c),
		cartService:      cart.NewCartService(constvar.ServiceCart, c),
		memberService:    member.NewMemberService(constvar.ServiceMember, c),
		marketService:    market.NewMarketService(constvar.ServiceMarket, c),
		warehouseService: warehouse.NewWarehouseService(constvar.ServiceWarehouse, c),
		productService:   product.NewProductService(constvar.ServiceProduct, c),
		thirdService:     core.NewThirdPartyService(constvar.ServiceThird, c),
	}
}

// Close logic
func (l *logic) Close() error {
	return l.repo.Close()
}
