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
		Section      bool
		HeadingTag   int
		HasRowHeader bool
		HasHeader    bool
		HeaderRowNum int
		CSV          [][]string
		Expected     string
	}{
		{ // 0
			Class:        "",
			ID:           "",
			HasHeader:    false,
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
		{ // 1
			Class:        "",
			ID:           "test",
			HasHeader:    false,
			HeaderRowNum: 0,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table class="test" id="test" border="">
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
			HasHeader:    false,
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
			HasHeader:    false,
			HeaderRowNum: 0,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section>
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
</section>
`,
		},
		{ // 4
			Class:        "",
			ID:           "",
			Section:      true,
			HasHeader:    false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingTag:   5,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section>
<h5>Test Table</h5>
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
</section>
`,
		},
		{ // 5
			Class:        "",
			ID:           "",
			HasHeader:    false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingTag:   3,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h3>Test Table</h3>
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
		{ // 6
			Class:        "",
			ID:           "",
			HasHeader:    false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h4>Test Table</h4>
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
		{ // 7
			Class:        "",
			ID:           "",
			HasHeader:    false,
			HeaderRowNum: 0,
			HeadingText:  "Test Table",
			HeadingTag:   10,
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<h4>Test Table</h4>
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
		{ // 8
			Class:        "people",
			HasHeader:    true,
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
			HasHeader:    true,
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
			HasHeader:    true,
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
			HasRowHeader: true,
			HasHeader:    true,
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
		if test.Class != "" {
			h.Class = test.Class
		}
		h.ID = test.ID
		h.Footer = test.Footer
		h.HasRowHeader = test.HasRowHeader
		h.HasHeader = test.HasHeader
		h.HeaderRowNum = test.HeaderRowNum
		h.Section.Include = test.Section
		h.HeadingText = test.HeadingText
		h.HeadingTag = test.HeadingTag
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
	h.HeadingTag = 3
	h.Border = "1"
	h.Caption = "caption"
	h.Class = "class"
	h.ID = "id"
	h.Footer = "footer"
	h.Cols = 4
	h.HasRowHeader = true
	h.HasHeader = false
	h.HeaderRowNum = 2
	h.Section.Include = true
	h.Section.Class = "abc"
	h.Section.ID = "123"
	h.CSV = [][]string{[]string{"a", "b", "c"}}
	h.Reset()
	if h.HeadingText != "" {
		t.Errorf("got %q, wanted an empty string", h.HeadingText)
	}
	if h.HeadingTag != 0 {
		t.Errorf("got %d, wanted 0", h.HeadingTag)
	}
	if h.Border != "" {
		t.Errorf("got %q, wanted an empty string", h.Border)
	}
	if h.Caption != "" {
		t.Errorf("got %q, wanted an empty string", h.Caption)
	}
	if h.Class != "test" {
		t.Errorf("got %q, wanted \"test\"", h.Class)
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
	if h.HasRowHeader != false {
		t.Errorf("got %t, wanted false", h.HasRowHeader)
	}
	if h.HasHeader != true {
		t.Errorf("got %t, wanted true", h.HasHeader)
	}
	if h.HeaderRowNum != 1 {
		t.Errorf("got %d, wanted 1", h.HeaderRowNum)
	}
	if h.Section.Include != false {
		t.Errorf("got %t, wanted false", h.Section.Include)
	}
	if h.Section.Class != "" {
		t.Errorf("got %q, wanted empty string", h.Section.Class)
	}
	if h.Section.ID != "" {
		t.Errorf("got %q, wanted empty string", h.Section.ID)
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
		HasHeader          bool
		HeaderRowNum       int
		HasRowHeader       bool
		HeaderRows         [][]string
		CSV                [][]string
		ExpectedHeaderRows [][]string
		ExpectedHTML       string
		ExpectedErr        string
	}{
		{ // 0
			HasHeader: true, HeaderRowNum: 1,
			ExpectedErr: "no table data found",
		},
		{ // 1
			HasHeader: true, HeaderRowNum: 0,
			ExpectedErr: "no table data found",
		},
		{ // 2
			HasHeader: true, HeaderRowNum: 0,
			CSV:         [][]string{[]string{"a"}, []string{"b"}},
			ExpectedErr: "no table header information found",
		},
		{ // 3
			HasHeader: true, HeaderRowNum: 1, HasRowHeader: true,
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
<table class="test" border="">
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
			HasHeader: true, HeaderRowNum: 2, HasRowHeader: true,
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
<table class="test" border="">
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
			HasHeader: true, HeaderRowNum: 0, HasRowHeader: true,
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
<table class="test" border="">
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
			HasHeader: true, HeaderRowNum: 0, HasRowHeader: true,
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
<table class="test" border="">
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
			HasHeader: true, HeaderRowNum: 1, HasRowHeader: true,
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
<table class="test" border="">
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
			HasHeader: true, HeaderRowNum: 2, HasRowHeader: true,
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
<table class="test" border="">
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
		h.HasHeader = test.HasHeader
		h.HeaderRowNum = test.HeaderRowNum
		h.HeaderRows = test.HeaderRows
		h.HasRowHeader = test.HasRowHeader
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

func TestSection(t *testing.T) {
	tests := []struct {
		Section      bool
		SectionID    string
		SectionClass string
		CSV          [][]string
		Expected     string
	}{
		{ // 0
			Section:      false,
			SectionID:    "",
			SectionClass: "",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<table class="test" border="">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tbody>
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
			Section:      true,
			SectionID:    "sid",
			SectionClass: "sclass",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section class="sclass" id="sid">
<table class="test" border="">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tbody>
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
		{ // 2
			Section:      true,
			SectionID:    "",
			SectionClass: "",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section>
<table class="test" border="">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tbody>
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
		{ // 3
			Section:      true,
			SectionID:    "",
			SectionClass: "sclass",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section class="sclass">
<table class="test" border="">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tbody>
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
			Section:      true,
			SectionID:    "sid",
			SectionClass: "",
			CSV: [][]string{
				[]string{"a", "b", "c"},
				[]string{"1", "2", "3"},
			},
			Expected: `
<section id="sid">
<table class="test" border="">
    <thead>
        <th>a</th>
        <th>b</th>
        <th>c</th>
    </thead>
    <tbody>
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
	}
	var buf bytes.Buffer
	h := New("test")
	for i, test := range tests {
		buf.Reset()
		h.Section.Include = test.Section
		h.Section.Class = test.SectionClass
		h.Section.ID = test.SectionID
		h.CSV = test.CSV
		err := h.Write(&buf)
		if err != nil {
			t.Errorf("%d: got %q: want nil", i, err)
			continue
		}
		if buf.String() != test.Expected {
			t.Errorf("%d got %q; want %q", i, buf.String(), test.Expected)
			//			t.Errorf("%d got %s; want %s", i, buf.String(), test.Expected)
		}
	}
}
