package maps

import "errors"

type Dictionary map[string]string

var WordNotFoundWarning = "could not find the word you're looking for"

func (d Dictionary) Search(searchTerm string) (string, error) {
	definition, ok := d[searchTerm]

	if !ok {
		return "", errors.New(WordNotFoundWarning)
	}
	return definition, nil
}
