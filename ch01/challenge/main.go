package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path"
)

/*
Instructions:

- Pass text to print as an argument
  - If not argument - read from stdin

- Use -width to specify width
  - Width should be bigger than 0 and less than 250
  - Default to 80

- Use -out to specify output file
  - Default to stdout
*/

var config struct {
	width int
	out   string
}

var usage = `Usage: %s [options] [TEXT]
Prints a banner. For example:
$ banner -w 6 Go
  Go
======
Options:
`

func main() {
	config.width = 80

	flag.Usage = func() {
		name := path.Base(os.Args[0])
		fmt.Fprintf(os.Stderr, usage, name)
		flag.PrintDefaults()
	}

	flag.Var(&Width{&config.width}, "width", "banner width [0 - 250]")
	flag.StringVar(&config.out, "out", "", "output file (default stdout)")
	flag.Parse()

	text := ""

	switch flag.NArg() {
	case 0:
		data, err := io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: can't read - %s\n", err)
			os.Exit(1)
		}
		text = string(data)
	case 1:
		text = flag.Arg(0)
	default:
		fmt.Fprintln(os.Stderr, "error: wrong number of arguments")
		os.Exit(1)
	}

	out := os.Stdout

	if config.out != "" {
		file, err := os.Create(config.out)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error: %s\n", err)
			os.Exit(1)
		}
		defer file.Close()
		out = file
	}

	Banner(out, text, config.width)
}
