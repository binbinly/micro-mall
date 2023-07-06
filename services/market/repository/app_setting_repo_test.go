package repository

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"market/model"
)

func TestRepo_AppPageData(t *testing.T) {
	type fields struct {
		Repo IRepo
	}
	type args struct {
		ctx  context.Context
		page int
	}
	fs := fields{Repo: repo}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "home_page",
			fields: fs,
			args: args{
				ctx:  context.Background(),
				page: model.AppPageHome,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotList, err := repo.AppPageData(tt.args.ctx, tt.args.page)
			if (err != nil) != tt.wantErr {
				t.Errorf("AppPageData() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			t.Logf("list len: %v", len(gotList))
		})
	}
}

func TestRepo_AppHomePageData(t *testing.T) {
	list, err := repo.AppHomePageData(context.Background(), 1)
	assert.NoError(t, err)
	t.Logf("list len: %v", len(list))
}
