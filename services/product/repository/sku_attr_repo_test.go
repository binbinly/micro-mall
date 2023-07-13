package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSkuAttrsBySpuID(t *testing.T) {
	type args struct {
		ctx   context.Context
		spuID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetSkuAttrsBySpuID",
			args: args{
				ctx:   context.Background(),
				spuID: 129,
			},
		},
		{
			name: "GetSkuAttrsBySpuID_empty",
			args: args{
				ctx:   context.Background(),
				spuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetSkuAttrsBySpuID(tt.args.ctx, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSkuAttrsBySpuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(gotList))
		})
	}
}
