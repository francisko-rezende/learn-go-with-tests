package maps

import "errors"

type Dictionary map[string]string

var ErrNotfound = errors.New("could not find the word you wre looking for")

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotfound
	}
	return definition, nil
}
