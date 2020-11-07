package main

import (
	"reflect"
	"testing"
)

func TestSplitterSplit(t *testing.T) {
	trySplitter(t, "", "", []string{})
	trySplitter(t, "", "abc", []string{"a", "b", "c"})
	trySplitter(t, "-", "a-b-c", []string{"a", "b", "c"})
	trySplitter(t, "\\s+", "a   b c", []string{"a", "b", "c"})
}

func trySplitter(t *testing.T, expr string, text string, expected []string) {
	splitter, err := ParseSplitter(expr)
	if err != nil {
		t.Errorf("Failed to create Splitter: %v", err)
	}

	actual := splitter.Split(text)

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Expression '%s' split '%s' into %v expecting %v", expr, text, actual, expected)
	}
}
