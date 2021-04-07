package user

import (
	"context"

	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func (*userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	panic("not implemented yet")
}
