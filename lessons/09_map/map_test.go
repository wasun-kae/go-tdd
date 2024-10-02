package gomap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {

	t.Run("should return programming language from given file extension", func(t *testing.T) {
		dictionary := Dictionary{
			"go": "Go",
		}

		expected := "Go"
		actual, err := dictionary.Search("go")

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	t.Run("should return error if programming language does NOT exists from given file extension", func(t *testing.T) {
		dictionary := Dictionary{
			"go": "Go",
		}

		expected := ""
		actual, err := dictionary.Search("java")

		assert.Equal(t, expected, actual)
		assert.EqualError(t, err, "programming language does not exist")
	})
}
