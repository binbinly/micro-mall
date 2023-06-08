package repository

import (
	"context"
	"market/model"
	"testing"
)

var repo IRepo

//func TestMain(m *testing.M) {
//	redis.InitTestRedis()
//	orm.InitTest("mall_sms")
//	repo = New(orm.GetDB(), util.NewCache())
//	if code := m.Run(); code != 0 {
//		panic(code)
//	}
//}

func TestRepo_GetConfigByName(t *testing.T) {
	type fields struct {
		Repo IRepo
	}
	type args struct {
		ctx  context.Context
		name string
		v    interface{}
	}

	fs := fields{Repo: repo}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   model.ConfigKeyHomeCat,
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: model.ConfigKeyHomeCat,
				v:    []*model.ConfigHomeCat{},
			},
			wantErr: false,
		},
		{
			name:   model.ConfigKeyPayList,
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: model.ConfigKeyPayList,
				v:    []*model.ConfigPayList{},
			},
			wantErr: false,
		},
		{
			name:   "test",
			fields: fs,
			args: args{
				ctx:  context.Background(),
				name: "test",
				v:    nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repo.GetConfigByName(tt.args.ctx, tt.args.name, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("GetConfigByName() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
