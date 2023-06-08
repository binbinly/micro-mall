package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSkuByID(t *testing.T) {
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
			name: "GetSkuByID",
			args: args{
				ctx:   context.Background(),
				skuID: 1300,
			},
		},
		{
			name: "GetSkuByID_1",
			args: args{
				ctx:   context.Background(),
				skuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := repo.GetSkuByID(tt.args.ctx, tt.args.skuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got: %v", got)
		})
	}
}

func TestRepo_GetSkusBySessionID(t *testing.T) {
	got, err := repo.GetSkusBySessionID(context.Background(), 1)
	if err != nil {
		t.Errorf("GetSkusBySessionID() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
