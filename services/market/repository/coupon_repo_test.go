package repository

import (
	"context"
	"testing"
)

func TestRepo_GetCouponList(t *testing.T) {
	type args struct {
		ctx   context.Context
		catID int64
		spuID int64
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "GetCouponList",
			args: args{
				ctx: context.TODO(),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := repo.GetCouponList(tt.args.ctx, tt.args.catID, tt.args.spuID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCouponList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("got len: %v", len(gotList))
		})
	}
}

func TestRepo_GetCouponByID(t *testing.T) {
	coupon, err := repo.GetCouponByID(context.Background(), 1)
	if err != nil {
		t.Errorf("GetCouponList() error = %v", err)
		return
	}
	t.Logf("coupon:%v", coupon)
}

func TestRepo_GetCouponMemberList(t *testing.T) {
	list, err := repo.GetCouponMemberList(context.Background(), 1)
	if err != nil {
		t.Errorf("GetCouponMemberList() error = %v", err)
		return
	}
	t.Logf("list len:%v", len(list))
}
