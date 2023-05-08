package ads

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUser_HasName(t *testing.T) {
	user := User{
		Nickname: "nickname",
	}
	tests := []Test{
		{In: "", Expect: false},
		{In: "lol", Expect: false},
		{In: "kek", Expect: false},
		{In: "n", Expect: false},
		{In: "nick", Expect: false},
		{In: "nickname", Expect: true},
	}
	for _, test := range tests {
		test := test
		t.Run("TestUser_HasName", func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, test.Expect, user.HasName(test.In))
		})
	}
}
