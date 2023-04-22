package grpc

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	"homework9/internal/ads"
	"homework9/internal/app"
)

type Server struct {
	UnimplementedAdServiceServer
	a app.App
}

func NewService(a app.App) *Server {
	return &Server{
		a: a,
	}
}

func adToAdResponse(ad *ads.Ad) *AdResponse {
	return &AdResponse{
		Id:        ad.ID,
		Title:     ad.Title,
		Text:      ad.Text,
		AuthorId:  ad.AuthorID,
		Published: ad.Published,
	}
}

func userToUserResponse(user *ads.User) *UserResponse {
	return &UserResponse{
		Id:   user.ID,
		Name: user.Nickname,
	}
}

func (s *Server) CreateAd(_ context.Context, req *CreateAdRequest) (*AdResponse, error) {
	ad, err := s.a.CreateAd(req.Title, req.Text, req.UserId)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) ChangeAdStatus(_ context.Context, req *ChangeAdStatusRequest) (*AdResponse, error) {
	ad, err := s.a.ChangeAdStatus(req.AdId, req.UserId, req.Published)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) UpdateAd(_ context.Context, req *UpdateAdRequest) (*AdResponse, error) {
	ad, err := s.a.UpdateAd(req.AdId, req.UserId, req.Title, req.Text)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) ListAds(_ context.Context, _ *emptypb.Empty) (*ListAdResponse, error) {
	result := make([]*AdResponse, 0)
	for _, ad := range s.a.ListAds(0) {
		result = append(result, adToAdResponse(ad))
	}
	return &ListAdResponse{List: result}, nil
}

func (s *Server) CreateUser(_ context.Context, req *CreateUserRequest) (*UserResponse, error) {
	user, err := s.a.CreateUser(req.Name, "")
	if err != nil {
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *Server) GetUser(_ context.Context, req *GetUserRequest) (*UserResponse, error) {
	user, err := s.a.GetUser(req.Id)
	if err != nil {
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *Server) DeleteUser(_ context.Context, req *DeleteUserRequest) (*emptypb.Empty, error) {
	_, err := s.a.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}

func (s *Server) DeleteAd(_ context.Context, req *DeleteAdRequest) (*emptypb.Empty, error) {
	_, err := s.a.DeleteAd(req.AdId, req.AuthorId)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
