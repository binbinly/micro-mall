package app

import (
	"fmt"
	"math/rand"
	"time"
)

// GenOrderNo 生成订单号
func GenOrderNo() string {
	now := time.Now()
	return fmt.Sprintf("%04d%02d%02d%02d%02d%02d%08d", now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second(), rand.Int31n(100000000))
}
