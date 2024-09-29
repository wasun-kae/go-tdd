package method_and_receiver

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
		actual := retangle.Perimeter()

		assert.Equal(t, expected, actual)
	})
}

func TestArea(t *testing.T) {

	t.Run("should return area from a given retangle", func(t *testing.T) {
		retangle := Retangle{
			Width:  10.0,
			Height: 5.0,
		}

		expected := 50.0
		actual := retangle.Area()

		assert.Equal(t, expected, actual)
	})
}
