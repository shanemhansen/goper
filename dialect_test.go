package goper

import "testing"

func TestDialect(t *testing.T) {
	dialects := []string{"mysql", "sqlite"}
	for _, dialect := range dialects {
		d := DialectByDriver(dialect)
		if d == nil {
			panic(d)
		}
		products := Table{Name: "products",
			Columns: []Column{
				Column{Name: "product_id",
					DbType: "integer"},
				Column{Name: "product_name",
					DbType: "text"},
			}}
		d.InsertOne(products)
		d.CreateTable(products)
		d.DropTable(products)
		d.ListTables("main")
	}
}
