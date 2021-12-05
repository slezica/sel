package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	cli := ParseCli()

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		fields := cli.splitter.Split(line)

		var selectedFields []string

		for _, selector := range cli.selectors {
			selectedFields = append(selectedFields, selector.Select(fields)...)
		}

		fmt.Println(cli.joiner.Join(selectedFields))
	}
}
