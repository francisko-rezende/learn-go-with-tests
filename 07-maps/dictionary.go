package maps

import "errors"

type Dictionary map[string]string

var (
	ErrNotFound   = errors.New("could not find the word you're looking for")
	ErrWordExists = errors.New("cannot add word because it already exists")
)

func (d Dictionary) Search(searchTerm string) (string, error) {
	definition, ok := d[searchTerm]

	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(key, definition string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
