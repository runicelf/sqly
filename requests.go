package sqly

import (
	"database/sql"
	"reflect"
)

func (db *DB) Exec(query string) (sql.Result, error) {
	return db.conn.Exec(query)
}

func (db *DB) Query(query string) (*Rows, error) {
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return newRows(rows), nil
}

func (db *DB) QueryRow(query string) (*Row, error) {
	rows, err := db.conn.Query(query)
	if err != nil {
		return nil, err
	}
	return newRow(rows), nil
}

func (db *DB) Get(dest interface{}, query string) error {
	row, err := db.QueryRow(query)
	if err != nil {
		return err
	}
	return row.StructScan(dest)
}

func (db *DB) Select(dest interface{}, query string) error {
	rows, err := db.SelectWithoutClosing(dest, query)
	if err != nil {
		return err
	}
	rows.Close()

	return nil
}

func (db *DB) SelectWithoutClosing(dest interface{}, query string) (*Rows, error) {
	typeOfSliceElement := reflect.TypeOf(dest).Elem().Elem()
	elementSlice := reflect.ValueOf(dest).Elem()

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		elementInstance := reflect.New(typeOfSliceElement).Interface()

		err := rows.StructScan(elementInstance)
		if err != nil {
			rows.Close()
			return nil, err
		}

		elementSlice.Set(reflect.Append(elementSlice, reflect.ValueOf(elementInstance).Elem()))
	}

	return rows, nil
}
