package goper

import "fmt"
import "bytes"

type MysqlDialect int

func (this MysqlDialect) Name() string { return "mysql" }
func (this MysqlDialect) CreateTable(table Table) string {
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
func (this MysqlDialect) DropTable(table Table) string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "DROP TABLE %s\n", table.Name)
	return string(buf.Bytes())
}
func (this MysqlDialect) InsertOne(table Table) string {
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
func (this MysqlDialect) ListTables(dbname string) string {
	return fmt.Sprintf(`SELECT table_name FROM  information_schema.tables where
table_schema = '%s' and table_type = 'BASE TABLE'`, dbname)
}
func (this MysqlDialect) ListColumns(dbname string, table Table) string {
	return fmt.Sprintf(`select column_name, data_type from information_schema.columns
where table_schema = '%s' and table_name='%s'`,dbname, table.Name)
}
func (this MysqlDialect) ListCollections(dbname string, table Table) string {
	return fmt.Sprintf(`
    SELECT DISTINCT i.TABLE_NAME, k.referenced_table_name, k.column_name, k.referenced_column_name
    FROM information_schema.TABLE_CONSTRAINTS i 
    LEFT JOIN information_schema.KEY_COLUMN_USAGE k
    ON i.CONSTRAINT_NAME = k.CONSTRAINT_NAME
    WHERE i.CONSTRAINT_TYPE = 'FOREIGN KEY'
    AND i.TABLE_SCHEMA = '%s'
    AND k.referenced_table_name='%s' and
    CONSTRAINT_TYPE='FOREIGN KEY'`, dbname, table.Name)
}
func (this MysqlDialect) ListReferences(dbname string, table Table) string {
	return fmt.Sprintf(`
    SELECT DISTINCT i.TABLE_NAME, k.referenced_table_name, k.column_name, k.referenced_column_name
    FROM information_schema.TABLE_CONSTRAINTS i 
    LEFT JOIN information_schema.KEY_COLUMN_USAGE k
    ON i.CONSTRAINT_NAME = k.CONSTRAINT_NAME
    WHERE i.CONSTRAINT_TYPE = 'FOREIGN KEY'
    AND i.TABLE_SCHEMA = '%s'
    AND k.table_name='%s' and
    CONSTRAINT_TYPE='FOREIGN KEY'`, dbname, table.Name)
}
