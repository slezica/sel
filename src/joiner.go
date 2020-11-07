package main

import "strings"

// Joiner can join fields using a delimiter
type Joiner struct {
	Delim string
}

// ParseJoiner creates a Joiner from a user-provided string
func ParseJoiner(expr string) (*Joiner, error) {
	return &Joiner{Delim: expr}, nil // no possible error for now
}

// Join concatenates fields using the delimiter between them
func (j *Joiner) Join(fields []string) string {
	return strings.Join(fields, j.Delim)
}
