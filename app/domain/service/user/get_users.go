package user

import (
	"context"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/presenter"
	appstatus "github.com/butterv/go-sqlx/app/status"

	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func (s *userService) GetUsers(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	con := s.r.NewConnection(ctx)
	defer con.Close()

	var uIDs []model.UserID
	for _, uID := range req.GetUserIds() {
		uIDs = append(uIDs, model.UserID(uID))
	}

	us, err := con.User().FindByIDs(uIDs)
	if err != nil {
		// TODO(butterv): output error log
		return nil, appstatus.FailedToGetUsers.Err()
	}

	return &pb.GetUsersResponse{
		Users: presenter.UsersToPbUsers(us),
	}, nil
}
