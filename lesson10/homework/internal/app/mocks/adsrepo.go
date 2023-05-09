package mocks

import (
	"homework10/internal/ads"
)

func NewAdsRepoMock() *AbstractRepoMock[*ads.Ad] {
	return NewAbstractRepoMock[*ads.Ad]()
}
