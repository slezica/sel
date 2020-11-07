package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Selector can select fields from a slice using human-friendly indexes, negative and positive
type Selector struct {
	start int
	end   int
}

// ParseSelectorError indicates an error parsing a user-provided string into a Selector
type ParseSelectorError struct {
	Expr string
}

// ParseSelector creates a Selector from a user-provided string
func ParseSelector(expr string) (*Selector, error) {
	exprParts := strings.Split(expr, ":")

	if len(exprParts) == 0 {
		return &Selector{start: 0, end: 0}, nil
	}

	start, err := strconv.Atoi(exprParts[0])
	if err != nil {
		return nil, NewParseSelectorError(expr)
	}

	if len(exprParts) == 1 {
		return &Selector{start: start, end: start}, nil
	}

	end, err := strconv.Atoi(exprParts[1])
	if err != nil {
		return nil, NewParseSelectorError(expr)
	}

	if len(exprParts) == 2 {
		return &Selector{start: start, end: end}, nil
	}

	return nil, NewParseSelectorError(expr)
}

// Select returns a slice of fields according to the Selector criterion
func (s *Selector) Select(fields []string) []string {
	fieldCount := len(fields)

	low := adjustStartIndex(s.start, fieldCount)
	high := adjustEndIndex(s.end, fieldCount)

	if low == high {
		return []string{fields[low]}
	}

	return fields[low:high]
}

// NewParseSelectorError creates a ParseSelectorError from a user-provided string
func NewParseSelectorError(expr string) *ParseSelectorError {
	return &ParseSelectorError{Expr: expr}
}

func (e *ParseSelectorError) Error() string {
	return fmt.Sprintf("Invalid selector '%s'", e.Expr)
}

func adjustStartIndex(index int, fieldCount int) int {
	// NOTE:
	// This function is a little redundant, but separating cases helps me reason about them.

	if index == 0 || fieldCount == 0 {
		return 0

	} else if index >= fieldCount {
		return fieldCount - 1

	} else if index < -fieldCount {
		return 0

	} else if index < 0 {
		return (fieldCount - 1) - (abs(index) - 1)

	} else {
		return index - 1
	}
}

func adjustEndIndex(index int, fieldCount int) int {
	// NOTE:
	// This function is a little redundant, but separating cases helps me reason about them.

	if index == 0 || fieldCount == 0 {
		return 0

	} else if index >= fieldCount {
		return fieldCount

	} else if index < -fieldCount {
		return 0

	} else if index < 0 {
		return fieldCount - abs(index) + 1

	} else {
		return index
	}
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
