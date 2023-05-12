package mocks

import (
	"github.com/stretchr/testify/mock"
	"homework10/internal/adapters/baserepo"
	"homework10/internal/adapters/filters"
	"homework10/internal/ads"
)

type AbstractRepoMock[T ads.RepoEntityInterface] struct {
	repo baserepo.Repository[T]
	mock.Mock
}

func (a *AbstractRepoMock[T]) GetAll(f filters.Filters[T]) []T {
	a.Called(f)
	return a.repo.GetAll(f)
}

func (a *AbstractRepoMock[T]) Add(elem T) error {
	a.Called(elem)
	return a.repo.Add(elem)
}

func (a *AbstractRepoMock[T]) FindByID(id int64) (T, error) {
	a.Called(id)
	return a.repo.FindByID(id)
}

func (a *AbstractRepoMock[T]) FindByName(name string) (T, error) {
	a.Called(name)
	return a.repo.FindByName(name)
}

func (a *AbstractRepoMock[T]) DeleteById(id int64) (T, error) {
	a.Called(id)
	return a.repo.DeleteById(id)
}

func NewAbstractRepoMock[T ads.RepoEntityInterface]() *AbstractRepoMock[T] {
	return &AbstractRepoMock[T]{
		repo: baserepo.New[T](),
	}
}
