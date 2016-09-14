package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/mohae/csv2htmltable"
)

var (
	input  string
	output string

	section     bool
	headingText string
	headingTag  int

	class        string
	id           string
	caption      string
	tableHeader  bool
	headerRowNum int
	rowHeader    bool
	footer       string
)

func init() {
	// c d f h i n o p r s x
	flag.StringVar(&input, "input", "stdin", "the path to the input file; if not specified stdin is used")
	flag.StringVar(&input, "i", "stdin", "the short flag for -input")
	flag.StringVar(&output, "output", "stdout", "output destination")
	flag.StringVar(&output, "o", "stdout", "output destination (short)")

	flag.BoolVar(&section, "section", false, "create the table in its own section")
	flag.BoolVar(&section, "s", false, "create the table in its own section")
	flag.StringVar(&headingText, "headingtext", "", "text for the heading element")
	flag.StringVar(&headingText, "x", "", "text for the heading element")
	flag.IntVar(&headingTag, "headingtag", 4, "int representing the heading tag size: valid values are 1-6")
	flag.IntVar(&headingTag, "t", 4, "int representing the heading tag size: valid values are 1-6")

	flag.StringVar(&class, "class", "", "the table's class")
	flag.StringVar(&class, "c", "", "the table's class")
	flag.StringVar(&id, "id", "", "the table's id")
	flag.StringVar(&id, "d", "", "the table's id")
	flag.StringVar(&caption, "caption", "", "table caption")
	flag.StringVar(&caption, "p", "", "table caption")
	flag.StringVar(&footer, "footer", "", "table footer")
	flag.StringVar(&footer, "f", "", "table footer")
	flag.BoolVar(&tableHeader, "tableheader", true, "include table heading in the output")
	flag.BoolVar(&tableHeader, "h", true, "include table heading in the output")
	flag.IntVar(&headerRowNum, "headerrownum", 1, "number of header rows in the csv")
	flag.IntVar(&headerRowNum, "n", 1, "number of header rows in the csv")
	flag.BoolVar(&rowHeader, "rowheader", false, "make the first column of each row a header")
	flag.BoolVar(&rowHeader, "r", false, "make the first column of each row a header")
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
	htable.HasRowHeader = rowHeader
	htable.Section.Include = section
	if headingText != "" {
		htable.HeadingText = headingText
	}
	htable.HeadingTag = headingTag
	htable.HasHeader = tableHeader
	htable.HeaderRowNum = headerRowNum
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
