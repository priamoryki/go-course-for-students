package tests

import (
	"context"
	"google.golang.org/grpc/credentials/insecure"
	"homework9/internal/adapters/userrepo"
	"log"
	"net"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"homework9/internal/adapters/adrepo"
	"homework9/internal/app"
	grpcPort "homework9/internal/ports/grpc"
)

func getGrpcTestClient(t *testing.T) (context.Context, grpcPort.AdServiceClient) {
	logger := log.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	t.Cleanup(func() {
		cancel()
	})

	lis := bufconn.Listen(1024 * 1024)
	t.Cleanup(func() {
		lis.Close()
	})

	srv := grpcPort.NewGRPCServer(logger, app.NewApp(adrepo.New(), userrepo.New()))
	t.Cleanup(func() {
		srv.Stop()
	})

	go func() {
		assert.NoError(t, srv.Serve(lis), "srv.Serve")
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	assert.NoError(t, err, "grpc.DialContext")

	t.Cleanup(func() {
		conn.Close()
	})

	return ctx, grpcPort.NewAdServiceClient(conn)
}

func TestGRPCCreateUser(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	res, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "Oleg", res.Name)
	assert.Equal(t, "test@gmail.com", res.Email)
}

func TestGRPCGetUser(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	res, err := client.GetUser(ctx, &grpcPort.GetUserRequest{Id: 0})
	assert.NoError(t, err, "client.GetUser")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "Oleg", res.Name)
	assert.Equal(t, "test@gmail.com", res.Email)
}

func TestGRPCUpdateUser(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	res, err := client.UpdateUser(ctx, &grpcPort.UpdateUserRequest{Id: 0, Name: "Oleg1", Email: "test1@gmail.com"})
	assert.NoError(t, err, "client.UpdateUser")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "Oleg1", res.Name)
	assert.Equal(t, "test1@gmail.com", res.Email)
}

func TestGRPCFindUser(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	res, err := client.FindUser(ctx, &grpcPort.FindUserRequest{Query: "Oleg"})
	assert.NoError(t, err, "client.FindUser")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "Oleg", res.Name)
	assert.Equal(t, "test@gmail.com", res.Email)
}

func TestGRPCDeleteUser(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	res, err := client.DeleteUser(ctx, &grpcPort.DeleteUserRequest{Id: 0})
	assert.NoError(t, err, "client.DeleteUser")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "Oleg", res.Name)
	assert.Equal(t, "test@gmail.com", res.Email)

	_, err = client.GetUser(ctx, &grpcPort.GetUserRequest{Id: 0})
	assert.Error(t, err, "client.GetUser")
}

func TestGRPCListAds(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.ListAds(ctx, &grpcPort.ListAdsRequest{Bitmask: 0})
	assert.NoError(t, err, "client.ListAds")
	assert.Equal(t, 0, len(res.List))

	_, err = client.ChangeAdStatus(ctx, &grpcPort.ChangeAdStatusRequest{AdId: 0, UserId: 0, Published: true})
	assert.NoError(t, err, "client.ChangeAdStatus")

	res, err = client.ListAds(ctx, &grpcPort.ListAdsRequest{Bitmask: 0})
	assert.NoError(t, err, "client.ListAds")
	assert.Equal(t, 1, len(res.List))
	assert.Equal(t, int64(0), res.List[0].Id)
}

func TestGRPCCreateAd(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	res, err := client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title", res.Title)
	assert.Equal(t, "text", res.Text)
	assert.Equal(t, false, res.Published)
}

func TestGRPCGetAd(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.GetAd(ctx, &grpcPort.GetAdRequest{Id: 0})
	assert.NoError(t, err, "client.GetAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title", res.Title)
	assert.Equal(t, "text", res.Text)
	assert.Equal(t, false, res.Published)
}

func TestGRPCUpdateAd(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.UpdateAd(ctx, &grpcPort.UpdateAdRequest{Title: "title1", Text: "text1", UserId: 0})
	assert.NoError(t, err, "client.UpdateAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title1", res.Title)
	assert.Equal(t, "text1", res.Text)
	assert.Equal(t, false, res.Published)
}

func TestGRPCChangeAdStatus(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.ChangeAdStatus(ctx, &grpcPort.ChangeAdStatusRequest{AdId: 0, UserId: 0, Published: true})
	assert.NoError(t, err, "client.UpdateAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title", res.Title)
	assert.Equal(t, "text", res.Text)
	assert.Equal(t, true, res.Published)
}

func TestGRPCFindAd(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.FindAd(ctx, &grpcPort.FindAdRequest{Query: "title"})
	assert.NoError(t, err, "client.FindAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title", res.Title)
	assert.Equal(t, "text", res.Text)
	assert.Equal(t, false, res.Published)
}

func TestGRPCDeleteAd(t *testing.T) {
	ctx, client := getGrpcTestClient(t)

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	assert.NoError(t, err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	assert.NoError(t, err, "client.CreateAd")

	res, err := client.DeleteAd(ctx, &grpcPort.DeleteAdRequest{AdId: 0, AuthorId: 0})
	assert.NoError(t, err, "client.DeleteAd")
	assert.Equal(t, int64(0), res.Id)
	assert.Equal(t, "title", res.Title)
	assert.Equal(t, "text", res.Text)
	assert.Equal(t, false, res.Published)

	_, err = client.GetAd(ctx, &grpcPort.GetAdRequest{Id: 0})
	assert.Error(t, err, "client.GetAd")
}
