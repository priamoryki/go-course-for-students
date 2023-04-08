package app

import (
	"errors"
	"github.com/priamoryki/validator"
	"github.com/valyala/fasthttp"
	"homework6/internal/ads"
	"net/http"
)

var AdNotFound = errors.New("can't find ad with such ID")
var NotUsersAd = errors.New("you don't have ad with such ID")

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

type AdValidatorStruct struct {
	Title string `validate:"min:1;max:100"`
	Text  string `validate:"min:1;max:500"`
}

type Impl struct {
	repository Repository
}

func (a Impl) CreateAd(c *fasthttp.RequestCtx, title string, text string, userId int64) (*ads.Ad, error) {
	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err := validator.Validate(validateStruct)
	if err != nil {
		c.Error(err.Error(), http.StatusBadRequest)
		return nil, err
	}

	ad := &ads.Ad{
		ID:        a.repository.GetNextID(),
		Title:     title,
		Text:      text,
		AuthorID:  userId,
		Published: false,
	}
	err = a.repository.Add(ad)
	if err != nil {
		c.Error(err.Error(), http.StatusInternalServerError)
		return nil, err
	}
	return ad, nil
}

func (a Impl) ChangeAdStatus(c *fasthttp.RequestCtx, adID int64, userID int64, published bool) (*ads.Ad, error) {
	ad, err := a.repository.FindByID(adID)
	if err != nil {
		c.Error(AdNotFound.Error(), http.StatusInternalServerError)
		return nil, AdNotFound
	}

	if ad.AuthorID != userID {
		c.Error(NotUsersAd.Error(), http.StatusForbidden)
		return nil, NotUsersAd
	}

	ad.Published = published
	return ad, nil
}

func (a Impl) UpdateAd(c *fasthttp.RequestCtx, adID int64, userID int64, title string, text string) (*ads.Ad, error) {
	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err := validator.Validate(validateStruct)
	if err != nil {
		c.Error(err.Error(), http.StatusBadRequest)
		return nil, err
	}

	ad, err := a.repository.FindByID(adID)
	if err != nil {
		c.Error(AdNotFound.Error(), http.StatusInternalServerError)
		return nil, AdNotFound
	}

	if ad.AuthorID != userID {
		c.Error(NotUsersAd.Error(), http.StatusForbidden)
		return nil, NotUsersAd
	}

	ad.Title = title
	ad.Text = text
	return ad, nil
}

func NewApp(repo Repository) App {
	return &Impl{repository: repo}
}
