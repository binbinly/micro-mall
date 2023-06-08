package eth

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Contract 合约连接结构
type Contract struct {
	ctx      context.Context
	client   *ethclient.Client
	Payments map[int64]*Payment
}

// NewContract 创建合约
func NewContract(c *ethclient.Client) *Contract {
	return &Contract{
		ctx:      context.Background(),
		client:   c,
		Payments: make(map[int64]*Payment),
	}
}

// Connect 连接合约地址
func (c *Contract) Connect(id int64, address string) error {
	if _, ok := c.Payments[id]; ok {
		return nil
	}
	pay, err := NewPayment(common.HexToAddress(address), c.client)
	if err != nil {
		return err
	}
	c.Payments[id] = pay
	return nil
}

// CheckPay 验证是否已支付
func (c *Contract) CheckPay(id int64, orderNo string) (bool, error) {
	if pay, ok := c.Payments[id]; ok {
		return pay.CheckPay(nil, orderNo)
	}
	return false, nil
}

// Close client
func (c *Contract) Close() {
	c.client.Close()
}
