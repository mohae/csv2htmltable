package csv2htmltable

import (
	"html/template"
	"io"
)

var tableTpl = `
<table{{if .Class}} class="{{.Class}}"{{end}}{{if .ID}} id="{{.ID}}"{{end}}>
{{range $index, $row := .CSV}}{{if eq $index 0}}    <thead>
{{range $row}}        <th>{{.}}</th>
{{end}}    </thead>
{{- else}}    <tr>
{{range $row}}        <td>{{.}}</td>
{{end}}    </tr>
{{- end}}
{{end}}</table>
`

type HTMLTable struct {
	Class string
	ID    string
	Title string
	CSV   [][]string
	tpl   *template.Template
}

func New(n string) *HTMLTable {
	return &HTMLTable{tpl: template.Must(template.New(n).Parse(tableTpl))}
}

func (h *HTMLTable) Write(w io.Writer) error {
	return h.tpl.Execute(w, h)
}
