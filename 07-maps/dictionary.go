package maps

type Dictionary map[string]string

func (d Dictionary) Search(searchTerm string) string {
	return d[searchTerm]
}

func Search(dictionary map[string]string, searchTerm string) string {
	return dictionary[searchTerm]
}
