package main

import (
	"context"
	"errors"
	"fmt"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/test/bufconn"
	"homework9/internal/adapters/adrepo"
	"homework9/internal/adapters/userrepo"
	"homework9/internal/app"
	grpcPort "homework9/internal/ports/grpc"
	"homework9/internal/ports/httpgin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	logger := log.Default()

	a := app.NewApp(adrepo.New(), userrepo.New())

	sigQuit := make(chan os.Signal, 1)
	signal.Ignore(syscall.SIGHUP, syscall.SIGPIPE)
	signal.Notify(sigQuit, syscall.SIGINT, syscall.SIGTERM)

	eg, ctx := errgroup.WithContext(context.Background())
	eg.Go(func() error {
		select {
		case s := <-sigQuit:
			return fmt.Errorf("signal: %v", s)
		case <-ctx.Done():
			return nil
		}
	})

	grpcServer := grpcPort.NewGRPCServer(logger, a)

	// start GRPC server
	eg.Go(func() error {
		logger.Println("starting GRPC server")
		errCh := make(chan error)

		defer func() {
			logger.Println("stopping GRPC server")
			grpcServer.GracefulStop()
			close(errCh)
		}()

		go func() {
			if err := grpcServer.Serve(bufconn.Listen(1024 * 1024)); err != nil {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("GRPC server error: %w", err)
		}
	})

	httpServer := httpgin.NewHTTPServer(":18080", a)

	// start HTTP server
	eg.Go(func() error {
		logger.Println("starting HTTP server")
		errCh := make(chan error)

		defer func() {
			logger.Println("stopping HTTP server")
			shutdownCtx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
			defer cancel()
			if err := httpServer.Shutdown(shutdownCtx); err != nil {
				logger.Printf("error on HTTP server closing occurred: %s", err.Error())
			}
			close(errCh)
		}()

		go func() {
			if err := httpServer.Listen(); !errors.Is(err, http.ErrServerClosed) {
				errCh <- err
			}
		}()

		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-errCh:
			return fmt.Errorf("HTTP server error: %w", err)
		}
	})

	if err := eg.Wait(); err != nil {
		logger.Printf("servers shutdown: %s\n", err.Error())
	}
}
