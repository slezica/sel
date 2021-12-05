package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cli := ParseCli()

	doTheThing(cli.splitter, cli.joiner, cli.selectors)
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
