// Package maps provides a simple dictionary implementation.
package maps

type Dictionary map[string]string

// Predefined errors
// An article: https://dave.cheney.net/2016/04/07/constant-errors
const (
	ErrNotFound   = DictionaryErr("could not find the word you were looking for")
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
)

type DictionaryErr string

// Error satisfies the error interface and returns the string value of e.
func (e DictionaryErr) Error() string {
	return string(e)
}

// Search looks up a word in the dictionary and returns its definition.
// If the word is found, it returns the definition and a nil error.
// If the word is not found, it returns an empty string and ErrNotFound.
func (d Dictionary) Search(word string) (string, error) {
	// The second value is a boolean which indicates if the key was found successfully.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

// Add a word and its definition to the dictionary.
//
// Returns ErrWordExists if the word already exists in the dictionary.
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}
