package adrepo

import (
	"errors"
	"homework6/internal/ads"
	"homework6/internal/app"
)

type Impl struct {
	currentId int64
	idToAd    map[int64]*ads.Ad
}

func (i *Impl) Add(ad *ads.Ad) error {
	i.idToAd[ad.ID] = ad
	i.currentId += 1
	return nil
}

func (i *Impl) GetNextID() int64 {
	return i.currentId
}

func (i *Impl) FindByID(adID int64) (*ads.Ad, error) {
	ad, ok := i.idToAd[adID]
	if !ok {
		return nil, errors.New("there is no ad with such ID")
	}
	return ad, nil
}

func New() app.Repository {
	return &Impl{
		currentId: 0,
		idToAd:    make(map[int64]*ads.Ad),
	}
}
