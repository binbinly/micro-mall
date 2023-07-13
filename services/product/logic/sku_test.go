package logic

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
	"product/config"
)

var srv Logic

func TestMain(m *testing.M) {
	// load config
	if err := app.LoadEnv(config.Cfg); err != nil {
		panic(err)
	}

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	srv = New(nil, db, redis.InitTestRedis())
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_SkuDetail(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SkuDetail",
			args: args{
				ctx: context.Background(),
				id:  249700,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := srv.SkuDetail(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("SkuDetail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestService_GetSkuSaleAttrs(t *testing.T) {
	got, err := srv.GetSkuSaleAttrs(context.Background(), 12900)
	if err != nil {
		t.Errorf("SkuDetail() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_GetSkuInfo(t *testing.T) {
	got, err := srv.GetSkuInfo(context.Background(), 12900)
	if err != nil {
		t.Errorf("GetSkuInfo() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_SpuComment(t *testing.T) {
	err := srv.SpuComment(context.Background(), []int64{12900}, 1, 1, 1, "test", "")
	if err != nil {
		t.Errorf("SpuComment() error = %v", err)
		return
	}
}

func TestService_CategoryTree(t *testing.T) {
	got, err := srv.CategoryTree(context.Background())
	if err != nil {
		t.Errorf("CategoryTree() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
