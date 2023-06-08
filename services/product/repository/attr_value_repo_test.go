package repository

import (
	"context"
	"testing"
)

func TestRepo_GetAttrsBySpuID(t *testing.T) {
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
			name: "GetAttrsBySpuID",
			args: args{
				ctx:   context.Background(),
				spuID: 39559,
			},
		},
		{
			name: "GetAttrsBySpuID_empty",
			args: args{
				ctx:   context.Background(),
				spuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := repo.GetAttrsBySpuID(tt.args.ctx, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttrsBySpuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(gotList))
		})
	}
}
