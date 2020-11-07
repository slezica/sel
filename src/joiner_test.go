package main

import "testing"

func TestParseJoiner(t *testing.T) {
	tryParseJoiner(t, "")
	tryParseJoiner(t, "-")
	tryParseJoiner(t, "hello world")
	tryParseJoiner(t, "this\n")
}

func tryParseJoiner(t *testing.T, expr string) {
	_, err := ParseJoiner(expr)

	if err != nil {
		t.Errorf("Error <%v> when parsing joiner '%v", err, expr)
	}
}
