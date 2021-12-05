package main

import (
	"math"
	"testing"
)

func TestAll(t *testing.T) {
	cli := ParseCli([]string{"-join", "a", "-split", "\\t", "3"})

	if cli.joiner.Delim != "a" {
		t.Fatalf("expected Joiner.Delim 'a', not '%s'", cli.joiner.Delim)
	}

	if cli.splitter.re.String() != "\\t" {
		t.Fatalf("expected Splitter.Re '\\t', not '%s'", cli.splitter.re.String())
	}

	if len(cli.selectors) != 3 {
		t.Fatalf("expected 3 selectors, not %d", len(cli.selectors))
	}
}

func TestDefaults(t *testing.T) {
	cli := ParseCli([]string{})

	if cli.joiner.Delim != " " {
		t.Fatalf("expected Joiner.Delim ' ', not '%s'", cli.joiner.Delim)
	}

	if cli.splitter.re.String() != "\\s+" {
		t.Fatalf("expected Splitter.Re '\\s+', not '%s'", cli.splitter.re.String())
	}
}

func TestSplitFlag(t *testing.T) {
	expectRe := "x12*"

	cli := ParseCli([]string{"-split=" + expectRe})
	actualRe := cli.splitter.re.String()

	if actualRe != expectRe {
		t.Fatalf("expected Splitter.Re '%s', not '%s'", expectRe, actualRe)
	}
}

func TestJoinFlag(t *testing.T) {
	expectDelim := "x12*"

	cli := ParseCli([]string{"-join=" + expectDelim})
	actualDelim := cli.joiner.Delim

	if actualDelim != expectDelim {
		t.Fatalf("expected Joiner.Delim '%s', not '%s'", expectDelim, actualDelim)
	}
}

func TestSelectors(t *testing.T) {
	selectorExprs := []string{"1", ":2", "3:", "4:5", "-1", ":-2"}
	expectedRanges := [][]int{{1, 1}, {1, 2}, {3, math.MaxInt32}, {4, 5}, {-1, -1}, {1, -2}}

	cli := ParseCli2(selectorExprs)

	for i, selector := range cli.selectors {
		if selector.start != expectedRanges[i][0] || selector.end != expectedRanges[i][1] {
			t.Fatalf("Selector %d expected range %v, got %+v", i, selector, expectedRanges[i])
		}
	}
}
