package main

import (
	"regexp"
)

// Splitter can split lines of text into fields
type Splitter struct {
	re *regexp.Regexp
}

// ParseSplitter creates a Splitter from a user-provided string
func ParseSplitter(expr string) (*Splitter, error) {
	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, err
	}

	return &Splitter{re: re}, nil
}

// Split returns a slice of all fields found in text
func (s *Splitter) Split(text string) []string {
	return s.re.Split(text, -1)
}
