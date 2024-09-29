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

func TestSumAllTails(t *testing.T) {

	t.Run("should return sum of each given arrays of numbers except first number in array", func(t *testing.T) {
		number1 := []int{1, 2}
		number2 := []int{3, 4, 5}

		expected := []int{2, 9}
		actual := SumAllTails(number1, number2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return zero if a given array is empty", func(t *testing.T) {
		number1 := []int{}
		number2 := []int{3, 4, 5}

		expected := []int{0, 9}
		actual := SumAllTails(number1, number2)

		assert.Equal(t, expected, actual)
	})

	t.Run("should return zero if a given array has only one element", func(t *testing.T) {
		number1 := []int{1}
		number2 := []int{3, 4, 5}

		expected := []int{0, 9}
		actual := SumAllTails(number1, number2)

		assert.Equal(t, expected, actual)
	})
}
