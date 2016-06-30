# csv2htmltable
[![GoDoc](https://godoc.org/github.com/mohae/csv2htmltable?status.svg)](https://godoc.org/github.com/mohae/csv2htmltable)[![Build Status](https://travis-ci.org/mohae/csv2htmltable.png)](https://travis-ci.org/mohae/csv2htmltable)

Generates a HTML Table from CSV and writes the output to the provided `io.Writer`; only table elements that are not deprecated by HTML5 are supported.

The table can be in its own section, with its own Heading element.  The following table elements are supported: `caption`, `class`, `id`, `footer`, and `thead`.  Also, the first column of each row can be a `th` element.

By default, the value used for the template name will be used for the table's class.  This can be overridden by setting the `Class` field.

### Table Header(s)
For CSV data that contain multiple rows of header information, the number of rows can be set.  The table headers can also be set by explicitly setting the `HeaderRows` field.  If the CSV data contains header information, but that information is to be overridden the `HeaderRows` can be set and the `HeaderRowNum` field should be set to the appropriate value.  If the CSV data does not contain any header information, the `HeaderRowNum` should be set to `0`; it's default is `1`.

## TODO:
* Revisit the handling of sections and headers.
* Possibly support adding html between a section header and the table.
* Add optional div element (should div and section be mutually exclusive)? (probably)
* Add lang element?
* Add title element?
