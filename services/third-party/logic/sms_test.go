package logic

import (
	"context"
	"strconv"
	"testing"

	"github.com/binbinly/pkg/storage/redis"

	"pkg/app"
	"pkg/constvar"
	"third-party/config"
)

var srv Logic

func TestMain(m *testing.M) {
	// load config
	if err := app.LoadEnv(constvar.ServiceWarehouse, config.Cfg); err != nil {
		panic(err)
	}

	rdb := redis.InitTestRedis()

	srv = New(rdb, nil)
	if code := m.Run(); code != 0 {
		panic(code)
	}
}

func TestThird_SendSMS(t1 *testing.T) {
	type args struct {
		ctx   context.Context
		phone string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "SendSMS",
			args: args{
				ctx:   context.Background(),
				phone: "15555555555",
			},
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			got, err := srv.SendSMS(tt.args.ctx, tt.args.phone)
			if (err != nil) != tt.wantErr {
				t1.Errorf("SendSMS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t1.Logf("got: %v", got)
			phone, _ := strconv.ParseInt(tt.args.phone, 10, 64)
			err = srv.CheckVCode(context.Background(), phone, got)
			if (err != nil) != tt.wantErr {
				t1.Errorf("CheckVCode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
