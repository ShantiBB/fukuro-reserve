package booking

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

	bookingv1 "booking/api/booking/v1"
	"booking/internal/config"
	"booking/internal/grpc/handler"
	"booking/internal/repository/postgres"
	"booking/internal/service"
)

type App struct {
	Config *config.Config
	Logger *slog.Logger
}

func (app *App) MustLoadGRPC() {
	slog.SetDefault(app.Logger)

	repo, err := postgres.New(app.Config)
	if err != nil {
		panic(err.Error())
	}

	validator, err := protovalidate.New()
	if err != nil {
		panic(err.Error())
	}

	svc := service.New(repo)
	h := handler.New(svc, validator)

	addr := fmt.Sprintf("%s:%d", app.Config.Server.Host, app.Config.Server.Port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err.Error())
	}

	grpcServer := newGRPCServer(app.Logger)

	bookingv1.RegisterBookingServiceServer(grpcServer, h)
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

	slog.Info("Shutting down gRPC server...")
	grpcServer.GracefulStop()
	slog.Info("gRPC server stopped")
}
