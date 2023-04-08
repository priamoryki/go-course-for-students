package app

import (
	"errors"
	"github.com/valyala/fasthttp"
	"homework6/internal/ads"
)

type App interface {
	CreateAd(ctx *fasthttp.RequestCtx, title string, text string, userId int64) (*ads.Ad, error)
	ChangeAdStatus(ctx *fasthttp.RequestCtx, adID int64, userID int64, published bool) (*ads.Ad, error)
	UpdateAd(ctx *fasthttp.RequestCtx, adID int64, userID int64, title string, text string) (*ads.Ad, error)
}

type Repository interface {
	Add(ad *ads.Ad) error
	GetNextID() int64
	FindByID(adID int64) (*ads.Ad, error)
}

type Impl struct {
	repository Repository
}

func (a Impl) CreateAd(_ *fasthttp.RequestCtx, title string, text string, userId int64) (*ads.Ad, error) {
	ad := &ads.Ad{
		ID:        a.repository.GetNextID(),
		Title:     title,
		Text:      text,
		AuthorID:  userId,
		Published: false,
	}
	err := a.repository.Add(ad)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (a Impl) ChangeAdStatus(_ *fasthttp.RequestCtx, adID int64, userID int64, published bool) (*ads.Ad, error) {
	ad, err := a.repository.FindByID(adID)
	if err != nil {
		return nil, err
	}
	if ad.AuthorID != userID {
		return nil, errors.New("you don't have ad with such ID")
	}
	ad.Published = published
	return ad, nil
}

func (a Impl) UpdateAd(_ *fasthttp.RequestCtx, adID int64, userID int64, title string, text string) (*ads.Ad, error) {
	ad, err := a.repository.FindByID(adID)
	if err != nil {
		return nil, err
	}
	if ad.AuthorID != userID {
		return nil, errors.New("you don't have ad with such ID")
	}
	ad.Title = title
	ad.Text = text
	return ad, nil
}

func NewApp(repo Repository) App {
	return &Impl{repository: repo}
}
