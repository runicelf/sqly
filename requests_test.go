package sqly

import (
	"github.com/jimlawless/whereami"
	"gopkg.in/DATA-DOG/go-sqlmock.v1"
	"testing"
)

func TestDBExec(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "insert into"
	mock.ExpectExec(query).WithArgs().WillReturnResult(sqlmock.NewResult(1, 1))

	dby := New(db)
	result, err := dby.Exec(query)
	checkErr(err, t)

	nInserted, err := result.LastInsertId()
	checkErr(err, t)

	expect(int64(1), nInserted, t)

	nAffected, err := result.RowsAffected()
	checkErr(err, t)

	expect(int64(1), nAffected, t)
}

type TestStruct struct {
	Column1 string `db:"column1"`
	Column2 string `db:"column2"`
}

func TestQuery(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "select from table"
	mock.ExpectQuery(query).WillReturnRows(mock.NewRows([]string{"column1", "column2"}).AddRow("field1", "field2"))

	dby := New(db)
	rows, err := dby.Query(query)
	checkErr(err, t)
	defer rows.Close()

	testStruct := new(TestStruct)
	for rows.Next() {
		err = rows.Scan(&testStruct.Column1, &testStruct.Column2)
		checkErr(err, t)
	}

	expect(TestStruct{"field1", "field2"}, *testStruct, t)
}

func TestQueryRow(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "select from table"
	mock.ExpectQuery(query).WillReturnRows(mock.NewRows([]string{"column1", "column2"}).AddRow("field1", "field2"))

	dby := New(db)
	row, err := dby.QueryRow(query)
	checkErr(err, t)

	testStruct := new(TestStruct)
	err = row.Scan(&testStruct.Column1, &testStruct.Column2)
	checkErr(err, t)

	expect(TestStruct{"field1", "field2"}, *testStruct, t)
}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "select from table"
	mock.ExpectQuery(query).WillReturnRows(mock.NewRows([]string{"column1", "column2"}).AddRow("field1", "field2"))

	dby := New(db)

	testStruct := new(TestStruct)
	err = dby.Get(testStruct, query)
	checkErr(err, t)

	expect(TestStruct{"field1", "field2"}, *testStruct, t)
}

func TestSelect(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "select from table"
	mock.ExpectQuery(query).WillReturnRows(mock.NewRows([]string{"column1", "column2"}).AddRow("field1", "field2"))

	dby := New(db)

	var testStruct []TestStruct
	err = dby.Select(&testStruct, query)
	checkErr(err, t)

	expect(TestStruct{"field1", "field2"}, testStruct[0], t)
}

func TestSelectWithoutClosing(t *testing.T) {
	db, mock, err := sqlmock.New()
	checkErr(err, t)

	query := "select from table"
	mock.ExpectQuery(query).WillReturnRows(mock.NewRows([]string{"column1", "column2"}).AddRow("field1", "field2"))

	dby := New(db)

	var testStruct []TestStruct
	rows, err := dby.SelectWithoutClosing(&testStruct, query)
	checkErr(err, t)
	rows.Close()

	expect(TestStruct{"field1", "field2"}, testStruct[0], t)
}

func expect(expected interface{}, got interface{}, t *testing.T) {
	source := whereami.WhereAmI(2)
	if expected != got {
		t.Fatal("\n", source, "\nExpected:", expected, "\nGot:", got)
	}
}

func checkErr(err error, t *testing.T) {
	source := whereami.WhereAmI(2)
	if err != nil {
		t.Fatal("\n", source, "\nNot expected error: ", err)
	}
}
