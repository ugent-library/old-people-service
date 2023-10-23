package gismo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniqStrings(t *testing.T) {
	assert.Equal(t, UniqStrings([]string{"a", "a", "b", "b"}), []string{"a", "b"})
	assert.Equal(t, UniqStrings([]string{"a", "b"}), []string{"a", "b"})
}
