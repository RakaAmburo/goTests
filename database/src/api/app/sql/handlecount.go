package sql

import (
	"database/sql"
	"github.com/mercadolibre/goTests/database/src/api/app/tools"
	"math"
	"strconv"
)

type HandleSqlCount struct {
	loopSize        int
	itemsPerPackage int
}

func (h *HandleSqlCount) Init(itemsPerPackage int) {
	h.itemsPerPackage = itemsPerPackage
}

func (h *HandleSqlCount) CalculateLoops(results []sql.RawBytes) {
	totalString := string(results[0])
	total, err := strconv.Atoi(totalString)
	tools.CheckError("HandleSqlCount - CalculateLoops", err)
	h.loopSize = int(math.Ceil(float64(total) / float64(h.itemsPerPackage)))
}

func (h *HandleSqlCount) GetLoopSize() int {
	return h.loopSize
}
