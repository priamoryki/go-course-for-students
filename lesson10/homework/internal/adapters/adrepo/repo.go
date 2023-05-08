package adrepo

import (
	"homework10/internal/adapters/baserepo"
	"homework10/internal/ads"
)

func New() baserepo.Repository[*ads.Ad] {
	return baserepo.New[*ads.Ad]()
}
