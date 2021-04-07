package user

import (
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

type userService struct {
	pb.UnimplementedUsersServer // embedding
}

// NewUserService generates the `UsersServer` implementation.
func NewUserService() pb.UsersServer {
	return &userService{}
}
