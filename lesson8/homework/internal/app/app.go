package app

import (
	"errors"
	"fmt"
	"github.com/priamoryki/validator"
	"homework8/internal/ads"
)

var ErrAdNotFound = errors.New("can't find ad with such ID")
var ErrNotUsersAd = errors.New("you don't have ad with such ID")
var ErrValidation = errors.New("validation error")

type App interface {
	CreateAd(title string, text string, userId int64) (*ads.Ad, error)
	ChangeAdStatus(adID int64, userID int64, published bool) (*ads.Ad, error)
	UpdateAd(adID int64, userID int64, title string, text string) (*ads.Ad, error)
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

func (a Impl) CreateAd(title string, text string, userId int64) (*ads.Ad, error) {
	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err := validator.Validate(validateStruct)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
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
		return nil, err
	}
	return ad, nil
}

func (a Impl) ChangeAdStatus(adID int64, userID int64, published bool) (*ads.Ad, error) {
	ad, err := a.repository.FindByID(adID)
	if err != nil {
		return nil, ErrAdNotFound
	}

	if ad.AuthorID != userID {
		return nil, ErrNotUsersAd
	}

	ad.Published = published
	return ad, nil
}

func (a Impl) UpdateAd(adID int64, userID int64, title string, text string) (*ads.Ad, error) {
	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err := validator.Validate(validateStruct)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	ad, err := a.repository.FindByID(adID)
	if err != nil {
		return nil, ErrAdNotFound
	}

	if ad.AuthorID != userID {
		return nil, ErrNotUsersAd
	}

	ad.Title = title
	ad.Text = text
	return ad, nil
}

func NewApp(repo Repository) App {
	return &Impl{repository: repo}
}
