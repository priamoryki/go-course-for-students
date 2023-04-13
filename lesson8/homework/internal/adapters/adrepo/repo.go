package adrepo

import (
	"errors"
	"homework8/internal/ads"
	"homework8/internal/app"
)

var ErrAdNotFound = errors.New("ad not found")

type Impl struct {
	currentId int64
	idToAd    map[int64]*ads.Ad
}

func (i *Impl) GetAll() []*ads.Ad {
	result := make([]*ads.Ad, 0)
	for j := int64(0); j < i.GetNextID(); j++ {
		ad := i.idToAd[j]
		if ad.Published {
			result = append(result, ad)
		}
	}
	return result
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

func (i *Impl) FindByName(name string) (*ads.Ad, error) {
	for _, ad := range i.idToAd {
		if ad.Title == name {
			return ad, nil
		}
	}
	return nil, ErrAdNotFound
}

func New() app.Repository[ads.Ad] {
	return &Impl{
		currentId: 0,
		idToAd:    make(map[int64]*ads.Ad),
	}
}
