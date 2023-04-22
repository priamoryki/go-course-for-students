package grpc

import (
	"context"
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
		Id:    user.ID,
		Name:  user.Nickname,
		Email: user.Email,
	}
}

func (s *Server) CreateUser(_ context.Context, req *CreateUserRequest) (*UserResponse, error) {
	user, err := s.a.CreateUser(req.Name, req.Email)
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

func (s *Server) UpdateUser(_ context.Context, req *UpdateUserRequest) (*UserResponse, error) {
	user, err := s.a.UpdateUser(req.Id, req.Name, req.Email)
	if err != nil {
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *Server) FindUser(_ context.Context, req *FindUserRequest) (*UserResponse, error) {
	user, err := s.a.FindUser(req.Query)
	if err != nil {
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *Server) DeleteUser(_ context.Context, req *DeleteUserRequest) (*UserResponse, error) {
	user, err := s.a.DeleteUser(req.Id)
	if err != nil {
		return nil, err
	}
	return userToUserResponse(user), nil
}

func (s *Server) ListAds(_ context.Context, req *ListAdsRequest) (*ListAdResponse, error) {
	result := make([]*AdResponse, 0)
	for _, ad := range s.a.ListAds(req.Bitmask) {
		result = append(result, adToAdResponse(ad))
	}
	return &ListAdResponse{List: result}, nil
}

func (s *Server) CreateAd(_ context.Context, req *CreateAdRequest) (*AdResponse, error) {
	ad, err := s.a.CreateAd(req.Title, req.Text, req.UserId)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) GetAd(_ context.Context, req *GetAdRequest) (*AdResponse, error) {
	ad, err := s.a.GetAd(req.Id)
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

func (s *Server) ChangeAdStatus(_ context.Context, req *ChangeAdStatusRequest) (*AdResponse, error) {
	ad, err := s.a.ChangeAdStatus(req.AdId, req.UserId, req.Published)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) FindAd(_ context.Context, req *FindAdRequest) (*AdResponse, error) {
	ad, err := s.a.FindAd(req.Query)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}

func (s *Server) DeleteAd(_ context.Context, req *DeleteAdRequest) (*AdResponse, error) {
	ad, err := s.a.DeleteAd(req.AdId, req.AuthorId)
	if err != nil {
		return nil, err
	}
	return adToAdResponse(ad), nil
}
