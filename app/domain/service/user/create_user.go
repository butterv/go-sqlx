package user

import (
	"context"

	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func (*userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	panic("not implemented yet")
}
