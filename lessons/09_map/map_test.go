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

func TestAdd(t *testing.T) {
	t.Run("should add new programming language if given file extension does NOT exist", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Add("go", "Go")
		expected := "Go"
		actual := dictionary["go"]

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	t.Run("should return error if given file extension already exists", func(t *testing.T) {
		dictionary := Dictionary{
			"go": "Go",
		}

		err := dictionary.Add("go", "Go")

		assert.EqualError(t, err, "programming language already exists")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("should update existing programming language if given file extension alreadys exists", func(t *testing.T) {
		dictionary := Dictionary{
			"go": "Go",
		}

		err := dictionary.Update("go", "Go - updated")

		expected := "Go - updated"
		actual := dictionary["go"]

		assert.Equal(t, expected, actual)
		assert.Nil(t, err)
	})

	t.Run("should return error if given file extension does NOT exist", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("go", "Go - updated")

		assert.EqualError(t, err, "programming language does not exist")
	})
}
