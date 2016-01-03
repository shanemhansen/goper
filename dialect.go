package goper

// A Dialect is a set of functions for generating the SQL
// to obtain various metadata about a database.
type Dialect interface {
	CreateTable(table Table) string
	DropTable(table Table) string
	InsertOne(table Table) string
	ListTables(string) string
	ListColumns(string, Table) string
	Name() string
}

var dialects map[string]Dialect

func RegisterDialect(driver string, dialect Dialect) {
	dialects[driver] = dialect
}
func DialectByDriver(driver string) Dialect {
	return dialects[driver]
}

func init() {
	dialects = make(map[string]Dialect)
	RegisterDialect("mysql", new(MysqlDialect))
	RegisterDialect("sqlite3", new(SqliteDialect))
	RegisterDialect("postgres", new(PgDialect))
}
