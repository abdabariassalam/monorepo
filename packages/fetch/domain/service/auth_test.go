package service

import (
	"reflect"
	"testing"

	"github.com/bariasabda/monorepo/packages/fetch/config"
	"github.com/bariasabda/monorepo/packages/fetch/domain/repository"
	"github.com/bariasabda/monorepo/packages/fetch/domain/repository/mock"
	"github.com/golang/mock/gomock"
)

func Test_service_VerifyToken(t *testing.T) {
	type fields struct {
		cfg  config.Config
		repo repository.RepositoryInterface
	}
	type args struct {
		reqToken string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			"test1 return error",
			fields{
				cfg:  config.Config{},
				repo: mock.NewMockRepositoryInterface(gomock.NewController(t)),
			},
			args{
				reqToken: "try",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &service{
				cfg:  tt.fields.cfg,
				repo: tt.fields.repo,
			}
			got, err := s.VerifyToken(tt.args.reqToken)
			if (err != nil) != tt.wantErr {
				t.Errorf("service.VerifyToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("service.VerifyToken() = %v, want %v", got, tt.want)
			}
		})
	}
}
