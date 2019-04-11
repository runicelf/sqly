package sqly

import (
	"database/sql"
)

type Row struct {
	*sql.Rows
}

func newRow(rows *sql.Rows) *Row {
	return &Row{rows}
}

func (r *Row) Scan(dest ...interface{}) error {
	defer r.Close()
	if !r.Next() {
		return sql.ErrNoRows
	}

	return r.Rows.Scan(dest...)
}

func (r *Row) StructScan(dest interface{}) error {
	defer r.Close()
	if !r.Next() {
		return sql.ErrNoRows
	}

	return scanRowsInStruct(newRows(r.Rows), dest)
}
