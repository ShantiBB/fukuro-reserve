package auth

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"os/signal"
	"syscall"

	"buf.build/go/protovalidate"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	userv1 "auth/api/user/v1"
	"auth/internal/config"
	"auth/internal/grpc/handler"
	"auth/internal/repository/postgres"
	"auth/internal/service"
	"auth/pkg/lib/utils/jwt"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
}

func (app *App) MustLoadGRPC() {
	slog.SetDefault(app.Logger)

	repo := postgres.New(app.Config)

	tokenCredentials := jwt.GetTokenCredentials(app.Config)
	svc := service.New(repo, tokenCredentials)

	validator, err := protovalidate.New()
	if err != nil {
		panic(err.Error())
	}
	h := handler.New(svc, validator)

	addr := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err.Error())
	}

	grpcServer := newGRPCServer(app.Logger, app.Config.JWT.AccessSecret)

	userv1.RegisterUserServiceServer(grpcServer, h)
	userv1.RegisterTokenServiceServer(grpcServer, h)
	reflection.Register(grpcServer)

	go func() {
		slog.Info("Starting gRPC server", "address", addr)
		if err = grpcServer.Serve(lis); err != nil {
			slog.Error("Failed to serve", "error", err)
		}
	}()

	app.gracefulShutdown(grpcServer)
}

func (app *App) gracefulShutdown(grpcServer *grpc.Server) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	slog.Info("Shutting down gRPC server")
	grpcServer.GracefulStop()
	slog.Info("gRPC server stopped")
}
