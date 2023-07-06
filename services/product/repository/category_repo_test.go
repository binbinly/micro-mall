package repository

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
	"pkg/constvar"
	"product/config"
)

var repo IRepo

func TestMain(m *testing.M) {
	if err := app.LoadEnv(constvar.ServiceProduct, config.Cfg); err != nil {
		panic(err)
	}
	repo = New(orm.NewDB(&config.Cfg.MySQL), cache.NewRedisCache(redis.InitTestRedis()))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_GetCategoryTree(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetCategoryTree",
			args: args{ctx: context.Background()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.CategoryAll(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GoodsCategoryTree() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(got))
		})
	}
}

func TestRepo_GetCategoryNameByID(t *testing.T) {
	name, err := repo.GetCategoryNameByID(context.Background(), 893)
	if err != nil {
		t.Errorf("GetCategoryNameByID() error = %v", err)
		return
	}
	t.Logf("name: %v", name)
}

func TestRepo_GetCategoryNamesByIds(t *testing.T) {
	names, err := repo.GetCategoryNamesByIds(context.Background(), []int64{893, 896, 897})
	if err != nil {
		t.Errorf("GetCategoryNamesByIds() error = %v", err)
		return
	}
	t.Logf("name len: %#v", names)
}
