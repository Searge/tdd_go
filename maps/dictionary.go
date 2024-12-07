package maps

import "fmt"

type Dictionary map[string]string

// Search looks up a word in the dictionary and returns its definition.
// If the word is not present in the dictionary, it returns an empty string.
func (d Dictionary) Search(word string) (string, error) {
	// The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", fmt.Errorf("could not find the word you were looking for")
	}
	return definition, nil
}
