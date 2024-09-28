package iteration

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeat(t *testing.T) {

	t.Run("should return repeated characters from given times", func(t *testing.T) {
		expected := "SSS"
		actual := Repeat("S", 3)

		assert.Equal(t, expected, actual)
	})
}
