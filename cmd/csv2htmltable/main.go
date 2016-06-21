package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/mohae/csv2htmltable"
)

var (
	input   string
	output  string
	class   string
	id      string
	caption string
	footer  string
)

func init() {
	flag.StringVar(&input, "input", "stdin", "the path to the input file; if not specified stdin is used")
	flag.StringVar(&input, "i", "stdin", "the short flag for -input")
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")
	flag.StringVar(&class, "class", "", "the table's class")
	flag.StringVar(&class, "c", "", "the table's class")
	flag.StringVar(&id, "id", "", "the table's id")
	flag.StringVar(&id, "d", "", "the table's id")
	flag.StringVar(&caption, "caption", "", "table caption")
	flag.StringVar(&caption, "p", "", "table caption")
	flag.StringVar(&footer, "footer", "", "table footer")
	flag.StringVar(&footer, "f", "", "table footer")
}

func main() {
	os.Exit(realMain())
}

func realMain() int {
	flag.Parse()

	var err error
	// by default set to stdin and stdout
	in := os.Stdin
	out := os.Stdout

	// If input was set, use that.
	if input != "stdin" {
		in, err = os.Open(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening input: %s\n", err)
			return 1
		}
		defer in.Close()
	}

	// If output was set, use that.
	if output != "stdout" {
		out, err = os.Open(output)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening output: %s\n", err)
			return 1
		}
		defer out.Close()
	}

	htable := csv2htmltable.New("htmltable")
	if caption != "" {
		htable.Caption = caption
	}
	if class != "" {
		htable.Class = class
	}
	if id != "" {
		htable.ID = id
	}
	if footer != "" {
		htable.Footer = footer
	}

	r := csv.NewReader(in)
	htable.CSV, err = r.ReadAll()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading CSV: %s\n", err)
		return 1
	}
	err = htable.Write(out)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing HTML table: %s\n", err)
		return 1
	}
	return 0
}
