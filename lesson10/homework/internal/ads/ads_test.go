package ads

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Test struct {
	In     string
	Expect bool
}

func TestAd_HasName(t *testing.T) {
	ad := Ad{
		Title: "title",
		Text:  "text",
	}
	tests := []Test{
		{In: "lol", Expect: false},
		{In: "kek", Expect: false},
		{In: "", Expect: true},
		{In: "t", Expect: true},
		{In: "tit", Expect: true},
		{In: "title", Expect: true},
	}
	for _, test := range tests {
		test := test
		t.Run("TestAd_HasName", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, ad.HasName(test.In))
		})
	}
}
