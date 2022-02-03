package main

import (
	"bytes"
	"errors"
	"io"
	"strings"
)

var (
	// ErrParseInput is returned if the input from stdin cannot be parsed
	ErrParseInput = errors.New("failed parse input from stdin")
	// ErrAbortInput is returned if the user aborted input action
	ErrAbortInput = errors.New("waiting for input was aborted")
	// ErrUnknownInput is returned if the user entered an unknown action
	ErrUnknownInput = errors.New("unknown answer entered")
)

// LineReader is an unbuffered line reader
type LineReader struct {
	reader io.Reader
}

// NewReader creates a new line reader
func NewReader(r io.Reader) *LineReader {
	return &LineReader{reader: r}
}

// Read implements io.Reader
func (lr LineReader) Read(p []byte) (int, error) {
	return lr.reader.Read(p)
}

// ReadLine reads one line without buffering
func (lr LineReader) ReadLine() (string, error) {
	output := &bytes.Buffer{}
	buffer := make([]byte, 1)

	for {
		n, err := lr.reader.Read(buffer)
		for i := 0; i < n; i++ {
			if buffer[i] == '\n' {
				return output.String(), nil
			}
			_ = output.WriteByte(buffer[i])
		}

		if err != nil {
			if err == io.EOF {
				return output.String(), nil
			}
			return output.String(), err
		}
	}
}

// AskForBool requires user bool input
func AskForBool(from io.Reader, def bool, skip bool) (bool, error) {
	if skip {
		return def, nil
	}

	choices := "y"
	if def {
		choices = "n"
	}

	str, err := AskForString(from, choices, skip)
	if err != nil {
		return false, ErrParseInput
	}

	switch str {
	case "y":
		return true, nil
	case "n":
		return false, nil
	}

	str = strings.ToLower(string(str[0]))
	switch str {
	case "y":
		return true, nil
	case "n":
		return false, nil
	case "q":
		return false, ErrAbortInput
	default:
		return false, ErrUnknownInput
	}
}

// AskForString requires user text input
func AskForString(from io.Reader, def string, skip bool) (string, error) {
	if skip {
		return def, nil
	}

	input, err := NewReader(from).ReadLine()
	if err != nil {
		return def, ErrParseInput
	}

	input = strings.TrimSpace(input)
	if input == "" {
		input = def
	}

	return input, nil
}
