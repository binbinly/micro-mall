package logic

import (
	"context"
	"testing"

	"github.com/binbinly/pkg/storage/orm"
	"github.com/binbinly/pkg/storage/redis"

	"member/config"
	"pkg/app"
)

var srv Logic

func TestMain(m *testing.M) {
	// load config
	if err := app.LoadEnv(config.Cfg); err != nil {
		panic(err)
	}

	// init dbs
	db := orm.NewDB(&config.Cfg.MySQL)

	srv = New(db, redis.InitTestRedis())
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestService_MemberEdit(t *testing.T) {
	err := srv.MemberEdit(context.Background(), 3, map[string]any{"nickname": "测试"})
	if err != nil {
		t.Errorf("MemberEdit() error = %v", err)
		return
	}
}

func TestService_MemberInfo(t *testing.T) {
	got, err := srv.MemberInfo(context.Background(), 3)
	if err != nil {
		t.Errorf("MemberInfo() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
