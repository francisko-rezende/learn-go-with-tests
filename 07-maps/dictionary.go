package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you're looking for")

func (d Dictionary) Search(searchTerm string) (string, error) {
	definition, ok := d[searchTerm]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, definition string) {
	d[key] = definition
}
