package user_test

import (
	"context"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/domain/service/user"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func TestUserService_GetUsers(t *testing.T) {
	want := &pb.GetUsersResponse{
		Users: []*pb.User{
			{
				UserId: "TEST_USER_ID1",
				Email:  "TEST_EMAIL1",
			},
			{
				UserId: "TEST_USER_ID2",
				Email:  "TEST_EMAIL2",
			},
			{
				UserId: "TEST_USER_ID3",
				Email:  "TEST_EMAIL3",
			},
		},
	}

	now := time.Now()
	s := test.NewStore()
	s.AddUsers([]*model.User{
		{
			ID:        "TEST_USER_ID1",
			Email:     "TEST_EMAIL1",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        "TEST_USER_ID2",
			Email:     "TEST_EMAIL2",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        "TEST_USER_ID3",
			Email:     "TEST_EMAIL3",
			CreatedAt: now,
			UpdatedAt: now,
		},
		{
			ID:        "TEST_USER_ID4",
			Email:     "TEST_EMAIL4",
			CreatedAt: now,
			UpdatedAt: now,
		},
	}...)

	r := inmemory.New(s)
	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUsersRequest{
		UserIds: []string{"TEST_USER_ID1", "TEST_USER_ID2", "TEST_USER_ID3"},
	}

	got, err := service.GetUsers(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUsers(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUsers(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUsers_NotFound(t *testing.T) {
	want := &pb.GetUsersResponse{}

	s := test.NewStore()
	r := inmemory.New(s)
	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUsersRequest{
		UserIds: []string{"TEST_USER_ID1", "TEST_USER_ID2", "TEST_USER_ID3"},
	}

	got, err := service.GetUsers(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUsers(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUsers(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}
