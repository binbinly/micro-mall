package repository

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/cache"
	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"
	"github.com/stretchr/testify/assert"

	"center/config"
	"center/model"
	"pkg/app"
	"pkg/dbs"
)

var r IRepo

func TestMain(m *testing.M) {
	if err := app.LoadEnv(config.Cfg); err != nil {
		panic(err)
	}
	r = New(orm.NewDB(&config.Cfg.MySQL), cache.NewRedisCache(redis.InitTestRedis()))
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestRepo_UserCreate(t *testing.T) {
	type args struct {
		ctx  context.Context
		user *model.UserModel
	}
	tests := []struct {
		name    string
		args    args
		wantID  int64
		wantErr bool
	}{
		{"zhangsan",
			args{
				ctx: context.Background(),
				user: &model.UserModel{
					PriID:    dbs.PriID{ID: 1},
					Username: "zhangsan",
					Password: "123456",
					Phone:    15555555555,
					Status:   model.StatusNormal,
				},
			},
			1,
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotID, err := r.UserCreate(tt.args.ctx, tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserCreate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotID != tt.wantID {
				t.Errorf("UserCreate() gotId = %v, want %v", gotID, tt.wantID)
			}
		})
	}
}

func TestRepo_UserUpdate(t *testing.T) {
	err := r.UserUpdate(context.Background(), 1, map[string]any{
		"nickname": "张三",
	})
	assert.NoError(t, err)
}

func TestRepo_GetUserByID(t *testing.T) {
	user, err := r.GetUserByID(context.Background(), 1)
	assert.NoError(t, err)
	t.Logf("user:%v", user)
	if user != nil {
		assert.Equal(t, user.ID, int64(1))
	}
}

func TestRepo_GetUserByUsername(t *testing.T) {
	user, err := r.GetUserByUsername(context.Background(), "zhangsan")
	assert.NoError(t, err)
	t.Logf("user:%v", user)
	if user != nil {
		assert.Equal(t, user.ID, int64(1))
	}
}
