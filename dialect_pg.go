package goper

import "fmt"
import "bytes"

type PgDialect int

func (this PgDialect) Name() string { return "postgres" }
func (this PgDialect) CreateTable(table Table) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "CREATE TABLE %s(", table.Name)
	count := len(table.Columns)
	for i := range table.Columns {
		c := table.Columns[i]
		fmt.Fprintf(&buf, "\n%s\t%s", c.Name, c.DbType)
		if i != (count - 1) {
			fmt.Fprintf(&buf, ",")
		}
	}
	fmt.Fprintf(&buf, "\n);\n")
	return string(buf.Bytes())
}
func (this PgDialect) DropTable(table Table) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "DROP TABLE %s\n", table.Name)
	return string(buf.Bytes())
}
func (this PgDialect) InsertOne(table Table) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "INSERT INTO %s(", table.Name)
	count := len(table.Columns)
	for i := range table.Columns {
		c := table.Columns[i]
		fmt.Fprintf(&buf, "\n\t%s", c.Name)
		if i != (count - 1) {
			fmt.Fprintf(&buf, ",")
		}
	}
	fmt.Fprintf(&buf, "\n) VALUES (")
	for i := range table.Columns {
		fmt.Fprintf(&buf, "?")
		if i != (count - 1) {
			fmt.Fprintf(&buf, ",")
		}
	}
	fmt.Fprintf(&buf, ");")
	return string(buf.Bytes())

}
func (this PgDialect) ListTables(dbname string) string {
	return `SELECT c.relname AS Tables_in FROM pg_catalog.pg_class c
LEFT JOIN pg_catalog.pg_namespace n ON n.oid = c.relnamespace
WHERE pg_catalog.pg_table_is_visible(c.oid)
AND c.relkind = 'r'
AND relname NOT LIKE 'pg_%'
ORDER BY 1`
}
func (this PgDialect) ListColumns(dbname string, table Table) string {
	return fmt.Sprintf(`SELECT column_name, data_type 
FROM information_schema.columns
WHERE table_name='%s'`, table.Name)

}
