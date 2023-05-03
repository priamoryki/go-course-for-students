package userrepo

import (
	"homework10/internal/adapters/baserepo"
	"homework10/internal/ads"
	"homework10/internal/app"
)

func New() app.Repository[*ads.User] {
	return baserepo.New[*ads.User]()
}
