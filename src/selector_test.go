package main

import (
	"errors"
	"testing"
)

func TestParse(t *testing.T) {
	tryParseSelector(t, "1", false)
	tryParseSelector(t, "-1", false)
	tryParseSelector(t, "1:1", false)
	tryParseSelector(t, "-1:1", false)
	tryParseSelector(t, "-1:-1", false)

	tryParseSelector(t, "", true)
	tryParseSelector(t, "1:", true)
	tryParseSelector(t, ":1", true)
	tryParseSelector(t, ":-1", true)
	tryParseSelector(t, "1:", true)
	tryParseSelector(t, "-1:", true)
	tryParseSelector(t, "hello", true)
	tryParseSelector(t, "hello", true)
	tryParseSelector(t, "hello:", true)
	tryParseSelector(t, "hello:bye", true)
	tryParseSelector(t, ":bye", true)
}

func TestSelect(t *testing.T) {
	// NOTE:
	// This is only testing the case where start == end. Other behavior is tested elsehwere in this file.

	selector := Selector{start: 1, end: 1}
	fields := []string{"a", "b", "c"}
	result := selector.Select(fields)

	if len(result) != 1 || result[0] != "a" {
		t.Errorf("Expected 'a' when selecting 1:1 from 'a b c', got %v", result)
	}
}

func TestAdjustStart(t *testing.T) {
	tryAdjustStart(t, -1, 0, 0)
	tryAdjustStart(t, 0, 0, 0)
	tryAdjustStart(t, 1, 0, 0)

	tryAdjustStart(t, -2, 1, 0)
	tryAdjustStart(t, -1, 1, 0)
	tryAdjustStart(t, 0, 1, 0)
	tryAdjustStart(t, 1, 1, 0)
	tryAdjustStart(t, 2, 1, 0)

	tryAdjustStart(t, -1, 3, 2)
	tryAdjustStart(t, -2, 3, 1)
	tryAdjustStart(t, -3, 3, 0)
	tryAdjustStart(t, -4, 3, 0)
	tryAdjustStart(t, 0, 3, 0)
	tryAdjustStart(t, 1, 3, 0)
	tryAdjustStart(t, 2, 3, 1)
	tryAdjustStart(t, 3, 3, 2)
	tryAdjustStart(t, 4, 3, 2)
}

func TestAdjustEnd(t *testing.T) {
	tryAdjustEnd(t, -1, 0, 0)
	tryAdjustEnd(t, 0, 0, 0)
	tryAdjustEnd(t, 1, 0, 0)

	tryAdjustEnd(t, -2, 1, 0)
	tryAdjustEnd(t, -1, 1, 1)
	tryAdjustEnd(t, 0, 1, 0)
	tryAdjustEnd(t, 1, 1, 1)
	tryAdjustEnd(t, 2, 1, 1)

	tryAdjustEnd(t, -1, 3, 3)
	tryAdjustEnd(t, -2, 3, 2)
	tryAdjustEnd(t, -3, 3, 1)
	tryAdjustEnd(t, -4, 3, 0)
	tryAdjustEnd(t, 0, 3, 0)
	tryAdjustEnd(t, 1, 3, 1)
	tryAdjustEnd(t, 2, 3, 2)
	tryAdjustEnd(t, 3, 3, 3)
	tryAdjustEnd(t, 4, 3, 3)
}

func tryAdjustStart(t *testing.T, start int, fieldCount int, expected int) {
	actual := adjustStartIndex(start, fieldCount)

	if actual != expected {
		t.Errorf("Start index %d of %d adjusted to %d expecting %d", start, fieldCount, actual, expected)
	}
}

func tryAdjustEnd(t *testing.T, end int, fieldCount int, expected int) {
	actual := adjustEndIndex(end, fieldCount)

	if actual != expected {
		t.Errorf("End index %d of %d adjusted to %d expecting %d", end, fieldCount, actual, expected)
	}
}

func tryParseSelector(t *testing.T, expr string, expectError bool) {
	_, err := ParseSelector(expr)
	hasError := (err != nil)

	if expectError == hasError {
		if expectError {
			var psErr *ParseSelectorError

			if !errors.As(err, &psErr) {
				t.Errorf("Error <%v> parsing selector '%v' expecting a ParseSelectorError", err, expr)
			}
		}

	} else {
		t.Errorf("Error <%v> parsing selector %v when expectError was %v", err, expr, expectError)
	}
}
