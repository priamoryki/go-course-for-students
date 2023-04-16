package ads

import (
	"strings"
	"time"
)

type Ad struct {
	RepoEntity
	Title          string
	Text           string
	AuthorID       int64
	Published      bool
	CreationTime   time.Time
	LastUpdateTime time.Time
}

func (ad *Ad) HasName(name string) bool {
	return strings.HasPrefix(ad.Title, name)
}
