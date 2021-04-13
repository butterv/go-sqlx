package main

import (
	"fmt"
	"net"
	"os"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"

	"github.com/butterv/go-sqlx/app/domain/model"
	"github.com/butterv/go-sqlx/app/domain/repository"
	"github.com/butterv/go-sqlx/app/domain/service/health"
	"github.com/butterv/go-sqlx/app/domain/service/user"
	"github.com/butterv/go-sqlx/app/infrastructure/persistence"
	healthpb "github.com/butterv/go-sqlx/app/interface/rpc/v1/health"
	userpb "github.com/butterv/go-sqlx/app/interface/rpc/v1/user"
)

func main() {
	listenPort, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatalln(err)
	}

	if host := os.Getenv("MYSQL_HOST"); host != "" {
		fmt.Printf("MYSQL_HOST: %s\n", host)
	}
	if user := os.Getenv("MYSQL_USER"); user != "" {
		fmt.Printf("MYSQL_USER: %s\n", user)
	}
	if password := os.Getenv("MYSQL_PASSWORD"); password != "" {
		fmt.Printf("MYSQL_PASSWORD: %s\n", password)
	}

	db := connectDB("")
	repo := persistence.New(db)

	s := newGRPCServer(repo)
	reflection.Register(s)
	_ = s.Serve(listenPort)
	s.GracefulStop()
}

func connectDB(connection string) *sqlx.DB {
	db, err := sqlx.Open("mysql", connection)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func newGRPCServer(repo repository.Repository) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	healthpb.RegisterHealthServer(s, health.NewHealthService())
	userpb.RegisterUsersServer(s, user.NewUserService(repo, model.NewDefaultUserIDGenerator()))

	return s
}
