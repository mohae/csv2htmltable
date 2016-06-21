package csv2htmltable

import (
	"html/template"
	"io"
)

var tableTpl = `
{{- $footer := .Footer }}
{{- $cols := .Cols}}
<table{{if .Class}} class="{{.Class}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}}>
    {{- if .Caption}}
    <caption>{{.Caption}}</caption>{{end}}
    {{- range $index, $row := .CSV -}}
        {{- if eq $index 0}}
    <thead>
        {{- range $row}}
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
        {{- range $row}}
        <td>{{.}}</td>{{end}}
    </tr>
        {{- end}}
    {{- end}}
</table>
`

type HTMLTable struct {
	Caption string
	Class   string
	ID      string
	Footer  string
	Cols    int
	CSV     [][]string
	tpl     *template.Template
}

func New(n string) *HTMLTable {
	return &HTMLTable{tpl: template.Must(template.New(n).Parse(tableTpl))}
}

func (h *HTMLTable) Write(w io.Writer) error {
	h.Cols = len(h.CSV[0])
	return h.tpl.Execute(w, h)
}
