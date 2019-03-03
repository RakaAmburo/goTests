package app

import (
	"database/sql"
	"fmt"
)

type SqlJob struct {

}

func (SqlJob) Init(){
	//pasar db topic consumer etc args y publisher func
}

func ( job SqlJob) run() {
	fmt.Println("DOING SOME WORK")
	args2 := []interface{}{}
	pepe := Pepe{}
	ExecAndDo(nil, CountUsuariosEntrantes, args2, pepe.doSom)
}

type Publisher struct {

}

func ( pub Publisher) publishResults(results []sql.RawBytes) {
	for _, data := range results {
		fmt.Print(string(data), ",")

	}
	fmt.Println()
}

func HandleSqlCount(esults []sql.RawBytes){
	//crear toda la extructura y instancia de publisher
}