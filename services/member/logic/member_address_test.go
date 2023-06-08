package logic

import (
	"context"
	"testing"
)

func TestService_MemberAddressAdd(t *testing.T) {
	got, err := srv.MemberAddressAdd(context.Background(), 3, "张三", "15555555555",
		"北京市", "北京市", "朝阳区", "朝阳广场东侧", 123123, 0)
	if err != nil {
		t.Errorf("MemberAddressAdd() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_MemberAddressEdit(t *testing.T) {
	err := srv.MemberAddressEdit(context.Background(), 1, 3, "张三", "15555555555",
		"北京市", "北京市", "朝阳区", "朝阳广场东侧", 123123, 1)
	if err != nil {
		t.Errorf("MemberAddressEdit() error = %v", err)
		return
	}
}

func TestService_MemberAddressList(t *testing.T) {
	got, err := srv.MemberAddressList(context.Background(), 3)
	if err != nil {
		t.Errorf("MemberAddressAdd() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestService_MemberAddressDel(t *testing.T) {
	err := srv.MemberAddressDel(context.Background(), 1, 3)
	if err != nil {
		t.Errorf("MemberAddressDel() error = %v", err)
		return
	}
}

func TestService_GetMemberAddressInfo(t *testing.T) {
	got, err := srv.GetMemberAddressInfo(context.Background(), 1, 3)
	if err != nil {
		t.Errorf("GetMemberAddressInfo() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
