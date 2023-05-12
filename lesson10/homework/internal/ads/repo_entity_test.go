package ads

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestEntity struct {
	RepoEntity
}

func TestRepoEntity_SetID(t *testing.T) {
	entity := &TestEntity{}
	entity.SetID(0)
	assert.Equal(t, int64(0), entity.ID)
}

func TestRepoEntity_GetID(t *testing.T) {
	entity := &TestEntity{}
	entity.SetID(0)
	assert.Equal(t, int64(0), entity.GetID())
}

func TestRepoEntity_HasName(t *testing.T) {
	entity := &TestEntity{}
	entity.SetID(0)
	assert.Equal(t, false, entity.HasName(""))
}
