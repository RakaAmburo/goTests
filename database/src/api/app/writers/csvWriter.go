package writers

import (
	"encoding/csv"
	"github.com/mercadolibre/goTests/database/src/api/app"
	"github.com/mercadolibre/goTests/database/src/api/app/tools"
	"os"
)

type CsvWriter struct {
	writer *csv.Writer
	file   *os.File
}

func (cw *CsvWriter) Init(fileName string) {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		os.Exit(1)
	}
	cw.file = file
	cw.writer = csv.NewWriter(file)
}

func (cw *CsvWriter) Write(line []string) {
	cw.writer.Write(line)
}

func (cw *CsvWriter) Close() {
	cw.file.Close()
}

func (cw *CsvWriter) BulkWrite(pkg app.Package) {
	lines := pkg.Extract()

	for _, line := range lines {
		if line != nil {
			err := cw.writer.Write(line)
			tools.CheckError("Cant write to file", err)
			cw.writer.Flush()
		}

	}
}
