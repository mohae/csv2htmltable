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
<table{{if .Class}} class="{{.Class}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}}>
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
</table>
{{- if .Section}}
</section>
{{- end}}
`

type HTMLTable struct {
	Caption   string
	Class     string
	ID        string
	Footer    string
	Cols      int
	RowHeader bool // if true the first column of each row is a header
	Section      bool // Whether the table should be in its own section.
	CSV       [][]string
	tpl       *template.Template
}

func New(n string) *HTMLTable {
	return &HTMLTable{tpl: template.Must(template.New(n).Parse(tableTpl))}
}

func (h *HTMLTable) Write(w io.Writer) error {
	h.Cols = len(h.CSV[0])
	return h.tpl.Execute(w, h)
}
