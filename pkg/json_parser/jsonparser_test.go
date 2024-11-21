package jsonparser_test

import (
	"os"
	"testing"

	jsonparser "github.com/sarathkumar17/go-adventure/pkg/json_parser"
)

// Successfully reads a valid JSON file and returns its content
func TestReadJSONSuccess(t *testing.T) {
	filename := "test.json"
	expectedContent := []byte(`{"key": "value"}`)

	err := os.WriteFile(filename, expectedContent, 0644)
	if err != nil {
		t.Fatalf("Failed to create test file: %v", err)
	}
	defer os.Remove(filename)

	content, err := jsonparser.ReadJSON(filename)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if string(content) != string(expectedContent) {
		t.Errorf("Expected %s, got %s", expectedContent, content)
	}
}

// File does not exist, should return an error
func TestReadJSONFileNotFound(t *testing.T) {
	filename := "nonexistent.json"

	_, err := jsonparser.ReadJSON(filename)
	if err == nil {
		t.Fatal("Expected an error, got nil")
	}
}
