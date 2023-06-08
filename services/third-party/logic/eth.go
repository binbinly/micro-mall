package logic

import (
	"context"
	"pkg/errno"

	"github.com/pkg/errors"
)

// CheckPay 检查订单是否已支付
func (l *logic) CheckPay(ctx context.Context, id int64, address, orderNo string) error {
	//连接合约
	if err := l.contract.Connect(id, address); err != nil {
		return errors.Wrapf(err, "[third.eth] connect contract address %v", address)
	}
	//调用合约
	check, err := l.contract.CheckPay(id, orderNo)
	if err != nil {
		return errors.Wrapf(err, "[third.eth] contract call checkPay")
	}
	if !check { //未支付
		return errno.ErrETHPayNotFound
	}

	return nil
}
