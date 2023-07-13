package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSkuByID(t *testing.T) {
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
			name: "GetSkuByID",
			args: args{
				ctx: context.Background(),
				id:  2636000,
			},
		},
		{
			name: "GetSkuByID_empty",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSku, err := r.GetSkuByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSku == nil {
				t.Logf("%d empty", tt.args.id)
				return
			}
			t.Logf("id:%v", gotSku.ID)
		})
	}
}

func TestRepo_GetSkusBySpuID(t *testing.T) {
	list, err := r.GetSkusBySpuID(context.Background(), 129)
	if err != nil {
		t.Errorf("GetSkusBySpuID() error = %v", err)
		return
	}
	t.Logf("got len: %v", len(list))
}

func TestRepo_GetSkusByIds(t *testing.T) {
	list, err := r.GetSkusByIds(context.Background(), []int64{12900, 12901, 3})
	if err != nil {
		t.Errorf("GetSkusByIds() error = %v", err)
		return
	}
	t.Logf("got len: %v", len(list))
}
