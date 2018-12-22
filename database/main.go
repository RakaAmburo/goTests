package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type Amount struct {
	CurrencyID string `json:"currency_id"`
	Value      float64  `json:"value"`
}

type Balance struct {
	Available     *Amount `json:"available_balance"`
	Invested      *Amount `json:"invested_balance"`
	NonReflected  *Amount `json:"non_reflected_balance"`
	Delta         *Amount `json:"delta"`
	OriginDelta   *Amount `json:"origin_delta,omitempty"`
	ShareQuantity float64       `json:"share_quantity"`
}

const (
	// BundleSelectStatusSQL selects a bundle status by id and provider id
	BundleSelectStatusSQL = `
		SELECT 

           o.balance_json AS balance_json,
           o.date_created as balance_date
            
		FROM
			bundle b
			INNER JOIN fund AS f ON b.fund_id = f.id
			INNER JOIN packet AS p ON b.id = p.bundle_id
			INNER JOIN operation AS o ON p.id = o.packet_id
		WHERE
			 o.origin_user_id = 236870040 LIMIT 1
	`

	// BundleSelectStatusPacketFilterSQL filter for packet id
	BundleSelectStatusPacketFilterSQL = `
		AND p.id = ?
	`
)

func main() {
	funcName()
}

func funcName() {
	fmt.Println("Go MySQL Tutorial")
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "ppaparini:VB5X9E1J@tcp(assetmgmtcore00.master.mlaws.com:6612)/asmgmtcore")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	// Execute the query
	rows, err := db.Query(BundleSelectStatusSQL)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Get column names
	columns, err := rows.Columns()
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// Make a slice for the values
	values := make([]sql.RawBytes, len(columns))
	// rows.Scan wants '[]interface{}' as an argument, so we must copy the
	// references into such a slice
	// See http://code.google.com/p/go-wiki/wiki/InterfaceSlice for details
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}
	// Fetch rows
	for rows.Next() {
		// get RawBytes from data
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		// Now do something with the data.
		// Here we just print each column as a string.
		var value string
		for i, col := range values {
			// Here we can check if the value is nil (NULL value)
			if col == nil {
				value = "NULL"
			} else {
				value = string(col)
			}
			fmt.Println(columns[i], ": ", value)

			//var balance Balance
			//json.Unmarshal([]byte("{}"), &balance)

			//fmt.Println(balance.Delta.Value)

		}
		fmt.Println("-----------------------------------")
	}
	if err = rows.Err(); err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
}