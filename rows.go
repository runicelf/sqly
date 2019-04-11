package sqly

import (
	"database/sql"
)

type Rows struct {
	*sql.Rows
	rowsInfo *rowsInfo
}

type rowsInfo struct {
	columns    []string
	isFirstRun bool
}

func newRows(rows *sql.Rows) *Rows {
	return &Rows{Rows: rows, rowsInfo: &rowsInfo{isFirstRun: true}}
}

func (r *Rows) StructScan(dest interface{}) error {
	return scanRowsInStruct(r, dest)
}

func (r *Rows) NextResultSet() bool {
	r.rowsInfo.isFirstRun = true
	return r.Rows.NextResultSet()
}
