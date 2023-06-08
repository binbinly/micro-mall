package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSpuByID(t *testing.T) {
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
			name: "GetSpuByID",
			args: args{
				ctx: context.Background(),
				id:  129,
			},
		},
		{
			name: "GetSpuByID_empty",
			args: args{
				ctx: context.Background(),
				id:  1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotSpu, err := repo.GetSpuByID(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpuByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotSpu == nil {
				t.Logf("%d empty", tt.args.id)
				return
			}
			t.Logf("id:%v", gotSpu.ID)
		})
	}
}
