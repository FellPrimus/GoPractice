package mydict

import "errors"

//Dictionary type
type Dictionary map[string]string

var errNotFound = errors.New("Not Found")
var errCantUpdate = errors.New("Cant update non-existing word")
var errWordExists = errors.New("That word already exists")

//Search for a word
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", errNotFound
}

//Add a word to dictionary
func (d Dictionary) Add(word string, def string) error {
	_, err := d.Search(word)
	switch err {
	case errNotFound:
		d[word] = def
	case nil:
		return errWordExists
	}
	return nil
	/*if err == errNotFound {
		d[word] = def
	} else if err == nil {
		return errWordExists
	}
	return nil*/
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case nil:
		d[word] = definition
	case errNotFound:
		return errCantUpdate
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
