package integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {

	t.Run("should sum of two integers", func(t *testing.T) {
		expected := 4
		actual := Add(2, 2)

		assert.Equal(t, expected, actual)
	})
}
