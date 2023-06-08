package repository

import (
	"context"
	"testing"
)

func TestRepo_GetSpuDescBySpuID(t *testing.T) {
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
			name: "GetSpuDescBySpuID",
			args: args{
				ctx:   context.Background(),
				spuID: 129,
			},
		},
		{
			name: "GetSpuDescBySpuID_empty",
			args: args{
				ctx:   context.Background(),
				spuID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDesc, err := repo.GetSpuDescBySpuID(tt.args.ctx, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetSpuDescBySpuID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotDesc == nil {
				t.Logf("%d empty", tt.args.spuID)
				return
			}
			t.Logf("id:%v", gotDesc.SpuID)
		})
	}
}
