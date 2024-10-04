package gomap

import "errors"

type Dictionary map[string]string

var (
	errItemNotExistErr    = errors.New("programming language does not exist")
	erroItemAlreadyExists = errors.New("programming language already exists")
)

func (dictionary Dictionary) Search(extension string) (string, error) {
	language, existed := dictionary[extension]

	if !existed {
		return language, errItemNotExistErr
	}

	return language, nil
}

func (dictionary Dictionary) Add(extension, language string) error {
	_, existed := dictionary[extension]

	if existed {
		return erroItemAlreadyExists
	}

	dictionary[extension] = language

	return nil
}
