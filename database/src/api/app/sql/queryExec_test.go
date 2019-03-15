package sql

import (
	"database/sql"
	"fmt"
	"testing"
)

func Test_Query(t *testing.T) {

	test()

}

var (
	count int
)

type Pepe struct {
	CurrencyID string `json:"currency_id"`
}

func (Pepe) doSom(results []sql.RawBytes) {
	for _, data := range results {
		fmt.Print(string(data), ",")

	}
	fmt.Println()
}

func test() {
	db, err := sql.Open("mysql", "")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	args2 := []interface{}{}
	pepe := Pepe{}
	ExecAndDo(db, CountNewUsersMLB, args2, pepe.doSom)

	fmt.Print("contduria", count)
}
