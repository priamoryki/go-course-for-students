package grpc

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"homework10/internal/app"
	"log"
)

func NewGRPCServer(logger *log.Logger, a app.App) *grpc.Server {
	// logger и panic interceptor
	server := grpc.NewServer(
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
	RegisterAdServiceServer(server, NewService(a))
	return server
}
