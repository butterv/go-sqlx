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

func TestUserService_GetUser(t *testing.T) {
	want := &pb.GetUserResponse{
		User: &pb.User{
			UserId: "TEST_USER_ID",
			Email:  "TEST_EMAIL",
		},
	}

	now := time.Now()
	s := test.NewStore()
	s.AddUsers(&model.User{
		ID:        "TEST_USER_ID",
		Email:     "TEST_EMAIL",
		CreatedAt: now,
		UpdatedAt: now,
	})

	r := inmemory.New(s)
	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: "TEST_USER_ID",
	}

	got, err := service.GetUser(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}

func TestUserService_GetUser_NotFound(t *testing.T) {
	want := &pb.GetUserResponse{}

	s := test.NewStore()
	r := inmemory.New(s)
	service := user.NewUserService(r, model.NewDefaultUserIDGenerator())

	ctx := context.Background()
	req := &pb.GetUserRequest{
		UserId: "TEST_USER_ID",
	}

	got, err := service.GetUser(ctx, req)
	if err != nil {
		t.Fatalf("service.GetUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.GetUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}
}
