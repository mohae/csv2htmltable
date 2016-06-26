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

var tableTpl = `
{{- $footer := .Footer }}
{{- $cols := .Cols}}
{{- $rowHeader := .RowHeader}}
{{- if .Section}}
<section>
{{- end}}
{{- if .HeadingText}}
{{ htag .HeadingType .HeadingText}}
{{- end}}
<table{{if .Class}} class="{{.Class}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}} border="{{.Border}}">
    {{- if .Caption}}
    <caption>{{.Caption}}</caption>{{end}}
    {{- range $index, $record := .CSV -}}
        {{- if eq $index 0}}
    <thead>
            {{- range $record}}
        <th>{{.}}</th>{{end}}
    </thead>
            {{- if $footer}}
    <tfoot>
        <tr>
            <td colspan="{{$cols}}">{{$footer}}</td>
        </tr>
    </tfoot>
            {{- end}}
        {{- else}}
            {{- if eq $index 1}}
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
    {{- end}}
    </tbody>
</table>
{{- if .Section}}
</section>
{{- end}}
`

type HTMLTable struct {
	HeadingText string
	// The heading tag int, valid values are 1-6, invalid value are set to the default.
	HeadingType int
	Border      string // Should either be empty or 1.
	Caption     string
	Class       string
	ID          string
	Footer      string
	Cols        int
	RowHeader   bool // if true the first column of each row is a header
	Section     bool // Whether the table should be in its own section.
	TableHeader bool // Whether the table has a header section.
	CSV         [][]string
	tpl         *template.Template
}

func New(n string) *HTMLTable {
	funcMap := template.FuncMap{
		"htag": Heading,
	}

	return &HTMLTable{TableHeader: true, tpl: template.Must(template.New(n).Funcs(funcMap).Parse(tableTpl))}
}

func (h *HTMLTable) Write(w io.Writer) error {
	// If this is not empty, set it to 1, regardless of what it was set to.  This
	// is always set to explicitly indicate that this is a non-layout table. The
	// value must be either "" or "1".
	// See: https://www.w3.org/TR/html5/tabular-data.html#attr-table-border
	if h.Border != "" {
		h.Border = "1"
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
	h.HeadingType = 0
	h.Border = ""
	h.Caption = ""
	h.Class = ""
	h.ID = ""
	h.Footer = ""
	h.Cols = 0
	h.RowHeader = false
	h.Section = false
	h.TableHeader = true
	h.CSV = h.CSV[:0]
}

// IsTableHeaderError
func IsTableHeaderErr(err error) bool {
	return err.Error() == errTableHeader.Error()
}
