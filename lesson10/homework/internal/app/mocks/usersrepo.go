package mocks

import (
	"homework10/internal/ads"
)

func NewUsersRepoMock() *AbstractRepoMock[*ads.User] {
	return NewAbstractRepoMock[*ads.User]()
}
