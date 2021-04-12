package user

import (
	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/domain/repository"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

type userService struct {
	r repository.Repository

	userIDGenerator model.UserIDGenerator

	pb.UnimplementedUsersServer // embedding
}

// NewUserService generates the `UsersServer` implementation.
func NewUserService(r repository.Repository, userIDGenerator model.UserIDGenerator) pb.UsersServer {
	return &userService{
		r:               r,
		userIDGenerator: userIDGenerator,
	}
}
