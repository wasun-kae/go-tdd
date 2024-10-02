package gomap

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(extension string) (string, error) {
	language, existed := dictionary[extension]

	if !existed {
		return language, errors.New("programming language does not exist")
	}

	return language, nil
}

func (dictionary Dictionary) Add(extension, language string) error {
	_, existed := dictionary[extension]

	if existed {
		return errors.New("programming language already exists")
	}

	dictionary[extension] = language

	return nil
}
