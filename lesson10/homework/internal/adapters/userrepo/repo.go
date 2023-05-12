package userrepo

import (
	"homework10/internal/adapters/baserepo"
	"homework10/internal/ads"
)

func New() baserepo.Repository[*ads.User] {
	return baserepo.New[*ads.User]()
}
