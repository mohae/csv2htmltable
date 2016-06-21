package csv2htmltable

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		Caption  string
		Class    string
		ID       string
		Footer   string
		CSV      [][]string
		Expected string
	}{
		{
			Class: "",
			ID:    "",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table>
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tr>
        <td>1</td>
        <td>2</td>
        <td>3</td>
    </tr>
</table>
`,
		},
		{
			Class: "",
			ID:    "test",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table id="test">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tr>
        <td>1</td>
        <td>2</td>
        <td>3</td>
    </tr>
</table>
`,
		},
		{
			Class: "test",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table class="test">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tr>
        <td>1</td>
        <td>2</td>
        <td>3</td>
    </tr>
</table>
`,
		},
		{
			Class: "people",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people">
    <thead>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tr>
        <td>Hello</td>
        <td>Mr.</td>
        <td>Bob</td>
    </tr>
    <tr>
        <td>Bonjour</td>
        <td>M.</td>
        <td>Genvieve</td>
    </tr>
</table>
`,
		},
		{
			Caption: "This is a test.",
			Class:   "people",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people">
    <caption>This is a test.</caption>
    <thead>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tr>
        <td>Hello</td>
        <td>Mr.</td>
        <td>Bob</td>
    </tr>
    <tr>
        <td>Bonjour</td>
        <td>M.</td>
        <td>Genvieve</td>
    </tr>
</table>
`,
		},
		{
			Caption: "This is a test.",
			Class:   "people",
			Footer:  "This is a footer.",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people">
    <caption>This is a test.</caption>
    <thead>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tfoot>
        <tr>
            <td colspan="3">This is a footer.</td>
        </tr>
    </tfoot>
    <tr>
        <td>Hello</td>
        <td>Mr.</td>
        <td>Bob</td>
    </tr>
    <tr>
        <td>Bonjour</td>
        <td>M.</td>
        <td>Genvieve</td>
    </tr>
</table>
`,
		},
	}
	var buf bytes.Buffer
	h := New("test")
	for i, test := range tests {
		buf.Reset()
		h.Caption = test.Caption
		h.Class = test.Class
		h.ID = test.ID
		h.Footer = test.Footer
		h.CSV = test.CSV
		err := h.Write(&buf)
		if err != nil {
			t.Errorf("%d: got %q: want nil", i, err)
			continue
		}
		if buf.String() != test.Expected {
			//t.Errorf("%d got %q; want %q", i, buf.String(), test.Expected)
			t.Errorf("%d got %s; want %s", i, buf.String(), test.Expected)
		}
	}
}
