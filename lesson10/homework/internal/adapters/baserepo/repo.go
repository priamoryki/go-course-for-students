package baserepo

import (
	"errors"
	"homework10/internal/adapters/filters"
	"homework10/internal/ads"
	"homework10/internal/app"
	"sync"
)

var ErrNotFound = errors.New("element not found")

func getZeroValue[T any]() T {
	var result T
	return result
}

type Impl[T ads.RepoEntityInterface] struct {
	currentId int64
	idToElem  map[int64]T
	mutex     *sync.RWMutex
}

func (i *Impl[T]) GetAll(f filters.Filters[T]) []T {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	result := make([]T, 0)
	for j := int64(0); j < i.currentId; j++ {
		result = append(result, i.idToElem[j])
	}
	return f.Filter(result)
}

func (i *Impl[T]) Add(elem T) error {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	elem.SetID(i.currentId)
	i.idToElem[elem.GetID()] = elem
	i.currentId += 1
	return nil
}

func (i *Impl[T]) FindByID(id int64) (T, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	elem, ok := i.idToElem[id]
	if !ok {
		return getZeroValue[T](), ErrNotFound
	}
	return elem, nil
}

func (i *Impl[T]) FindByName(name string) (T, error) {
	i.mutex.RLock()
	defer i.mutex.RUnlock()
	for _, elem := range i.idToElem {
		if elem.HasName(name) {
			return elem, nil
		}
	}
	return getZeroValue[T](), ErrNotFound
}

func (i *Impl[T]) DeleteById(id int64) (T, error) {
	i.mutex.Lock()
	defer i.mutex.Unlock()
	elem, ok := i.idToElem[id]
	if !ok {
		return getZeroValue[T](), ErrNotFound
	}
	delete(i.idToElem, id)
	return elem, nil
}

func New[T ads.RepoEntityInterface]() app.Repository[T] {
	return &Impl[T]{
		currentId: 0,
		idToElem:  make(map[int64]T),
		mutex:     new(sync.RWMutex),
	}
}
