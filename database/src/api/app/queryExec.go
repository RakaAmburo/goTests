package app

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"strconv"
)

type workerFunc func(results []sql.RawBytes)

var (
	count int
)

type Pepe struct {
	CurrencyID string  `json:"currency_id"`
}

func (Pepe) doSom(results []sql.RawBytes){
	for _, data := range results {
		fmt.Print(string(data), ",")

	}
	fmt.Println()
}

func test() {
	db, err := sql.Open("mysql", "ppaparini:FZS6RYQN@tcp(assetmgmtcore01.master.mlaws.com:6612)/asmgmtmlb")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	args2 := []interface{}{}
	pepe := Pepe{}
	ExecAndDo(db, CountUsuariosEntrantes, args2, pepe.doSom)

	fmt.Print("contduria",count)

	//args := []interface{}{"2019-02-08", "2019-02-15"}

	//ExecAndDo(db, UsuariosEntrantesMLB, args, printResults)
}

func setCount(results []sql.RawBytes){
	count, _ = strconv.Atoi(string(results[0]))
}

func printResults(results []sql.RawBytes) {
	for _, data := range results {
		fmt.Print(string(data), ",")

	}
	fmt.Println()
}

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

const (
	UsuariosEntrantesMLB = `
SELECT  
	MIN(date_created) AS date_investing, 
	user_id AS user_id
FROM  
	user_status_history 
WHERE 
	new_user_status_id = 'investing' AND 
    date_created >= ? AND 
    date_created < ?
GROUP BY  
	user_id 
ORDER BY  
	1, 2 ASC LIMIT 1000
`

	CountUsuariosEntrantes = `
    
      SELECT 
    COUNT(*)
    FROM
    (SELECT 
        MIN(date_created) AS date_investing
    FROM
        user_status_history
    WHERE
        new_user_status_id = 'investing'
            AND date_created >= '2019-02-08'
            AND date_created < '2019-02-15'
    GROUP BY user_id) AS users
      
`
)
