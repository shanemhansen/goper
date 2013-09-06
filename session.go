package goper

import "database/sql"
import "reflect"
import "errors"
import "strings"
import "fmt"

//  Session is deprecated in favor of sqlx
var Done = errors.New("iteration done")

type Session struct {
	DB *sql.DB
}
type TableStruct interface {
	Table() string
}
type StructGenerator func(interface{}) error

func (this *Session) FromQuery(query string, args ...interface{}) (gen StructGenerator, err error) {
	rows, err := this.DB.Query(query, args...)
	if err != nil {
		return
	}
	gen = func(obj interface{}) (err error) {
		if !rows.Next() {
			return Done
		}
		var value reflect.Value
		switch obj := obj.(type) {
		case reflect.Value:
			value = reflect.Indirect(obj)
		default:
			value = reflect.Indirect(reflect.ValueOf(obj))
		}
		t := value.Type()
		fields := make([]interface{}, 0)
		for i := 0; i < t.NumField(); i++ {
			field := value.Field(i).Addr().Interface()
			tag := t.Field(i).Tag
			if len(tag.Get("db")) == 0 || strings.Contains(tag.Get("db"), ":") {
				continue
			}
			fields = append(fields, field)
		}
		err = rows.Scan(fields...)
		if err != nil {
			panic(err)
			return
		}
		return
	}
	return
}
func (this *Session) GetChildren(obj interface{}, children string) error {
	var value reflect.Value
	switch obj := obj.(type) {
	case reflect.Value:
		value = reflect.Indirect(obj)
	default:
		value = reflect.Indirect(reflect.ValueOf(obj))
	}
	t := value.Type()
	fields := make([]interface{}, 0)

	sql := "SELECT * FROM %s WHERE %s = %d\n"
	child := value.FieldByName(children)
	field := child.Addr().Interface()
	field_t, found := t.FieldByName(children)
	if !found {
		return errors.New(children + " not found")
	}
	tag := field_t.Tag
	if len(tag.Get("db")) == 0 {
		return errors.New("No db struct tag")
	}
	f := strings.Split(tag.Get("db"), ":")
	fields = append(fields, field)
	table := f[0]
	column := f[1]
	//assumption first field is pk
	childName := children[:len(children)-3]
	sql = fmt.Sprintf(sql, table, column, reflect.Indirect(value.FieldByName(childName)).Int())
	fmt.Println(sql)
	gen, err := this.FromQuery(sql)
	if err != nil {
		return err
	}
	for {
		t := reflect.New(field_t.Type.Elem())
		if gen(t) == Done {
			break
		}
		if child.Kind() == reflect.Slice {
			child.Set(reflect.Append(child, reflect.Indirect(t)))
		} else {
			child.Set(t)
		}
	}
	return nil
}
