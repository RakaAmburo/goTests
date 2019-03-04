package app

import "database/sql"

type HandleCount struct {

}

func (h HandleCount) CalculateLoops(results []sql.RawBytes){

}

func (h HandleCount) GetResult() int{
	return 1
}

