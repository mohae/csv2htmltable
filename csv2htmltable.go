package csv2htmltable

import (
	"html/template"
	"io"
)

var tableTpl = `
{{- $footer := .Footer }}
{{- $cols := .Cols}}
{{- $rowHeader := .RowHeader}}
{{- if .Section}}
<section>
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
	return &HTMLTable{TableHeader: true, tpl: template.Must(template.New(n).Parse(tableTpl))}
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
