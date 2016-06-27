package csv2htmltable

import (
	"bytes"
	"errors"
	"testing"

	json "github.com/mohae/unsafejson"
)

func TestWrite(t *testing.T) {
	tests := []struct {
		Caption      string
		Class        string
		ID           string
		Footer       string
		HeadingText  string
		HeadingType  int
		RowHeader    bool
		Section      bool
		TableHeader  bool
		HeaderRowNum int
		CSV          [][]string
		Expected     string
	}{
		{ // 0
			Class:        "",
			ID:           "",
			TableHeader:  false,
			HeaderRowNum: 0,
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
			Class:        "",
			ID:           "test",
			TableHeader:  false,
			HeaderRowNum: 0,
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
			Class:        "test",
			TableHeader:  false,
			HeaderRowNum: 0,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table class="test" border="">
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
		{ // 3
			Class:        "",
			ID:           "",
			Section:      true,
			TableHeader:  false,
			HeaderRowNum: 0,
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
			Class:        "",
			ID:           "",
			Section:      true,
			TableHeader:  false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingType:  5,
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
			Class:        "",
			ID:           "",
			TableHeader:  false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingType:  3,
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
			Class:        "",
			ID:           "",
			TableHeader:  false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
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
			Class:        "",
			ID:           "",
			TableHeader:  false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingType:  10,
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
			Class:        "people",
			TableHeader:  true,
			HeaderRowNum: 1,
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
			Caption:      "This is a test.",
			Class:        "people",
			TableHeader:  true,
			HeaderRowNum: 1,
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
			Caption:      "This is a test.",
			Class:        "people",
			Footer:       "This is a footer.",
			TableHeader:  true,
			HeaderRowNum: 1,
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
			Class:        "greetings",
			RowHeader:    true,
			TableHeader:  true,
			HeaderRowNum: 1,
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
		h.Reset()
		h.Caption = test.Caption
		h.Class = test.Class
		h.ID = test.ID
		h.Footer = test.Footer
		h.RowHeader = test.RowHeader
		h.TableHeader = test.TableHeader
		h.HeaderRowNum = test.HeaderRowNum
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
			t.Errorf("%d got %q; want %q", i, buf.String(), test.Expected)
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
	h.HeaderRowNum = 2
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
	if h.HeaderRowNum != 1 {
		t.Errorf("got %d, wanted 1", h.HeaderRowNum)
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

func TestIsNoDataErr(t *testing.T) {
	tests := []struct {
		err      error
		expected bool
	}{
		{err: errors.New("some error"), expected: false},
		{err: errNoData, expected: true},
	}
	for i, test := range tests {
		b := IsNoDataErr(test.err)
		if b != test.expected {
			t.Errorf("%d: got %t; want %t", i, b, test.expected)
		}
	}
}

func TestHeaderHandling(t *testing.T) {
	tests := []struct {
		TableHeader        bool
		HeaderRowNum       int
		RowHeader          bool
		HeaderRows         [][]string
		CSV                [][]string
		ExpectedHeaderRows [][]string
		ExpectedHTML       string
		ExpectedErr        string
	}{
		{ // 0
			TableHeader: true, HeaderRowNum: 1,
			ExpectedErr: "no table data found",
		},
		{ // 1
			TableHeader: true, HeaderRowNum: 0,
			ExpectedErr: "no table data found",
		},
		{ // 2
			TableHeader: true, HeaderRowNum: 0,
			CSV:         [][]string{[]string{"a"}, []string{"b"}},
			ExpectedErr: "no table header information found",
		},
		{ // 3
			TableHeader: true, HeaderRowNum: 1, RowHeader: true,
			HeaderRows: nil,
			CSV: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
			},
			ExpectedHTML: `
<table border="">
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
		{ // 4
			TableHeader: true, HeaderRowNum: 2, RowHeader: true,
			HeaderRows: nil,
			CSV: [][]string{
				[]string{"Language", "Greeting", "Title", "Name"},
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"Language", "Greeting", "Title", "Name"},
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
			},
			ExpectedHTML: `
<table border="">
    <thead>
        <th>Language</th>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
        <td>Langue</td>
        <td>Salutation</td>
        <td>Titre</td>
        <td>Prénom</td>
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
		{ // 5
			TableHeader: true, HeaderRowNum: 0, RowHeader: true,
			HeaderRows: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
			},
			CSV: [][]string{
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
			},
			ExpectedHTML: `
<table border="">
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
		{ // 6
			TableHeader: true, HeaderRowNum: 0, RowHeader: true,
			HeaderRows: [][]string{
				[]string{"Language", "Greeting", "Title", "Name"},
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
			},
			CSV: [][]string{
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"Language", "Greeting", "Title", "Name"},
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
			},
			ExpectedHTML: `
<table border="">
    <thead>
        <th>Language</th>
        <th>Greeting</th>
        <th>Title</th>
        <th>Name</th>
        <td>Langue</td>
        <td>Salutation</td>
        <td>Titre</td>
        <td>Prénom</td>
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
		{ // 7
			TableHeader: true, HeaderRowNum: 1, RowHeader: true,
			HeaderRows: [][]string{
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
			},
			CSV: [][]string{
				[]string{"", "Greeting", "Title", "Name"},
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
			},
			ExpectedHTML: `
<table border="">
    <thead>
        <th>Langue</th>
        <th>Salutation</th>
        <th>Titre</th>
        <th>Prénom</th>
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
		{ // 8
			TableHeader: true, HeaderRowNum: 2, RowHeader: true,
			HeaderRows: [][]string{
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
				[]string{"Idioma", "Saludo", "Título", "Nombre"},
			},
			CSV: [][]string{
				[]string{"Language", "Greeting", "Title", "Name"},
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
				[]string{"English", "Hello", "Mr.", "Bob"},
				[]string{"French", "Bonjour", "M.", "Genvieve"},
			},
			ExpectedHeaderRows: [][]string{
				[]string{"Langue", "Salutation", "Titre", "Prénom"},
				[]string{"Idioma", "Saludo", "Título", "Nombre"},
			},
			ExpectedHTML: `
<table border="">
    <thead>
        <th>Langue</th>
        <th>Salutation</th>
        <th>Titre</th>
        <th>Prénom</th>
        <td>Idioma</td>
        <td>Saludo</td>
        <td>Título</td>
        <td>Nombre</td>
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
		h.TableHeader = test.TableHeader
		h.HeaderRowNum = test.HeaderRowNum
		h.HeaderRows = test.HeaderRows
		h.RowHeader = test.RowHeader
		h.CSV = test.CSV
		err := h.Write(&buf)
		if err != nil {
			if test.ExpectedErr == "" {
				t.Errorf("%d: got %q: want nil", i, err)
			} else {
				if test.ExpectedErr != err.Error() {
					t.Errorf("%d: got %q want %q", i, err, test.ExpectedErr)
				}
			}
			continue
		}
		if buf.String() != test.ExpectedHTML {
			t.Errorf("%d got %q; want %q", i, buf.String(), test.ExpectedHTML)
			//			t.Errorf("%d got %s; want %s", i, buf.String(), test.Expected)
		}
		if json.MarshalToString(h.HeaderRows) != json.MarshalToString(test.ExpectedHeaderRows) {
			t.Errorf("%d: got %v; want %v", i, h.HeaderRows, test.ExpectedHeaderRows)
		}
	}
}
