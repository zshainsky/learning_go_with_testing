package main

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists")
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, val string) error {
	_, err := d.Search(word)

	switch err {
	// If word not found
	case ErrNotFound:
		d[word] = val
	// Word found
	case nil:
		return ErrWordExists
	default:
		return err

	}
	return nil
}

func (d Dictionary) Update(word, val string) error {
	_, err := d.Search(word)
	switch err {
	// Word found
	case nil:
		d[word] = val
	case ErrNotFound:
		return ErrWordDoesNotExist
	default:
		return err
	}
	return nil
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}

func (e DictionaryErr) Error() string {
	return string(e)
}
