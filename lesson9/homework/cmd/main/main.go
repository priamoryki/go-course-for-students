package main

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"homework9/internal/adapters/adrepo"
	"homework9/internal/adapters/userrepo"
	"homework9/internal/app"
	grpcPort "homework9/internal/ports/grpc"
	"homework9/internal/ports/httpgin"
	"log"
)

func createServer(ctx context.Context, logger *log.Logger) *grpc.Server {
	// logger и panic interceptor
	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			logging.UnaryServerInterceptor(
				logging.LoggerFunc(func(_ context.Context, _ logging.Level, msg string, fields ...any) {
					logger.Printf("message: %s, fields: %v\n", msg, fields)
				}),
			),
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(func(p interface{}) (err error) {
				logger.Printf("panic: %v\n", p)
				return
			})),
		),
	)
	// graceful shutdown
	go func(ctx context.Context) {
		<-ctx.Done()
		srv.GracefulStop()
	}(ctx)
	return srv
}

func main() {
	ctx := context.Background()

	logger := log.Default()

	a := app.NewApp(adrepo.New(), userrepo.New())

	grpcServer := createServer(ctx, logger)
	grpcPort.RegisterAdServiceServer(grpcServer, grpcPort.NewService(a))
	err := grpcServer.Serve(bufconn.Listen(1024 * 1024))
	if err != nil {
		fmt.Printf("error occurred on GRPC server start: %s", err.Error())
		return
	}

	httpServer := httpgin.NewHTTPServer(":18080", a)
	err = httpServer.Listen()
	if err != nil {
		fmt.Printf("error occurred on HTTP server start: %s", err.Error())
		return
	}
}
