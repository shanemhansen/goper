package goper

import "io"
import "fmt"
import "strings"
import "database/sql"

// A SchemaWriter writes a set of tables to the writer denoted by Outfile
type SchemaWriter struct {
	PackageName string
	Outfile     io.Writer
	Tables      []*Table
}

// Write the schema
func (this *SchemaWriter) WriteSchema() {
	//Write the package declaration
	fmt.Fprintf(this.Outfile, "package %s\n\n", this.PackageName)
	for _, table := range this.Tables {
		this.WriteType(table)
	}
}

// Write an individual table
func (this *SchemaWriter) WriteType(table *Table) {
	fmt.Fprintf(this.Outfile, "type %s struct {\n", CamelCase(table.Name))
	for _, column := range table.Columns {
		this.WriteField(&column)
	}
	fmt.Fprintf(this.Outfile, "}\n")
	fmt.Fprintf(this.Outfile,
		`
func (this *%s)Table() string {
    return "%s"
}
`,
		CamelCase(table.Name), table.Name)

}

// Write an individual field
func (this *SchemaWriter) WriteField(column *Column) {
	fmt.Fprintf(this.Outfile, "\t%s %s `db:\"%s\"`\n",
		CamelCase(column.Name), column.GoType(), column.Name)
}

// Load the database schema into memory using introspection, populating .Tables
func (this *SchemaWriter) LoadSchema(driver string, schema string, db *sql.DB) error {
	dialect := DialectByDriver(driver)
	tables, err := db.Query(dialect.ListTables(schema))
	if err != nil {
		return err
	}
	for tables.Next() {
		var ignored sql.NullString
		t := new(Table)
		tables.Scan(&t.Name)
		this.Tables = append(this.Tables, t)
		cols, err := db.Query(dialect.ListColumns(schema, *t))
		if err != nil {
			return err
		}
		for cols.Next() {
			c := new(Column)
			if strings.EqualFold(dialect.Name(), "sqlite3") {
				err = cols.Scan(&ignored, &c.Name, &c.DbType,
					&ignored, &ignored, &ignored)
			} else {
				err = cols.Scan(&c.Name, &c.DbType)
			}
			if err != nil {
				panic(err)
			}
			t.Columns = append(t.Columns, *c)
		}
		this.WriteType(t)
	}
	return nil
}
