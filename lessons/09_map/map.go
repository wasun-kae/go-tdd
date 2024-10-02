package gomap

import "errors"

type Dictionary map[string]string

func (dictionary Dictionary) Search(key string) (string, error) {
	value, ok := dictionary[key]

	if !ok {
		return value, errors.New("programming language does not exist")
	}

	return value, nil
}
