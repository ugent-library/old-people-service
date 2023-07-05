package arrayutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUniq(t *testing.T) {
	assert.ElementsMatch(t, Uniq([]string{"a", "a", "a", "b"}), []string{"a", "b"})
}
