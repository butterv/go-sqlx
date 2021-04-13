package user

import (
	"context"

	"github.com/butterv/go-sqlx/app/domain/model"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
	"github.com/butterv/go-sqlx/app/presenter"
	appstatus "github.com/butterv/go-sqlx/app/status"
)

func (s *userService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	con := s.r.NewConnection(ctx)
	defer con.Close()

	uID := model.UserID(req.GetUserId())
	u, err := con.User().FindByID(uID)
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToGetUser.Err()
	}

	return &pb.GetUserResponse{
		User: presenter.UserToPbUser(u),
	}, nil
}
