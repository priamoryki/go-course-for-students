package userrepo

import (
	"errors"
	"homework8/internal/ads"
	"homework8/internal/app"
)

var ErrUserNotFound = errors.New("user not found")

type Impl struct {
	currentId int64
	idToUser  map[int64]*ads.User
}

func (i *Impl) GetAll() []*ads.User {
	result := make([]*ads.User, 0)
	for j := int64(0); j < i.GetNextID(); j++ {
		result = append(result, i.idToUser[j])
	}
	return result
}

func (i *Impl) Add(user *ads.User) error {
	i.idToUser[user.ID] = user
	i.currentId += 1
	return nil
}

func (i *Impl) GetNextID() int64 {
	return i.currentId
}

func (i *Impl) FindByID(userID int64) (*ads.User, error) {
	ad, ok := i.idToUser[userID]
	if !ok {
		return nil, errors.New("there is no user with such ID")
	}
	return ad, nil
}

func (i *Impl) FindByName(name string) (*ads.User, error) {
	for _, user := range i.idToUser {
		if user.Nickname == name {
			return user, nil
		}
	}
	return nil, ErrUserNotFound
}

func New() app.Repository[ads.User] {
	return &Impl{
		currentId: 0,
		idToUser:  make(map[int64]*ads.User),
	}
}
