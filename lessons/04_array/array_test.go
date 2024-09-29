package array

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {

	t.Run("should return sum from given array of numbers in any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		expected := 15
		actual := Sum(numbers)

		assert.Equal(t, expected, actual)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should return sum of each given arrays of numbers", func(t *testing.T) {
		number1 := []int{1, 2}
		number2 := []int{3, 4, 5}

		expected := []int{3, 12}
		actual := SumAll(number1, number2)

		assert.Equal(t, expected, actual)
	})
}
