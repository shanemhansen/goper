package goper

import (
	"bytes"
	"regexp"
	"strings"
)

var camelingRegex = regexp.MustCompile("[0-9A-Za-z]+")

// A Table represents the metadata associated with a table (columns, datatypes, etc)
type Table struct {
	Name    string
	Columns []Column
}

// A column has a name and a database type (hackily converted to a gotype)
type Column struct {
	Name   string
	DbType string
}

var typemap map[string]string = map[string]string{
	"int":       "*int64",
	"decimal":   "*float64", //fixme
	"varchar":   "*string",
	"text":      "*string",
	"float":     "*float64",
	"datetime":  "*string",
	"timestamp": "*string",
	"enum":      "*string",
	"date":      "*string",
	"double":    "float64",
	"char":      "*string",
	"bit":       "*int64",
	"longblob":  "*int64",
	"blob":      "[]byte",
	"BIGINT":    "*int64",
}

// Return the go type a database column should be mapped to.
// We always use pointers to handle null
func (this *Column) GoType() string {
	for key, value := range typemap {
		if strings.Contains(strings.ToLower(this.DbType), key) {
			return value
		}
	}
	panic("Unknown type" + this.DbType)
}

// A helper function that maps strings like product_variant to ProductVariant
func CamelCase(src string) string {
	byteSrc := []byte(src)
	chunks := camelingRegex.FindAll(byteSrc, -1)
	for idx, val := range chunks {
		chunks[idx] = bytes.Title(val)
	}
	return string(bytes.Join(chunks, nil))
}
