package repository

import (
	"context"
	"testing"
)

func TestRepo_GetBrandByID(t *testing.T) {
	got, err := repo.GetBrandByID(context.Background(), 271)
	if err != nil {
		t.Errorf("GetBrandByID() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}

func TestRepo_GetBrandsByIds(t *testing.T) {
	got, err := repo.GetBrandsByIds(context.Background(), []int64{271, 274, 276})
	if err != nil {
		t.Errorf("GetBrandsByIds() error = %v", err)
		return
	}
	t.Logf("got: %v", got)
}
