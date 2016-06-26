package csv2htmltable

import (
	"bytes"
	"errors"
	"testing"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		Caption     string
		Class       string
		ID          string
		Footer      string
		HeadingText string
		HeadingType int
		RowHeader   bool
		Section     bool
		TableHeader bool
		CSV         [][]string
		Expected    string
	}{
		{ // 0
			Class:       "",
			ID:          "",
			TableHeader: false,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 1
			Class:       "",
			ID:          "test",
			TableHeader: false,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table id="test" border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 2
			Class:       "test",
			TableHeader: false,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table class="test" border="">
    <tbody>
        <tr>
            <td>a</tr>
            <td>b</tr>
            <td>c</tr>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 3
			Class:       "",
			ID:          "",
			Section:     true,
			TableHeader: false,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section>
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
</section>
`,
		},
		{ // 4
			Class:       "",
			ID:          "",
			Section:     true,
			TableHeader: false,
			HeadingText: "Test Table",
			HeadingType: 5,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section>
<h5>Test Table</h5>
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
</section>
`,
		},
		{ // 5
			Class:       "",
			ID:          "",
			TableHeader: false,
			HeadingText: "Test Table",
			HeadingType: 3,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h3>Test Table</h3>
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 6
			Class:       "",
			ID:          "",
			TableHeader: false,
			HeadingText: "Test Table",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h4>Test Table</h4>
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 7
			Class:       "",
			ID:          "",
			TableHeader: false,
			HeadingText: "Test Table",
			HeadingType: 10,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h4>Test Table</h4>
<table border="">
    <tbody>
        <tr>
            <td>a</td>
            <td>b</td>
            <td>c</td>
        </tr>
        <tr>
            <td>1</td>
            <td>2</td>
            <td>3</td>
        </tr>
    </tbody>
</table>
`,
		},
		{ // 8
			Class: "people",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people" border="">
    <thead>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tbody>
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
    </tbody>
</table>
`,
		},
		{ // 9
			Caption: "This is a test.",
			Class:   "people",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people" border="">
    <caption>This is a test.</caption>
    <thead>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tbody>
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
    </tbody>
</table>
`,
		},
		{ // 10
			Caption: "This is a test.",
			Class:   "people",
			Footer:  "This is a footer.",
			CSV: [][]string{
				[]string{"Greeting", "Title", "Name"},
				[]string{"Hello", "Mr.", "Bob"},
				[]string{"Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="people" border="">
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
    <tbody>
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
    </tbody>
</table>
`,
		},
		{ // 11
			Class:     "greetings",
			RowHeader: true,
			CSV: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			Expected: `
<table class="greetings" border="">
    <thead>
        <th></th>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
    </thead>
    <tbody>
        <tr>
            <th>English</th>
            <td>Hello</td>
            <td>Mr.</td>
            <td>Bob</td>
        </tr>
        <tr>
            <th>French</th>
            <td>Bonjour</td>
            <td>M.</td>
            <td>Genvieve</td>
        </tr>
    </tbody>
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
		h.RowHeader = test.RowHeader
		h.TableHeader = test.TableHeader
		h.Section = test.Section
		h.HeadingText = test.HeadingText
		h.HeadingType = test.HeadingType
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

func TestReset(t *testing.T) {
	h := New("test")
	h.HeadingText = "heading"
	h.HeadingType = 3
	h.Border = "1"
	h.Caption = "caption"
	h.Class = "class"
	h.ID = "id"
	h.Footer = "footer"
	h.Cols = 4
	h.RowHeader = true
	h.TableHeader = false
	h.Section = true
	h.CSV = [][]string{[]string{"a", "b", "c"}}
	h.Reset()
	if h.HeadingText != "" {
		t.Errorf("got %q, wanted an empty string", h.HeadingText)
	}
	if h.HeadingType != 0 {
		t.Errorf("got %d, wanted 0", h.HeadingType)
	}
	if h.Border != "" {
		t.Errorf("got %q, wanted an empty string", h.Border)
	}
	if h.Caption != "" {
		t.Errorf("got %q, wanted an empty string", h.Caption)
	}
	if h.Class != "" {
		t.Errorf("got %q, wanted an empty string", h.Class)
	}
	if h.ID != "" {
		t.Errorf("got %q, wanted an empty string", h.ID)
	}
	if h.Footer != "" {
		t.Errorf("got %q, wanted an empty string", h.Footer)
	}
	if h.Cols != 0 {
		t.Errorf("got %d, wanted 0", h.Cols)
	}
	if h.RowHeader != false {
		t.Errorf("got %t, wanted false", h.RowHeader)
	}
	if h.TableHeader != true {
		t.Errorf("got %t, wanted true", h.TableHeader)
	}
	if h.Section != false {
		t.Errorf("got %t, wanted false", h.Section)
	}
	if len(h.CSV) != 0 {
		t.Errorf("CSV len was %d, wanted 0", len(h.CSV))
	}
}

func TestIsTableHeaderErr(t *testing.T) {
	tests := []struct {
		err      error
		expected bool
	}{
		{err: errors.New("some error"), expected: false},
		{err: errTableHeader, expected: true},
	}
	for i, test := range tests {
		b := IsTableHeaderErr(test.err)
		if b != test.expected {
			t.Errorf("%d: got %t; want %t", i, b, test.expected)
		}
	}
}
