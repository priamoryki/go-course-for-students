package userrepo

import (
	"errors"
	"homework8/internal/adapters/filters"
	"homework8/internal/ads"
	"homework8/internal/app"
	"sync"
)

var ErrUserNotFound = errors.New("user not found")

type Impl struct {
	currentId int64
	idToUser  map[int64]*ads.User
	mutex     *sync.RWMutex
}

func (i *Impl) GetAll(f filters.Filters[ads.User]) []*ads.User {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	result := make([]*ads.User, 0)
	for j := int64(0); j < i.currentId; j++ {
		result = append(result, i.idToUser[j])
	}
	return f.Filter(result)
}

func (i *Impl) Add(user *ads.User) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	user.SetID(i.currentId)
	i.idToUser[user.GetID()] = user
	i.currentId += 1
	return nil
}

func (i *Impl) FindByID(userID int64) (*ads.User, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	ad, ok := i.idToUser[userID]
	if !ok {
		return nil, errors.New("there is no user with such ID")
	}
	return ad, nil
}

func (i *Impl) FindByName(name string) (*ads.User, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	for _, user := range i.idToUser {
		if user.HasName(name) {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func New() app.Repository[ads.User] {
	return &Impl{
		currentId: 0,
		idToUser:  make(map[int64]*ads.User),
		mutex:     new(sync.RWMutex),
	}
}
