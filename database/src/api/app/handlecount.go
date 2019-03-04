package app

import "database/sql"

type HandleSqlCount struct {

}

func (h HandleSqlCount) CalculateLoops(results []sql.RawBytes){

}

func (h HandleSqlCount) GetResult() int{
	return 1
}

