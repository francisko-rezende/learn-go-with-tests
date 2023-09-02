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

func (d Dictionary) Add(word, definition string) {
	d[word] = definition
	// read here to find out why theres no need to use pointers here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps#pointers-copies-et-al
}

// To initialize an empty map do like so to avoid a runtime panic when trying to write on an empty map

// var dictionary = map[string]string{}

// // OR

// var dictionary = make(map[string]string)
