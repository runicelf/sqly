**SQLY**

Sugar for working with SQL

+ `Open(connectionString, driverName string) (*DB, error)`

`type sqly.DB`
+ `func (db *DB) Exec(query string) (sql.Result, error)`
+ `func (db *DB) Query(query string) (*Rows, error)`
+ `func (db *DB) QueryRow(query string) (*Row, error)`
+ `func (db *DB) Get(dest interface{}, query string) error`
+ `func (db *DB) Select(dest interface{}, query string) error`
+ `func (db *DB) SelectWithoutClosing(dest interface{}, query string) (*Rows, error)`

`type sqly.Rows`
+ `func (r *Rows) Scan(dest ...interface{}) error`
+ `func (r *Rows) StructScan(dest interface{}) error`
+ `func (r *Rows) NextResultSet() bool`

`type sqly.Row`
+ `func (r *Row) Scan(dest ...interface{}) error`
+ `func (r *Row) StructScan(dest interface{}) error`


Code examples:

    type CustomStruct struct {
        Field1 sql.NullString `db:"column1"`
        Field2 sql.NullString `db:"column2"`
        Field3 sql.NullString
    }
    
    func f1(db *sqly.DB)  {
        rows, err := db.Query("SELECT column1, column2 FROM procedure();")
        checkErr(err)
        defer rows.Close()
    
        if rows.Next() {
            var cStruct CustomStruct
            err = rows.StructScan(&cStruct)
            checkErr(err)
        }
    }
    
    func f2(db *sqly.DB)  {
        row, err := db.QueryRow("SELECT column1, column2 FROM procedure();"))
        checkErr(err)
    
        var cStruct CustomStruct
        err = row.StructScan(&cStruct)
        checkErr(err)
    }
    
    func f3(db *sqly.DB)  {
        row, err := db.QueryRow("SELECT column1, column2 FROM procedure();"))
        checkErr(err)
    
        var cStruct CustomStruct
        err = row.StructScan(&cStruct)
        checkErr(err)
    }
    
    func f4(db *sqly.DB)  {
        var cStruct CustomStruct
        err := db.Get(&cStruct, "SELECT column1, column2 FROM procedure();"))
        checkErr(err)
    }
    
    func f5(db *sqly.DB)  {
        var cStruct []CustomStruct
        err := db.Select(&cStruct, "SELECT column1, column2 FROM procedure();"))
        checkErr(err)
    }
    
    func f6(db *sqly.DB)  {
    	var cStruct []CustomStruct
    	rows, err := db.SelectWithoutClosing(&currency, "SELECT column1, column2 FROM procedure();")
    	checkErr(err)
    	defer rows.Close()
    }