package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

/* The approach for mocking os.Stdin came from https://stackoverflow.com/questions/46365221/fill-os-stdin-for-function-that-reads-from-it */

func TestPromptForNameContinuesUntilNonEmpty(t *testing.T) {
	content := "   \n\t\nWoody"
	tempFile, err := writeToATempFile(content)
	if tempFile != nil {
		defer os.Remove(tempFile.Name())
	}
	if err != nil {
		t.Error(err)
	}

	oldStdIn := os.Stdin
	defer func() { os.Stdin = oldStdIn }()

	os.Stdin = tempFile

	name := promptForName()
	if name != "Woody" {
		t.Errorf("Name is %s, expected Woody", name)
	}

	if err := tempFile.Close(); err != nil {
		t.Error("Unable to close the temp file", err)
	}
}

func TestPromptForNameTrimsSpaces(t *testing.T) {
	content := "   Woody  \t"
	tempFile, err := writeToATempFile(content)
	if tempFile != nil {
		defer os.Remove(tempFile.Name())
	}
	if err != nil {
		t.Error(err)
	}

	oldStdIn := os.Stdin
	defer func() { os.Stdin = oldStdIn }()

	os.Stdin = tempFile

	name := promptForName()
	if name != "Woody" {
		t.Errorf("Name is %s, expected Woody", name)
	}

	if err := tempFile.Close(); err != nil {
		t.Error("Unable to close the temp file", err)
	}
}

func writeToATempFile(content string) (*os.File, error) {
	tempFile, err := ioutil.TempFile(".", "promptForName")
	if err != nil {
		return tempFile, fmt.Errorf("Unable to write contents to the temp file: %s", err)
	}

	if _, err := tempFile.WriteString(content); err != nil {
		return tempFile, fmt.Errorf("Unable to write contents to the temp file: %s", err)
	}

	if _, err := tempFile.Seek(0, 0); err != nil {
		return tempFile, fmt.Errorf("Unable to seek to the beginning of the temp file: %s", err)
	}

	return tempFile, nil
}
