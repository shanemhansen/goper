package goper

import (
	"fmt"
)

type SqliteDialect struct {
	MysqlDialect
}

func (this SqliteDialect) Name() string { return "sqlite3" }
func (this SqliteDialect) ListTables(dbname string) string {
	return fmt.Sprintf(`SELECT name as table_name FROM sqlite_master
WHERE type='table' ORDER BY name`)
}
func (this SqliteDialect) ListColumns(dbname string, table Table) string {
	return fmt.Sprintf(`PRAGMA table_info(%s)`, table.Name)
}
