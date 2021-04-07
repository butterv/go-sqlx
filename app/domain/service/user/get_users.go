package user

import (
	"context"

	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func (*userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	panic("not implemented yet")
}
