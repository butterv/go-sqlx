package presenter

import (
	"github.com/butterv/go-sqlx/app/domain/model"
	pb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func UserToPbUser(u *model.User) *pb.User {
	if u == nil {
		return nil
	}

	return &pb.User{
		UserId: string(u.ID),
		Email:  u.Email,
	}
}

func UsersToPbUsers(us []*model.User) []*pb.User {
	if len(us) == 0 {
		return nil
	}

	var out []*pb.User
	for _, u := range us {
		out = append(out, UserToPbUser(u))
	}

	return out
}
