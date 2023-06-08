package logic

import (
	"context"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/redis/go-redis/v9"
	"third-party/eth"
)

var _ Logic = (*logic)(nil)

// Logic 第三方服务接口定义
type Logic interface {
	SendSMS(ctx context.Context, phone string) (string, error)
	CheckVCode(ctx context.Context, phone int64, vCode string) error
	CheckPay(ctx context.Context, id int64, address, orderNo string) error

	Close() error
}

type logic struct {
	contract *eth.Contract
	rdb      *redis.Client
}

// New init logic
func New(rdb *redis.Client, c *ethclient.Client) Logic {
	return &logic{
		contract: eth.NewContract(c),
		rdb:      rdb,
	}
}

// Close logic
func (l *logic) Close() error {
	l.contract.Close()

	return l.rdb.Close()
}
