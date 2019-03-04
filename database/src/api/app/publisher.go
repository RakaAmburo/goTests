package app

import (
	"database/sql"
)

type Publisher struct {
	topic Topic
	pkg SqlPackage
}

func (pub Publisher) publishResults(results []sql.RawBytes) {
	pkg := new(SqlPackage)
	for _, field := range results {

	}

}
