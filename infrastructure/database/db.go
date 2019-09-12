package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/hmarf/sample_clean/config"
	"github.com/hmarf/sample_clean/interface/datastore"
)

type MysqlData struct {
	DB *sql.DB
}

func SetUpDB(database config.Database) *MysqlData {
	connectDB := fmt.Sprintf(
		"%s:%s@%s(%s:%s)/%s?%s",
		database.UserName,
		database.Password,
		database.Protocol,
		database.Host,
		database.Port,
		database.DBName,
		database.Option,
	)
	fmt.Println(connectDB)
	DB, err := sql.Open("mysql", connectDB)
	if err != nil {
		log.Fatal(err)
	}
	conn := MysqlData{DB: DB}
	return &conn
}

//interface層で使用可能なsqlのクエリ操作定義
func (conn *MysqlData) Exec(cmd string, args ...interface{}) (datastore.Result, error) {
	result, err := conn.DB.Exec(cmd, args...)
	if err != nil {
		return nil, err
	}
	return &SqlResult{Result: result}, nil
}

//interface層で使用可能なsqlのクエリ操作定義
func (conn *MysqlData) Query(cmd string, args ...interface{}) (datastore.Rows, error) {
	rows, err := conn.DB.Query(cmd, args...)
	if err != nil {
		return nil, err
	}
	return &SqlRows{Rows: rows}, nil
}

func (conn *MysqlData) QueryRow(cmd string, args ...interface{}) datastore.Row {
	row := conn.DB.QueryRow(cmd, args...)
	return &SqlRow{Row: row}
}

type SqlResult struct {
	Result sql.Result
}

func (r *SqlResult) LastInsertId() (int64, error) {
	return r.Result.LastInsertId()
}

func (r *SqlResult) RowsAffected() (int64, error) {
	return r.Result.RowsAffected()
}

//なぜポインタ型なのか？
type SqlRows struct {
	Rows *sql.Rows
}

func (r SqlRows) Scan(ctr ...interface{}) error {
	return r.Rows.Scan(ctr...)
}

func (r SqlRows) Next() bool {
	return r.Rows.Next()
}

func (r SqlRows) Close() error {
	return r.Rows.Close()
}

type SqlRow struct {
	Row *sql.Row
}

func (r SqlRow) Scan(ctr ...interface{}) error {
	return r.Row.Scan(ctr...)
}
