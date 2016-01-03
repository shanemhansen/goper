// schema -driver mysql|sqlite3|postgres -dsn dsn
// Generate a set of a golang structs
package main
            
import (    
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    _ "github.com/mattn/go-sqlite3"
	_ "github.com/lib/pq"
    "flag"
	"fmt"
    "os"
	"github.com/shanemhansen/goper"
	"log"
)

var dsn string
var driver string
var schema string
var logger *log.Logger
var verbose bool

func init() {
	flag.StringVar(&dsn,"dsn", "user:password@tcp(127.0.0.1:3306)/main", "database dsn")
    flag.StringVar(&driver, "driver", "mysql", "driver")
    flag.StringVar(&schema, "schema", "main", "schema")
	flag.BoolVar(&verbose, "verbose", false, "Print debugging")
    flag.Parse()

	logger = log.New(goper.ColourStream{os.Stderr}, " [XXXX] ", log.LstdFlags)
}

func main() {
	conn, err := sql.Open(driver, dsn)
    if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        panic(err)
    }
	err = conn.Ping()
    if err != nil {
		logger.Panic(err)
    } else if verbose {
		logger.Printf("Ping Worked\n")
	}
    os.Stdout.Write([]byte("package data\n"))
    writer := &goper.SchemaWriter{Outfile: os.Stdout, PackageName: "data"}
	//os.Stdout.Write([]byte(schema))
	err = writer.LoadSchema(driver, schema, conn)
    if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
        panic(err)
    }
}
