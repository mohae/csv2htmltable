package csv2htmltable

import (
	"bytes"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		Class    string
		CSV      [][]string
		Expected string
	}{
		{
			Class: "",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table>
    <tr>
        <td>a</td><td>b</td><td>c</td>
    </tr>
    <tr>
        <td>1</td><td>2</td><td>3</td>
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
    <tr>
        <td>a</td><td>b</td><td>c</td>
    </tr>
    <tr>
        <td>1</td><td>2</td><td>3</td>
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
    <tr>
        <td>Greeting</td><td>Title</td><td>Name</td>
    </tr>
    <tr>
        <td>Hello</td><td>Mr.</td><td>Bob</td>
    </tr>
    <tr>
        <td>Bonjour</td><td>M.</td><td>Genvieve</td>
    </tr>
</table>
`,
		},
	}
	var buf bytes.Buffer
	h := New("test")
	for i, test := range tests {
		buf.Reset()
		h.Class = test.Class
		h.CSV = test.CSV
		err := h.Write(&buf)
		if err != nil {
			t.Errorf("%d: got %q: want nil", i, err)
			continue
		}
		if buf.String() != test.Expected {
			t.Errorf("%d got %s; want %s", i, buf.String(), test.Expected)
		}
	}
}
