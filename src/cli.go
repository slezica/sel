package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Cli struct {
	splitter  *Splitter
	joiner    *Joiner
	selectors []*Selector
}

const dashPlaceholderSymbol = "௸"

func ParseCli2(args []string) *Cli {
	flagSet := flag.NewFlagSet("sel", flag.ExitOnError)

	splitExpr := flagSet.String("split", "\\s+", "Regex to split fields")
	joinExpr := flagSet.String("join", " ", "String to join selected fields")
	help := flagSet.Bool("help", false, "Show help")

	transformArgsBeforeParse(args)

	flagSet.Parse(args)
	selectorArgs := flagSet.Args()

	// transformArgsAfterParse(selectorArgs)

	if *help {
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	splitter, err := ParseSplitter(*splitExpr)
	check(err, "Invalid regex to split fields: %s", *splitExpr)

	joiner, err := ParseJoiner(*joinExpr)
	check(err, "Invalid string to join fields: %s", *joinExpr)

	var selectors []*Selector

	for _, selectorExpr := range selectorArgs {
		selector, err := ParseSelector(selectorExpr)
		check(err, "Invalid selector: '%s'", selectorExpr)

		selectors = append(selectors, selector)
	}

	return &Cli{splitter, joiner, selectors}
}

func ParseCli(args []string) *Cli {
	flagSet := flag.NewFlagSet("sel", flag.ExitOnError)

	splitExpr := flagSet.String("split", "\\s+", "Regex to split fields")
	joinExpr := flagSet.String("join", " ", "String to join selected fields")
	help := flagSet.Bool("help", false, "Show help")

	// Before we call `flagSet.Parse`, we need to separate selectors from flags ourselves. This is because negative
	// selectors (such as "-1") look like invalid flags to the implementation, and there's no option to avoid that:
	flags, selectorExprs := classifyArgs(args)
	flagSet.Parse(flags)

	if *help {
		flagSet.PrintDefaults()
		os.Exit(1)
	}

	splitter, err := ParseSplitter(*splitExpr)
	check(err, "Invalid regex to split fields: %s", *splitExpr)

	joiner, err := ParseJoiner(*joinExpr)
	check(err, "Invalid string to join fields: %s", *joinExpr)

	var selectors []*Selector

	for _, selectorExpr := range selectorExprs {
		selector, err := ParseSelector(selectorExpr)
		check(err, "Invalid selector: '%s'", selectorExpr)

		selectors = append(selectors, selector)
	}

	return &Cli{splitter, joiner, selectors}
}

// classifyArgs separates flags and positional arguments, because negative selectors (eg "-1") look like invalid flags
func classifyArgs(args []string) ([]string, []string) {
	var flags []string
	var selectorExprs []string

	for _, arg := range args {
		// If this argument is obviously not a flag, just add it to the selector expression list:
		if !strings.HasPrefix(arg, "-") {
			selectorExprs = append(selectorExprs, arg)
			continue
		}

		// Now this could be a flag or a negative selector. Try to parse it, see what happens:
		_, err := ParseSelector(arg)

		if err != nil {
			flags = append(flags, arg)
		} else {
			selectorExprs = append(selectorExprs, arg)
		}

		// NOTE:
		// For the negative selector vs flag case, we're running `ParseSelector` twice (will also happen in `main`). Not
		// ideal, but not a problem that merits additional complexity. I blame the `flags` module for forcing me into this
		// position. Yeah, that's it. Not my fault.
	}

	return flags, selectorExprs
}

func transformArgsBeforeParse(args []string) {
	negativeSelectorRe := regexp.MustCompile("^-([0-9:]+)$")

	for i, arg := range args {
		args[i] = negativeSelectorRe.ReplaceAllString(arg, dashPlaceholderSymbol+"$1")
	}
}

func transformArgsAfterParse(args []string) {
	for i, arg := range args {
		args[i] = strings.ReplaceAll(arg, dashPlaceholderSymbol, "-")
	}
}

func check(err error, message string, params ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", fmt.Sprintf(message, params...))
		os.Exit(1)
	}
}
