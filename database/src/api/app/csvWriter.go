package app

import (
	"encoding/csv"
	"log"
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

func (cw *CsvWriter) BulkWrite(pkg Package) {
	lines := pkg.Extract()

	for _, line := range lines {
		err := cw.writer.Write(line)
		checkError("Cant write to file", err)
		cw.writer.Flush()
	}
}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}




