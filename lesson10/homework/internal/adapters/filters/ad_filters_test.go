package filters

import (
	"github.com/stretchr/testify/assert"
	"homework10/internal/ads"
	"testing"
	"time"
)

func TestFilterNonPublished(t *testing.T) {
	filter := NewFilterNonPublished()
	ad1 := &ads.Ad{Published: true}
	ad2 := &ads.Ad{Published: false}

	tests := []Test[*ads.Ad]{
		{In: []*ads.Ad{ad2, ad1}, Expect: []*ads.Ad{ad1}},
		{In: []*ads.Ad{ad1, ad2, ad1}, Expect: []*ads.Ad{ad1, ad1}},
	}

	for _, test := range tests {
		test := test
		t.Run("TestFilterNonPublished", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, filter.Filter(test.In))
		})
	}
}

func TestFilterByAuthor(t *testing.T) {
	filter := NewFilterByAuthor()
	ad1 := &ads.Ad{AuthorID: 0}
	ad2 := &ads.Ad{AuthorID: 1}

	tests := []Test[*ads.Ad]{
		{In: []*ads.Ad{ad2, ad1}, Expect: []*ads.Ad{ad1, ad2}},
		{In: []*ads.Ad{ad1, ad2, ad1}, Expect: []*ads.Ad{ad1, ad1, ad2}},
	}

	for _, test := range tests {
		test := test
		t.Run("TestFilterByAuthor", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, filter.Filter(test.In))
		})
	}
}

func TestFilterByCreationTime(t *testing.T) {
	filter := NewFilterByCreationTime()
	curTime := time.Now().UTC()
	ad1 := &ads.Ad{CreationTime: curTime}
	ad2 := &ads.Ad{CreationTime: curTime.Add(time.Minute)}

	tests := []Test[*ads.Ad]{
		{In: []*ads.Ad{ad2, ad1}, Expect: []*ads.Ad{ad1, ad2}},
		{In: []*ads.Ad{ad1, ad2, ad1}, Expect: []*ads.Ad{ad1, ad1, ad2}},
	}

	for _, test := range tests {
		test := test
		t.Run("TestFilterByAuthor", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, filter.Filter(test.In))
		})
	}
}
