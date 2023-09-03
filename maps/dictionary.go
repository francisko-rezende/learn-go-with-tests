package maps

type Dictionary map[string]string

const (
	ErrNotfound         = DictionaryErr("could not find the word you wre looking for")
	ErrWordExists       = DictionaryErr("word already exists in dictionary")
	ErrWordDoesNotExist = DictionaryErr("word already exists in dictionary")
)

// read here to find more info on why it's cool to use contant errors https://dave.cheney.net/2016/04/07/constant-errors

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotfound
	}
	return definition, nil
}

func (d Dictionary) Edit(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotfound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = definition
	default:
		return err
	}

	return nil
}

func (d Dictionary) Add(word, definition string) error {

	_, err := d.Search(word)

	switch err {
	case ErrNotfound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
	// read here to find out why theres no need to use pointers here: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/maps#pointers-copies-et-al
}

// To initialize an empty map do like so to avoid a runtime panic when trying to write on an empty map

// var dictionary = map[string]string{}

// // OR

// var dictionary = make(map[string]string)
