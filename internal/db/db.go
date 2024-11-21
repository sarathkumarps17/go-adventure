package db

import (
	"fmt"

	"github.com/sarathkumar17/go-adventure/internal/adventure"

	jsonparser "github.com/sarathkumar17/go-adventure/pkg/json_parser"
)

func StartDb(filename string) (adventure.Adventure, error) {
	bs, err := jsonparser.ReadJSON(filename)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return nil, err
	}
	return jsonparser.ParseJSON(bs)
}
