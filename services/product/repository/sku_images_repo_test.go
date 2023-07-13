package repository

import (
	"context"
	"testing"
)

func TestRepo_GetImagesBySkuID(t *testing.T) {
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
			name: "GetImagesBySkuID",
			args: args{
				ctx: context.Background(),
				id:  2636000,
			},
		},
		{
			name: "GetImagesBySkuID_empty",
			args: args{
				ctx: context.Background(),
				id:  0,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetImagesBySkuID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetImagesBySkuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(gotList))
		})
	}
}
