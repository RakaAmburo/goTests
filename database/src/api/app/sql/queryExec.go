package sql

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type workerFunc func(results []sql.RawBytes)

func ExecAndDo(db *sql.DB, someQuery string, withArgs []interface{}, do workerFunc) {

	rows, err := db.Query(someQuery, withArgs...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	values := make([]sql.RawBytes, len(columns))
	results := make([]interface{}, len(values))
	for i := range values {
		results[i] = &values[i]
	}

	for rows.Next() {
		err := rows.Scan(results...)
		if err != nil {
			log.Fatal(err)
		}
		do(values)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
