package main

import "errors"

var ErrNotFound = errors.New("could not find the word you were looking for")

//Create type for "thin" wrapper around Dictionary and give custom method
type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", errors.New("could not find the word")
	}
	return definition, nil
}

//Below code now redundant = refactored

//Takes a map named dictionary, a word and returns a string
//func Search(dictionary map[string]string, word string) string {
//	return dictionary[word]
//}
