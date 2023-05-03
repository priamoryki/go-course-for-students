package adrepo

import (
	"homework10/internal/adapters/baserepo"
	"homework10/internal/ads"
	"homework10/internal/app"
)

func New() app.Repository[*ads.Ad] {
	return baserepo.New[*ads.Ad]()
}
