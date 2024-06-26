package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/janobono/writing-command-line-applications-in-go/logs"
)

var config struct {
	count   int
	verbose bool
}

func main() {
	flag.IntVar(&config.count, "count", 0, "number of records to parse")
	flag.BoolVar(&config.verbose, "verbose", false, "emit more information")
	flag.Parse()

	if config.count < 0 {
		fmt.Fprintf(os.Stderr, "error: count should be positive, got %d\n", config.count)
		os.Exit(1)
	}

	var r io.Reader
	var fileName string
	switch flag.NArg() {
	case 0:
		r, fileName = os.Stdin, "<stdin>"
	case 1:
		fileName = flag.Arg(0)
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()

		r = file
	default:
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	s := logs.NewScanner(r)
	n := 0
	for s.Next() {
		n++
		if config.count > 0 && config.count == n {
			break
		}
	}

	if err := s.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %s:%d: scanning - %s\n", fileName, s.Line(), err)
		os.Exit(1)
	}

	if config.verbose {
		fmt.Printf("%s: successfully validated %d records\n", fileName, n)
	}
}
