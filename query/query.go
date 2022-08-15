package query

import (
	"fmt"
	"reflect"
	"strings"
)

//
//type Query[T any, PtrT *T] struct {
//	table string // if this not present reflect it from the model
//}

type Builder[T any] struct {
	table string // if this not present reflect it from the model
	query string
}

func NewBuilder[T any](table string) Builder[T] {
	return Builder[T]{
		//table: reflect.TypeOf(T).Name(),
		table: table,
	}
}

func (b Builder[T]) where(in T) string {
	v := reflect.ValueOf(in)
	var values []string
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("db")
		keys := strings.Split(key, ",")
		value := v.Field(i)

		if v.Type().Field(i).Type.Kind() == reflect.Ptr {
			value = v.Field(i).Elem()
		}

		// if value exists, then add it to the query
		if value.IsZero() == false {
			values = append(values, fmt.Sprintf("%s = '%v'", keys[0], value))
		}
	}

	if len(values) == 0 {
		return ""
	}

	return fmt.Sprintf("WHERE %s", strings.Join(values, " AND "))
}

func (b Builder[T]) limit(limit int) string {
	if limit == 0 {
		return ""
	}
	return fmt.Sprintf("LIMIT %d", limit)
}

func (b Builder[T]) offset(offset int) string {
	if offset == 0 {
		return ""
	}
	return fmt.Sprintf("OFFSET %d", offset)
}

func (b Builder[T]) orderBy(orderBy string) string {
	if orderBy == "" {
		return ""
	}
	return fmt.Sprintf("ORDER BY %s", orderBy)
}

func (b Builder[T]) columns(in T) string {
	v := reflect.ValueOf(in)
	var fields []string
	for i := 0; i < v.NumField(); i++ {
		key := v.Type().Field(i).Tag.Get("db")
		keys := strings.Split(key, ",")
		fields = append(fields, keys[0])
	}
	return strings.Join(fields, ", ")
}

func (b Builder[T]) Find(in T) string {
	return fmt.Sprintf("SELECT %s FROM %s %s", b.columns(in), b.table, b.where(in))
}

func (b Builder[T]) FindAll() string {
	return fmt.Sprintf("SELECT * FROM %s", b.table)
}
