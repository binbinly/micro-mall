package repository

import (
	"context"
	"testing"
)

func TestRepo_GetAttrGroupByCatID(t *testing.T) {
	type args struct {
		ctx   context.Context
		catID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetAttrGroupByCatID",
			args: args{
				ctx:   context.Background(),
				catID: 896,
			},
		},
		{
			name: "GetAttrGroupByCatID_empty",
			args: args{
				ctx:   context.Background(),
				catID: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := r.GetAttrGroupByCatID(tt.args.ctx, tt.args.catID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAttrGroupByCatID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len:%v", len(gotList))
		})
	}
}
