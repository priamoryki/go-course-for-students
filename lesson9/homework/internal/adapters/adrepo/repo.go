package adrepo

import (
	"errors"
	"homework9/internal/adapters/filters"
	"homework9/internal/ads"
	"homework9/internal/app"
	"sync"
)

var ErrAdNotFound = errors.New("ad not found")

type Impl struct {
	currentId int64
	idToAd    map[int64]*ads.Ad
	mutex     *sync.RWMutex
}

func (i *Impl) GetAll(f filters.Filters[ads.Ad]) []*ads.Ad {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	result := make([]*ads.Ad, 0)
	for j := int64(0); j < i.currentId; j++ {
		result = append(result, i.idToAd[j])
	}
	return f.Filter(result)
}

func (i *Impl) Add(ad *ads.Ad) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	ad.SetID(i.currentId)
	i.idToAd[ad.GetID()] = ad
	i.currentId += 1
	return nil
}

func (i *Impl) FindByID(adID int64) (*ads.Ad, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	ad, ok := i.idToAd[adID]
	if !ok {
		return nil, ErrAdNotFound
	}
	return ad, nil
}

func (i *Impl) FindByName(name string) (*ads.Ad, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	for _, ad := range i.idToAd {
		if ad.HasName(name) {
			return ad, nil
		}
	}
	return nil, ErrAdNotFound
}

func (i *Impl) DeleteById(adID int64) (*ads.Ad, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	ad, ok := i.idToAd[adID]
	if !ok {
		return nil, ErrAdNotFound
	}
	delete(i.idToAd, adID)
	return ad, nil
}

func New() app.Repository[ads.Ad] {
	return &Impl{
		currentId: 0,
		idToAd:    make(map[int64]*ads.Ad),
		mutex:     new(sync.RWMutex),
	}
}
