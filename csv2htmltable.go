package csv2htmltable

import (
	"errors"
	"fmt"
	"html/template"
	"io"
)

// DefaultHTag is the default value for the Heading Element.
const DefaultHTag = "h4"

var errTableHeader = errors.New("no table header information found")
var errNoData = errors.New("no table data found")

var tableTpl = `
{{- $footer := .Footer }}
{{- $cols := .Cols}}
{{- $tableHeader := .TableHeader}}
{{- $rowHeader := .RowHeader}}
{{- if .Section}}
<section>
{{- end}}
{{- if .HeadingText}}
{{ htag .HeadingTag .HeadingText}}
{{- end}}
<table{{if .Class}} class="{{.Class}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} border="{{.Border}}">
{{- if .Caption}}
    <caption>{{.Caption}}</caption>
{{- end}}
{{- if $tableHeader }}
    <thead>
    {{- range $i, $row := .HeaderRows}}
        {{- range $fld := $row}}
            {{- if eq $i 0}}
        <th>{{$fld}}</th>
            {{- else}}
        <td>{{$fld}}</td>
            {{- end}}
        {{- end}}
    {{- end}}
    </thead>
{{- end}}
{{- if $footer}}
    <tfoot>
        <tr>
            <td colspan="{{$cols}}">{{$footer}}</td>
        </tr>
    </tfoot>
{{- end}}
{{- range $index, $record := .CSV -}}
    {{- if eq $index 0}}
    <tbody>
    {{- end}}
        <tr>
    {{- range $ndx, $field := $record}}
        {{- if eq $ndx 0}}
            {{- if $rowHeader}}
            <th>{{$field}}</th>
            {{- else}}
            <td>{{$field}}</td>
            {{- end}}
        {{- else}}
            <td>{{$field}}</td>
        {{- end}}
    {{- end}}
        </tr>
{{- end}}
    </tbody>
</table>
{{- if .Section}}
</section>
{{- end}}
`

// HTMLTable contains all of the information and configuration for a table
// and its data, including the template itself.  The generated table will be
// consistent with w3.org documentation for HTML5.  There may be elements
// missing or partially supported but any deprecated elements are not
// supported.  It is assumed that the styling of the table will be done
// with CSS.
//
// The border is always included and is either an empty string or "0" to
// indicate that this is not a table for layout purposes.
//
// The table may be preceded by a heading element with a custom string.  The
// heading element type is specified by setting HeadingType to an int between
// 1 and 6, inclusive.  If the value is invalid, it will be set to 4, the
// default, which corresponds with <h4>.
//
// Optionally, the table can be in its own section by setting Section to true.
// If the Class and ID fields are non-empty strings, they are applied to the
// <table> element.  An optional Caption and Footer is supported.  The
// Footer is assumed to span all of the columns in the table.
//
// The table's header rows output is controlled by the TableHeader field.
// When false, no table headers will be generated.  If the CSV data has
// record header rows, the HeaderRowNum should be set to the number of
// header rows it contains so that those rows will be skipped.  If true
// either the HeaderRowNum field can be set to the number of header rows
// that the CSV data starts with or the HeaderRows fields can be set.  If
// the CSV contains header row information and you want to use custom Header
// information, the HeaderRowNum should be set to the number of rows that the
// CSV data contains, so that the number of rows in the CSV data to skip is
// known, and the HeaderRows field should be set with the desired header
// information.
type HTMLTable struct {
	HeadingText string
	// The heading tag int, valid values are 1-6, invalid value are set to the default.
	HeadingTag   int
	Border       string // Should either be empty or 1.
	Caption      string
	Class        string
	ID           string
	Footer       string
	Cols         int
	RowHeader    bool // if true the first column of each row is a header
	Section      bool // Whether the table should be in its own section.
	TableHeader  bool // Whether the table has a header section.
	HeaderRowNum int  // Number of header rows in the CSV field.
	// Header information, if this is explicitly set and the CSV has header records,
	// the CSV header records will be ignored.
	HeaderRows [][]string
	CSV        [][]string
	tpl        *template.Template
}

// New returns a HTMLTable struct with a compiled table template whose name
// is set to the received value.  It is assumed that the CSV contains one
// header row and that it should be part of the generated table.  If that is
// not the case, the table header information must be explicitly set.
func New(n string) *HTMLTable {
	funcMap := template.FuncMap{
		"htag": Heading,
	}

	return &HTMLTable{TableHeader: true, HeaderRowNum: 1, tpl: template.Must(template.New(n).Funcs(funcMap).Parse(tableTpl))}
}

// Write accepts an io.Writer, validates the current configuration, and
// executes the HTML table template, writing the output to the received
// io.Writer.
func (h *HTMLTable) Write(w io.Writer) error {
	// Return an error if there's no table data.
	if len(h.CSV) == 0 {
		return errNoData
	}
	// If this is not empty, set it to 1, regardless of what it was set to.  This
	// is always set to explicitly indicate that this is a non-layout table. The
	// value must be either "" or "1".
	// See: https://www.w3.org/TR/html5/tabular-data.html#attr-table-border
	if h.Border != "" {
		h.Border = "1"
	}

	// If the CSV has header records, process them
	if h.HeaderRowNum > 0 {
		// if there weren't any custom headers set, copy the header records from
		// the CSV to the header rows
		if len(h.HeaderRows) == 0 {
			for i := 0; i < h.HeaderRowNum; i++ {
				h.HeaderRows = append(h.HeaderRows, h.CSV[i])
			}
		}
		// remove the header rows from the CSV
		h.CSV = h.CSV[h.HeaderRowNum:]
	}
	// If the table has headers; but there aren't any header rows: error.
	if h.TableHeader && len(h.HeaderRows) == 0 {
		return errTableHeader
	}
	h.Cols = len(h.CSV[0])
	return h.tpl.Execute(w, h)
}

// Heading returns the heading element as template.HTML.  If the HeadingType
// is < 0 || > 6, the DefaultHTag will be used.
func Heading(i int, s string) template.HTML {
	var htag string
	switch i {
	case 1:
		htag = "h1"
	case 2:
		htag = "h2"
	case 3:
		htag = "h3"
	case 4:
		htag = "h4"
	case 5:
		htag = "h5"
	case 6:
		htag = "h6"
	default:
		htag = DefaultHTag
	}
	return template.HTML(fmt.Sprintf("<%s>%s</%s>", htag, s, htag))
}

// Reset resets all of the structs settings to their defaults
func (h *HTMLTable) Reset() {
	h.HeadingText = ""
	h.HeadingTag = 0
	h.Border = ""
	h.Caption = ""
	h.Class = ""
	h.ID = ""
	h.Footer = ""
	h.Cols = 0
	h.RowHeader = false
	h.Section = false
	h.TableHeader = true
	h.HeaderRowNum = 1
	h.HeaderRows = h.HeaderRows[:0]
	h.CSV = h.CSV[:0]
}

// IsTableHeaderErr returns whether or not the error returned was a result of
// an error in the Table Header.
func IsTableHeaderErr(err error) bool {
	return err.Error() == errTableHeader.Error()
}

// IsNoDataErr returns whether or not the error was a result of no table data
// being present.
func IsNoDataErr(err error) bool {
	return err.Error() == errNoData.Error()
}
