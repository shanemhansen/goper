// schema -driver mysql|sqlite3 -dsn dsn
// Generate a set of a golang structs
package main
            
import (    
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
    "flag"
    "os"
	"github.com/shanemhansen/goper"
)

func main() {
    var dsn string
    var driver string
    var schema string
	flag.StringVar(&dsn,"dsn", "user:password@tcp(127.0.0.1:3306)/main", "database dsn")
    flag.StringVar(&driver, "driver", "mysql", "driver")
    flag.StringVar(&schema, "schema", "main", "schema")
    flag.Parse()
	conn, err := sql.Open(driver, dsn)
    if err != nil {
        panic(err)
    }
    os.Stdout.Write([]byte("package data\n"))
    writer := &goper.SchemaWriter{Outfile: os.Stdout, PackageName: "data"}
	err = writer.LoadSchema(driver, schema, conn)
    if err != nil {
        panic(err)
    }
}
