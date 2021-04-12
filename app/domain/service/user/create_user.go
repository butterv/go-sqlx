package user

import (
	"context"

	"github.com/butterv/go-sqlx/app/domain/repository"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
	appstatus "github.com/butterv/go-sqlx/app/status"
)

func (s *userService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	con := s.r.NewConnection(ctx)
	defer con.Close()

	uID := s.userIDGenerator.Generate()
	err := con.RunTransaction(func(tx repository.Transaction) error {
		err := tx.User().Create(uID, req.GetEmail())
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToCreateUser.Err()
	}

	return &pb.CreateUserResponse{
		UserId: string(uID),
	}, nil
}
