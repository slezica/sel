package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const outOfBounds = -1

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
	if expr == "" {
		return nil, NewParseSelectorError(expr)
	}

	exprParts := strings.Split(expr, ":")

	start, err := parseIndex(exprParts[0], 1)
	if err != nil {
		return nil, NewParseSelectorError(expr)
	}

	if len(exprParts) == 1 {
		return &Selector{start: start, end: start}, nil
	}

	end, err := parseIndex(exprParts[1], math.MaxInt32)
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

	if low == outOfBounds || high == outOfBounds || low > high {
		return []string{}
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

func parseIndex(expr string, defaultValue int) (int, error) {
	if len(expr) == 0 {
		return defaultValue, nil
	}

	value, err := strconv.Atoi(expr)

	if value == 0 || err != nil {
		return -1, fmt.Errorf("Invalid index %s (cause %v)", expr, err)
	}

	return value, nil
}

func adjustStartIndex(index int, fieldCount int) int {
	if index == 0 {
		panic("Internal error: invalid Selector was built")
	}

	if fieldCount == 0 {
		return outOfBounds // no possible selection for 0 fields, anything is out of bounds

	} else if index > fieldCount {
		return outOfBounds // selection starts after fields end, this is out of bounds

	} else if index < -fieldCount {
		return 0 // selection starts before fields do, begin at the first index

	} else if index < 0 {
		return (fieldCount - 1) - (abs(index) - 1) // valid negative selection start, count from the last index

	} else {
		return index - 1 // valid positive selection start, count from the first index
	}
}

func adjustEndIndex(index int, fieldCount int) int {
	if index == 0 {
		panic("Internal error: invalid Selector was built")
	}

	if fieldCount == 0 {
		return outOfBounds // no possible selection for 0 fields, anything is out of bounds

	} else if index < -fieldCount {
		return outOfBounds // selection ends before fields begin, this is out of bounds

	} else if index >= fieldCount {
		return fieldCount // selection ends after fields do, end at the last index

	} else if index < 0 {
		return fieldCount - abs(index) + 1 // valid negative selection end, count from the last index (inclusive)

	} else {
		return index // valid positive selection end, count from the first index (inclusive)
	}
}

func abs(x int) int {
	return int(math.Abs(float64(x)))
}
