package sqly

import (
	"database/sql"
	"fmt"
	"reflect"
)

const tag = "db"

type DB struct {
	conn *sql.DB
}

func Open(connectionString, driverName string) (*DB, error) {
	db, err := sql.Open(driverName, connectionString)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return New(db), nil
}

func New(db *sql.DB) *DB {
	return &DB{db}
}

func scanRowsInStruct(r *Rows, value interface{}) error {
	if r.rowsInfo.isFirstRun {
		columns, err := r.Columns()
		if err != nil {
			return err
		}
		r.rowsInfo.columns = columns
		r.rowsInfo.isFirstRun = false
	}

	pointers, err := getPointers(r.rowsInfo.columns, value)
	if err != nil {
		return err
	}

	return r.Rows.Scan(pointers...)
}

func getPointers(columns []string, responseModel interface{}) ([]interface{}, error) {
	var pointers []interface{}
	value := reflect.ValueOf(responseModel).Elem()
	mappingMap, err := mapping(value)
	if err != nil {
		return nil, err
	}

	for _, column := range columns {
		pointer, ok := mappingMap[column]
		if !ok {
			return nil, fmt.Errorf("missing field in struct: %s", column)
		}
		pointers = append(pointers, pointer)
	}
	return pointers, nil
}

func mapping(value reflect.Value) (map[string]interface{}, error) {
	mappingMap := make(map[string]interface{})
	for i := 0; i < value.NumField(); i++ {
		tag := value.Type().Field(i).Tag.Get(tag)
		if tag == "" {
			continue
		}
		var address interface{}
		//if value.Field(i).Kind() == reflect.Ptr {
		//	address = value.Field(i).Interface()
		//}else {
			address = value.Field(i).Addr().Interface()
		//}

		mappingMap[tag] = address
	}
	return mappingMap, nil
}
