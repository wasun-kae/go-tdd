package gostruct

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPerimeter(t *testing.T) {

	t.Run("should return perimeter from a given retangle", func(t *testing.T) {
		retangle := Retangle{
			Width:  10.0,
			Height: 5.0,
		}

		expected := 30.0
		actual := Perimeter(retangle)

		assert.Equal(t, expected, actual)
	})
}
