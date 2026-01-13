package booking

import (
	"fmt"
	"log/slog"
	"net"

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
}

func (app *App) MustLoadGRPC() {
	repo, err := postgres.NewRepository(app.Config)
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

	grpcServer := grpc.NewServer()

	bookingv1.RegisterBookingServiceServer(grpcServer, h)
	reflection.Register(grpcServer)

	slog.Info("Starting gRPC server", "address", addr)
	if err = grpcServer.Serve(lis); err != nil {
		panic(err.Error())
	}
}
