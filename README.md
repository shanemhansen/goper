goper
=======

Utility for generating structs based off of database schemas. Intended to be
used with something like sqlx.

Included:

   * support for mysql/sqlite (postgres pending)


Documentation
-------------

[godoc](http://godoc.org/github.com/shanemhansen/goper)

Installation
--------

    # install the library:
    # install the sqlite development libraries
    go get github.com/shanemhansen/goper
    

Examples
--------

     #Add GOPATH/bin to your PATH
     go install github.com/shanemhansen/goper/...
     schema -driver sqlite3 -dsn path/to/file.sqlite3 > data.go
     schema -driver mysql -dsn "user:password@tcp(127.0.0.1:3306)/main" > data.go
