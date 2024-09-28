package lesson_01_hello_world

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

	t.Run("should return hello world message", func(t *testing.T) {
		expected := "Hello, World"
		actual := HelloWorld()

		assert.Equal(t, expected, actual)
	})
}
