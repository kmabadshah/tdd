package main

type Dict map[string]string

const (
	ErrNotFound   = DictionaryErr("word doesn't exist")
	ErrWordExists = DictionaryErr("word exists")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dict) Search(word string) (string, error) {
	def, exists := d[word]

	if !exists {
		return "", ErrNotFound
	}

	return def, nil
}

func (d Dict) Add(word, def string) error {
	if d[word] != "" {
		return ErrWordExists
	}

	d[word] = def
	return nil
}

func (d Dict) Update(word, newDef string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	d[word] = newDef
	return nil
}

func (d Dict) Delete(word string) error {
	_, err := d.Search(word)
	if err != nil {
		return err
	}

	delete(d, word)
	return nil
}
