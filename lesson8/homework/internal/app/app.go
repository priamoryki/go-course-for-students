package app

import (
	"errors"
	"fmt"
	"github.com/priamoryki/validator"
	"homework8/internal/adapters/filters"
	"homework8/internal/ads"
	"time"
)

const (
	NonPublished = 1 << iota
	ByAuthor
	ByCreationTime
)

var ErrUserNotFound = errors.New("can't find user with such ID")
var ErrAdNotFound = errors.New("can't find ad with such ID")
var ErrNotUsersAd = errors.New("you don't have ad with such ID")
var ErrValidation = errors.New("validation error")

type App interface {
	CreateUser(nickname string, email string) (*ads.User, error)
	UpdateUser(userID int64, nickname string, email string) (*ads.User, error)
	FindUser(nickname string) (*ads.User, error)
	ListAds(bitmask int64) []*ads.Ad
	CreateAd(title string, text string, userId int64) (*ads.Ad, error)
	ChangeAdStatus(adID int64, userID int64, published bool) (*ads.Ad, error)
	GetAd(adID int64) (*ads.Ad, error)
	UpdateAd(adID int64, userID int64, title string, text string) (*ads.Ad, error)
	FindAd(title string) (*ads.Ad, error)
}

type Repository[T any] interface {
	GetAll(f filters.Filters[T]) []*T
	Add(ad *T) error
	FindByID(id int64) (*T, error)
	FindByName(name string) (*T, error)
}

type AdValidatorStruct struct {
	Title string `validate:"min:1;max:100"`
	Text  string `validate:"min:1;max:500"`
}

type Impl struct {
	adsRepository   Repository[ads.Ad]
	usersRepository Repository[ads.User]
}

func (a Impl) CreateUser(nickname string, email string) (*ads.User, error) {
	user := &ads.User{
		Nickname: nickname,
		Email:    email,
	}
	err := a.usersRepository.Add(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a Impl) UpdateUser(userID int64, nickname string, email string) (*ads.User, error) {
	user, err := a.usersRepository.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	user.Nickname = nickname
	user.Email = email
	return user, nil
}

func (a Impl) FindUser(nickname string) (*ads.User, error) {
	return a.usersRepository.FindByName(nickname)
}

func (a Impl) ListAds(bitmask int64) []*ads.Ad {
	f := make(filters.Filters[ads.Ad], 0)
	if !(bitmask&NonPublished != 0) {
		f = append(f, filters.NewFilterNonPublished())
	}
	if bitmask&ByAuthor != 0 {
		f = append(f, filters.NewFilterByAuthor())
	}
	if bitmask&ByCreationTime != 0 {
		f = append(f, filters.NewFilterByCreationTime())
	}
	return a.adsRepository.GetAll(f)
}

func (a Impl) CreateAd(title string, text string, userID int64) (*ads.Ad, error) {
	_, err := a.usersRepository.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err = validator.Validate(validateStruct)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	ad := &ads.Ad{
		Title:        title,
		Text:         text,
		AuthorID:     userID,
		Published:    false,
		CreationTime: time.Now().UTC(),
	}
	ad.LastUpdateTime = ad.CreationTime
	err = a.adsRepository.Add(ad)
	if err != nil {
		return nil, err
	}
	return ad, nil
}

func (a Impl) ChangeAdStatus(adID int64, userID int64, published bool) (*ads.Ad, error) {
	_, err := a.usersRepository.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	ad, err := a.adsRepository.FindByID(adID)
	if err != nil {
		return nil, ErrAdNotFound
	}

	if ad.AuthorID != userID {
		return nil, ErrNotUsersAd
	}

	ad.Published = published
	return ad, nil
}

func (a Impl) GetAd(adID int64) (*ads.Ad, error) {
	return a.adsRepository.FindByID(adID)
}

func (a Impl) UpdateAd(adID int64, userID int64, title string, text string) (*ads.Ad, error) {
	_, err := a.usersRepository.FindByID(userID)
	if err != nil {
		return nil, ErrUserNotFound
	}

	validateStruct := AdValidatorStruct{
		Title: title,
		Text:  text,
	}
	err = validator.Validate(validateStruct)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrValidation, err.Error())
	}

	ad, err := a.adsRepository.FindByID(adID)
	if err != nil {
		return nil, ErrAdNotFound
	}

	if ad.AuthorID != userID {
		return nil, ErrNotUsersAd
	}

	ad.LastUpdateTime = time.Now().UTC()
	ad.Title = title
	ad.Text = text
	return ad, nil
}

func (a Impl) FindAd(title string) (*ads.Ad, error) {
	return a.adsRepository.FindByName(title)
}

func NewApp(adsRepository Repository[ads.Ad], usersRepository Repository[ads.User]) App {
	return &Impl{
		adsRepository:   adsRepository,
		usersRepository: usersRepository,
	}
}
