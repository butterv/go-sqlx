package status

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errorpb "github.com/butterv/go-sqlx/app/interface/rpc/v1/error"
)

// LocaleJapan is a Japanese local id in Japan.
const LocaleJapan = "ja-JP"

// LocaleEnUS is a locale id for English in U.S..
const LocaleEnUS = "en-US"

var (
	FailedToCreateUser = must(status.New(codes.Internal, "failed to create user").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "FAILED_TO_CREATE_USER",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJapan,
			Message: "ユーザー作成に失敗しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUS,
			Message: "Failed to create user.",
		},
	))

	FailedToGetUser = must(status.New(codes.Internal, "failed to get user").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "FAILED_TO_GET_USER",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJapan,
			Message: "ユーザーの取得に失敗しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUS,
			Message: "Failed to get user.",
		},
	))

	FailedToGetUsers = must(status.New(codes.Internal, "failed to get users").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "FAILED_TO_GET_USERS",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJapan,
			Message: "ユーザーの取得に失敗しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUS,
			Message: "Failed to get users.",
		},
	))
)

func must(s *status.Status, err error) *status.Status {
	if err != nil {
		panic(err)
	}
	return s
}
