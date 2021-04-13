package user_test

import (
	"context"
	"testing"

	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/domain/service/user"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory"
	"github.com/butterv/go-sqlx/app/infrastructure/inmemory/test"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

type testUserIDGenerator struct {
	userID model.UserID
}

func newTestUserIDGenerator(uID model.UserID) model.UserIDGenerator {
	return &testUserIDGenerator{
		userID: uID,
	}
}

func (g *testUserIDGenerator) Generate() model.UserID {
	return g.userID
}

func TestUserService_CreateUser(t *testing.T) {
	want := &pb.CreateUserResponse{
		UserId: "TEST_USER_ID",
	}

	s := test.NewStore()
	r := inmemory.New(s)

	uID := model.UserID("TEST_USER_ID")
	userIDGenerator := newTestUserIDGenerator(uID)
	service := user.NewUserService(r, userIDGenerator)

	ctx := context.Background()
	req := &pb.CreateUserRequest{
		Email: "TEST_EMAIL",
	}

	got, err := service.CreateUser(ctx, req)
	if err != nil {
		t.Fatalf("service.CreateUser(ctx, %v) = _, %#v; want nil", req, err)
	}
	if diff := cmp.Diff(got, want, protocmp.Transform()); diff != "" {
		t.Errorf("service.CreateUser(ctx, %v) = %#v, _; want %v\ndiff = %s", req, got, want, diff)
	}

	con := r.NewConnection(ctx)
	gotUser, err := con.User().FindByID(uID)
	if err != nil {
		t.Fatalf("con.User().FindByID(%s) = _, %#v; want nil", uID, err)
	}
	if gotUser == nil {
		t.Errorf("con.User().FindByID(%s) = nil, _; want not nil", uID)
	}
}
