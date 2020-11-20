package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestParse(t *testing.T) {
	tryParseSelector(t, "1", false)
	tryParseSelector(t, "-1", false)
	tryParseSelector(t, "1:1", false)
	tryParseSelector(t, "-1:1", false)
	tryParseSelector(t, "-1:-1", false)

	tryParseSelector(t, "1:", false)
	tryParseSelector(t, ":1", false)
	tryParseSelector(t, ":-1", false)
	tryParseSelector(t, "1:", false)
	tryParseSelector(t, "-1:", false)

	tryParseSelector(t, "", true)
	tryParseSelector(t, "0", true)
	tryParseSelector(t, "0:0", true)
	tryParseSelector(t, "0:", true)
	tryParseSelector(t, ":0", true)
	tryParseSelector(t, "hello", true)
	tryParseSelector(t, "hello", true)
	tryParseSelector(t, "hello:", true)
	tryParseSelector(t, "hello:bye", true)
	tryParseSelector(t, ":bye", true)
}

func TestSelectEmpty(t *testing.T) {
	st := newSelectorTester(t, []string{})

	// st.try("-2:-2", []string{})
	st.try("-2:-1", []string{})
	st.try("-2:1", []string{})
	st.try("-2:2", []string{})

	st.try("-1:-2", []string{})
	st.try("-1:-1", []string{})
	st.try("-1:1", []string{})
	st.try("-1:2", []string{})

	st.try("1:-2", []string{})
	st.try("1:-1", []string{})
	st.try("1:1", []string{})
	st.try("1:2", []string{})

	st.try("2:-2", []string{})
	st.try("2:-1", []string{})
	st.try("2:1", []string{})
	st.try("2:2", []string{})

	st.try(":1", []string{})
	st.try(":2", []string{})
	st.try("1:", []string{})
	st.try("2:", []string{})

	st.try(":-1", []string{})
	st.try(":-2", []string{})
	st.try("-1:", []string{})
	st.try("-2:", []string{})

	st.try(":", []string{})
}

func TestSelect1(t *testing.T) {
	st := newSelectorTester(t, []string{"a"})

	st.try("-2:-2", []string{})
	st.try("-2:-1", []string{"a"})
	st.try("-2:1", []string{"a"})
	st.try("-2:2", []string{"a"})
	st.try("-2:2", []string{"a"})

	st.try("-1:-2", []string{})
	st.try("-1:-1", []string{"a"})
	st.try("-1:1", []string{"a"})
	st.try("-1:2", []string{"a"})

	st.try("1:-2", []string{})
	st.try("1:-1", []string{"a"})
	st.try("1:1", []string{"a"})
	st.try("1:2", []string{"a"})

	st.try("2:-2", []string{})
	st.try("2:-1", []string{})
	st.try("2:1", []string{})
	st.try("2:2", []string{})

	st.try(":1", []string{"a"})
	st.try(":2", []string{"a"})
	st.try("1:", []string{"a"})
	st.try("2:", []string{})

	st.try(":-1", []string{"a"})
	st.try(":-2", []string{})
	st.try("-1:", []string{"a"})
	st.try("-2:", []string{"a"})

	st.try(":", []string{"a"})
}

func TestSelect2(t *testing.T) {
	st := newSelectorTester(t, []string{"a", "b"})

	st.try("-3:-2", []string{"a"})
	st.try("-3:-1", []string{"a", "b"})
	st.try("-3:1", []string{"a"})
	st.try("-3:2", []string{"a", "b"})

	st.try("-2:-3", []string{})
	st.try("-2:-2", []string{"a"})
	st.try("-2:-1", []string{"a", "b"})
	st.try("-2:1", []string{"a"})
	st.try("-2:2", []string{"a", "b"})
	st.try("-2:3", []string{"a", "b"})

	st.try("-1:-3", []string{})
	st.try("-1:-2", []string{})
	st.try("-1:-1", []string{"b"})
	st.try("-1:1", []string{})
	st.try("-1:2", []string{"b"})
	st.try("-1:3", []string{"b"})

	st.try("1:-3", []string{})
	st.try("1:-2", []string{"a"})
	st.try("1:-1", []string{"a", "b"})
	st.try("1:1", []string{"a"})
	st.try("1:2", []string{"a", "b"})
	st.try("1:3", []string{"a", "b"})

	st.try("2:-3", []string{})
	st.try("2:-2", []string{})
	st.try("2:-1", []string{"b"})
	st.try("2:1", []string{})
	st.try("2:2", []string{"b"})
	st.try("2:3", []string{"b"})

	st.try("3:-3", []string{})
	st.try("3:-2", []string{})
	st.try("3:-1", []string{})
	st.try("3:1", []string{})
	st.try("3:2", []string{})
	st.try("3:3", []string{})

	st.try(":1", []string{"a"})
	st.try(":2", []string{"a", "b"})
	st.try(":3", []string{"a", "b"})
	st.try("1:", []string{"a", "b"})
	st.try("2:", []string{"b"})
	st.try("3:", []string{})

	st.try(":-1", []string{"a", "b"})
	st.try(":-2", []string{"a"})
	st.try(":-3", []string{})
	st.try("-1:", []string{"b"})
	st.try("-2:", []string{"a", "b"})
	st.try("-3:", []string{"a", "b"})

	st.try(":", []string{"a", "b"})
}

