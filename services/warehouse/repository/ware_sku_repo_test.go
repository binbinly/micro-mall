package repository

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"pkg/app"
	"warehouse/config"
)

var repo IRepo

func TestMain(m *testing.M) {
	if err := app.LoadEnv(config.Cfg); err != nil {
		panic(err)
	}
	repo = New(orm.NewDB(&config.Cfg.MySQL), cache.NewRedisCache(redis.InitTestRedis()))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_GetWareSkuStock(t *testing.T) {
	type args struct {
		ctx   context.Context
		skuID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetWareSkuStock",
			args: args{
				ctx:   context.TODO(),
				skuID: 13,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStock, err := repo.GetWareSkuStock(tt.args.ctx, tt.args.skuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetWareSkuStock() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("stock: %v", gotStock)
		})
	}
}

func TestRepo_BatchGetWareSkuStocks(t *testing.T) {
	list, err := repo.BatchGetWareSkusStock(context.TODO(), []int64{13, 15})
	if err != nil {
		t.Errorf("BatchGetWareSkuStocks() error = %v", err)
		return
	}
	t.Logf("list len:%v", len(list))
}

func TestRepo_BatchGetWareSkus(t *testing.T) {
	list, err := repo.BatchGetWareSkus(context.TODO(), []int64{13, 15})
	if err != nil {
		t.Errorf("BatchGetWareSkus() error = %v", err)
		return
	}
	t.Logf("list len:%v", len(list))
}
