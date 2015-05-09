package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/docopt/docopt-go"
)

const (
	version = "1.0.0"
	usage   = `Grip.

Filters input matching given regex.

Usage:
  grip <regex>
  grip -h | --help
  grip --version

Options:
  -h --help     Show this screen.
  --version     Show version.
`
)

func main() {
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		log.Fatal(err)
	}

	regex := args["<regex>"].(string)
	pattern := regexp.MustCompile(regex)

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		if matched := pattern.MatchString(text); matched {
			fmt.Printf("%v\n", text)
		}
	}
}
