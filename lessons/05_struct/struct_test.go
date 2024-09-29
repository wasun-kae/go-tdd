package gostruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {

	t.Run("should return perimeter from a given retangle", func(t *testing.T) {
		width := 10.0
		height := 5.0
		retangle := Retangle{width, height}

		expected := 30.0
		actual := Perimeter(retangle)

		assert.Equal(t, expected, actual)
	})
}
