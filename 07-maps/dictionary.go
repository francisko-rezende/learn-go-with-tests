package maps

const (
	ErrNotFound   = DictionaryErr("could not find the word you're looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Update(key, updatedDefinition string) {
	d[key] = updatedDefinition
}

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
