package jsonparser

import (
	"encoding/json"
	"os"

	"github.com/sarathkumar17/go-adventure/internal/adventure"
)

// ReadJSON reads the contents of a JSON file specified by the filename
// and returns the contents as a byte slice. If an error occurs while
// reading the file, it returns an error.
func ReadJSON(filename string) ([]byte, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

// ParseJSON takes a byte slice representing a JSON file and returns an
// adventure.Adventure representing the file's contents. If the JSON in the
// byte slice is invalid, an error is returned.
func ParseJSON(bs []byte) (adventure.Adventure, error) {
	var s adventure.Adventure
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return nil, err
	}
	return s, nil
}
