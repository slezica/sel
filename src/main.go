package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	flagSet := flag.NewFlagSet("main", flag.ExitOnError)

	splitExpr := flagSet.String("split", "\\s+", "Regex to split fields")
	joinExpr := flagSet.String("join", " ", "String to join selected fields")
	help := flagSet.Bool("help", false, "Show help")

	flags, selectorExprs := classifyArgs(os.Args[1:])
	flagSet.Parse(flags)

	if *help {
		flagSet.PrintDefaults()
		os.Exit(0)
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

	doTheThing(splitter, joiner, selectors)
}

// classifyArgs separates flags and positional arguments, because negative selectors (eg "-1") look like invalid flags
func classifyArgs(args []string) ([]string, []string) {
	var flags []string
	var selectorExprs []string

	for _, arg := range args {
		if !strings.HasPrefix(arg, "-") {
			selectorExprs = append(flags, arg)
			continue
		}

		// Between this and main(), we're parsing selectors twice. Not optimal, but no big deal.
		_, err := ParseSelector(arg)

		if err != nil {
			flags = append(flags, arg)
		} else {
			selectorExprs = append(selectorExprs, arg)
		}
	}

	return flags, selectorExprs
}

func doTheThing(splitter *Splitter, joiner *Joiner, selectors []*Selector) {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := splitter.Split(line)

		var selectedFields []string

		for _, selector := range selectors {
			selectedFields = append(selectedFields, selector.Select(fields)...)
		}

		fmt.Println(joiner.Join(selectedFields))
	}
}

func check(err error, message string, params ...interface{}) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", fmt.Sprintf(message, params...))
		os.Exit(1)
	}
}