func TestSelect3(t *testing.T) {
	st := newSelectorTester(t, []string{"a", "b", "c"})

	st.try("-4:-4", []string{})
	st.try("-4:-3", []string{"a"})
	st.try("-4:-2", []string{"a", "b"})
	st.try("-4:-1", []string{"a", "b", "c"})
	st.try("-4:1", []string{"a"})
	st.try("-4:2", []string{"a", "b"})
	st.try("-4:3", []string{"a", "b", "c"})
	st.try("-4:4", []string{"a", "b", "c"})

	st.try("-3:-4", []string{})
	st.try("-3:-3", []string{"a"})
	st.try("-3:-2", []string{"a", "b"})
	st.try("-3:-1", []string{"a", "b", "c"})
	st.try("-3:1", []string{"a"})
	st.try("-3:2", []string{"a", "b"})
	st.try("-3:3", []string{"a", "b", "c"})
	st.try("-3:4", []string{"a", "b", "c"})

	st.try("-2:-4", []string{})
	st.try("-2:-3", []string{})
	st.try("-2:-2", []string{"b"})
	st.try("-2:-1", []string{"b", "c"})
	st.try("-2:1", []string{})
	st.try("-2:2", []string{"b"})
	st.try("-2:3", []string{"b", "c"})
	st.try("-2:4", []string{"b", "c"})

	st.try("-1:-4", []string{})
	st.try("-1:-3", []string{})
	st.try("-1:-2", []string{})
	st.try("-1:-1", []string{"c"})
	st.try("-1:1", []string{})
	st.try("-1:2", []string{})
	st.try("-1:3", []string{"c"})
	st.try("-1:4", []string{"c"})

	st.try("1:-4", []string{})
	st.try("1:-3", []string{"a"})
	st.try("1:-2", []string{"a", "b"})
	st.try("1:-1", []string{"a", "b", "c"})
	st.try("1:1", []string{"a"})
	st.try("1:2", []string{"a", "b"})
	st.try("1:3", []string{"a", "b", "c"})
	st.try("1:4", []string{"a", "b", "c"})

	st.try("2:-4", []string{})
	st.try("2:-3", []string{})
	st.try("2:-2", []string{"b"})
	st.try("2:-1", []string{"b", "c"})
	st.try("2:1", []string{})
	st.try("2:2", []string{"b"})
	st.try("2:3", []string{"b", "c"})
	st.try("2:4", []string{"b", "c"})

	st.try("3:-4", []string{})
	st.try("3:-3", []string{})
	st.try("3:-2", []string{})
	st.try("3:-1", []string{"c"})
	st.try("3:1", []string{})
	st.try("3:2", []string{})
	st.try("3:3", []string{"c"})
	st.try("3:4", []string{"c"})

	st.try("4:-4", []string{})
	st.try("4:-3", []string{})
	st.try("4:-2", []string{})
	st.try("4:-1", []string{})
	st.try("4:1", []string{})
	st.try("4:2", []string{})
	st.try("4:3", []string{})
	st.try("4:4", []string{})

	st.try(":1", []string{"a"})
	st.try(":2", []string{"a", "b"})
	st.try(":3", []string{"a", "b", "c"})
	st.try(":4", []string{"a", "b", "c"})
	st.try("1:", []string{"a", "b", "c"})
	st.try("2:", []string{"b", "c"})
	st.try("3:", []string{"c"})
	st.try("4:", []string{})

	st.try(":-1", []string{"a", "b", "c"})
	st.try(":-2", []string{"a", "b"})
	st.try(":-3", []string{"a"})
	st.try(":-4", []string{})
	st.try("-1:", []string{"c"})
	st.try("-2:", []string{"b", "c"})
	st.try("-3:", []string{"a", "b", "c"})
	st.try("-4:", []string{"a", "b", "c"})

	st.try(":", []string{"a", "b", "c"})
}

type SelectorTester struct {
	t      *testing.T
	fields []string
}

func newSelectorTester(t *testing.T, fields []string) *SelectorTester {
	return &SelectorTester{t, fields}
}

func (s *SelectorTester) try(expr string, expectedFields []string) {
	selector, err := ParseSelector(expr)
	if err != nil {
		s.t.Error(err)
		return
	}

	actualFields := selector.Select(s.fields)

	if !reflect.DeepEqual(actualFields, expectedFields) {
		s.t.Errorf("Expression '%s' selected from %v fields %v expecting %v", expr, s.fields, actualFields, expectedFields)
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
