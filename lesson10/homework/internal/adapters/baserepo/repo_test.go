package baserepo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestType struct {
	ID   int64
	Name string
}

func (t *TestType) GetID() int64 {
	return t.ID
}

func (t *TestType) SetID(ID int64) {
	t.ID = ID
}

func (t *TestType) HasName(name string) bool {
	return t.Name == name
}

func TestAdd(t *testing.T) {
	repo := New[*TestType]()

	entity := &TestType{ID: 123123, Name: "a"}
	err := repo.Add(entity)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)
}

func TestDeleteById(t *testing.T) {
	repo := New[*TestType]()

	entity := &TestType{Name: "a"}
	err := repo.Add(entity)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)

	entity, err = repo.DeleteById(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)

	entity, err = repo.DeleteById(0)
	assert.Error(t, err)
}

func TestFindByID(t *testing.T) {
	repo := New[*TestType]()

	entity, err := repo.FindByID(0)
	assert.Error(t, err)

	entity = &TestType{Name: "a"}
	err = repo.Add(entity)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)

	entity, err = repo.FindByID(0)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)
}

func TestFindByName(t *testing.T) {
	repo := New[*TestType]()

	entity, err := repo.FindByName("a")
	assert.Error(t, err)

	entity = &TestType{Name: "a"}
	err = repo.Add(entity)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), entity.ID)

	entity, err = repo.FindByName("a")
	assert.NoError(t, err)
	assert.Equal(t, "a", entity.Name)
}
