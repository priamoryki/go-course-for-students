package tests

import (
	"context"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/credentials/insecure"
	"homework10/internal/adapters/userrepo"
	"log"
	"net"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/test/bufconn"
	"homework10/internal/adapters/adrepo"
	"homework10/internal/app"
	grpcPort "homework10/internal/ports/grpc"
)

type GRPCSuite struct {
	suite.Suite
	Ctx    context.Context
	Cancel context.CancelFunc
	Lis    *bufconn.Listener
	Srv    *grpc.Server
	Conn   *grpc.ClientConn
	Client grpcPort.AdServiceClient
}

func (s *GRPCSuite) TearDownTests() {
	s.Cancel()
	s.Lis.Close()
	s.Srv.Stop()
	s.Conn.Close()
}

func (s *GRPCSuite) SetupTest() {
	logger := log.Default()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	s.Ctx = ctx
	s.Cancel = cancel

	lis := bufconn.Listen(1024 * 1024)
	s.Lis = lis

	srv := grpcPort.NewGRPCServer(logger, app.NewApp(adrepo.New(), userrepo.New()))
	s.Srv = srv

	go func() {
		s.NoError(srv.Serve(lis), "srv.Serve")
	}()

	dialer := func(context.Context, string) (net.Conn, error) {
		return lis.Dial()
	}

	conn, err := grpc.DialContext(ctx, "", grpc.WithContextDialer(dialer), grpc.WithTransportCredentials(insecure.NewCredentials()))
	s.NoError(err, "grpc.DialContext")
	s.Conn = conn

	s.Client = grpcPort.NewAdServiceClient(conn)
}

func (s *GRPCSuite) TestGRPCCreateUser() {
	ctx, client := s.Ctx, s.Client

	res, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")
	s.Equal(int64(0), res.Id)
	s.Equal("Oleg", res.Name)
	s.Equal("test@gmail.com", res.Email)
}

func (s *GRPCSuite) TestGRPCGetUser() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	res, err := client.GetUser(ctx, &grpcPort.GetUserRequest{Id: 0})
	s.NoError(err, "client.GetUser")
	s.Equal(int64(0), res.Id)
	s.Equal("Oleg", res.Name)
	s.Equal("test@gmail.com", res.Email)
}

func (s *GRPCSuite) TestGRPCUpdateUser() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	res, err := client.UpdateUser(ctx, &grpcPort.UpdateUserRequest{Id: 0, Name: "Oleg1", Email: "test1@gmail.com"})
	s.NoError(err, "client.UpdateUser")
	s.Equal(int64(0), res.Id)
	s.Equal("Oleg1", res.Name)
	s.Equal("test1@gmail.com", res.Email)
}

func (s *GRPCSuite) TestGRPCFindUser() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	res, err := client.FindUser(ctx, &grpcPort.FindUserRequest{Query: "Oleg"})
	s.NoError(err, "client.FindUser")
	s.Equal(int64(0), res.Id)
	s.Equal("Oleg", res.Name)
	s.Equal("test@gmail.com", res.Email)
}

func (s *GRPCSuite) TestGRPCDeleteUser() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	res, err := client.DeleteUser(ctx, &grpcPort.DeleteUserRequest{Id: 0})
	s.NoError(err, "client.DeleteUser")
	s.Equal(int64(0), res.Id)
	s.Equal("Oleg", res.Name)
	s.Equal("test@gmail.com", res.Email)

	_, err = client.GetUser(ctx, &grpcPort.GetUserRequest{Id: 0})
	s.Error(err, "client.GetUser")
}

func (s *GRPCSuite) TestGRPCListAds() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.ListAds(ctx, &grpcPort.ListAdsRequest{Bitmask: 0})
	s.NoError(err, "client.ListAds")
	s.Equal(0, len(res.List))

	_, err = client.ChangeAdStatus(ctx, &grpcPort.ChangeAdStatusRequest{AdId: 0, UserId: 0, Published: true})
	s.NoError(err, "client.ChangeAdStatus")

	res, err = client.ListAds(ctx, &grpcPort.ListAdsRequest{Bitmask: 0})
	s.NoError(err, "client.ListAds")
	s.Equal(1, len(res.List))
	s.Equal(int64(0), res.List[0].Id)
}

func (s *GRPCSuite) TestGRPCCreateAd() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	res, err := client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *GRPCSuite) TestGRPCGetAd() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.GetAd(ctx, &grpcPort.GetAdRequest{Id: 0})
	s.NoError(err, "client.GetAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *GRPCSuite) TestGRPCUpdateAd() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.UpdateAd(ctx, &grpcPort.UpdateAdRequest{Title: "title1", Text: "text1", UserId: 0})
	s.NoError(err, "client.UpdateAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title1", res.Title)
	s.Equal("text1", res.Text)
	s.Equal(false, res.Published)
}

func (s *GRPCSuite) TestGRPCChangeAdStatus() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.ChangeAdStatus(ctx, &grpcPort.ChangeAdStatusRequest{AdId: 0, UserId: 0, Published: true})
	s.NoError(err, "client.UpdateAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(true, res.Published)
}

func (s *GRPCSuite) TestGRPCFindAd() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.FindAd(ctx, &grpcPort.FindAdRequest{Query: "title"})
	s.NoError(err, "client.FindAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)
}

func (s *GRPCSuite) TestGRPCDeleteAd() {
	ctx, client := s.Ctx, s.Client

	_, err := client.CreateUser(ctx, &grpcPort.CreateUserRequest{Name: "Oleg", Email: "test@gmail.com"})
	s.NoError(err, "client.CreateUser")

	_, err = client.CreateAd(ctx, &grpcPort.CreateAdRequest{Title: "title", Text: "text", UserId: 0})
	s.NoError(err, "client.CreateAd")

	res, err := client.DeleteAd(ctx, &grpcPort.DeleteAdRequest{AdId: 0, AuthorId: 0})
	s.NoError(err, "client.DeleteAd")
	s.Equal(int64(0), res.Id)
	s.Equal("title", res.Title)
	s.Equal("text", res.Text)
	s.Equal(false, res.Published)

	_, err = client.GetAd(ctx, &grpcPort.GetAdRequest{Id: 0})
	s.Error(err, "client.GetAd")
}

func TestGRPCSuite(t *testing.T) {
	suite.Run(t, new(GRPCSuite))
}
