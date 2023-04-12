package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInArray(t *testing.T) {
	assert.Equal(t, InArray([]string{"a", "b"}, "a"), true)
	assert.Equal(t, InArray([]string{"a", "b"}, "c"), false)
}

func TestUniq(t *testing.T) {
	assert.ElementsMatch(t, Uniq([]string{"a", "a", "a", "b"}), []string{"a", "b"})
}
