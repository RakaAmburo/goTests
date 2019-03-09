package app

import (
	"database/sql"
	"fmt"
)

type SqlJob struct {
	query string
	args  []interface{}
	topic Topic
	pkg Package
	db *sql.DB
}

func (job *SqlJob) Init(args []interface{}, query string, topic Topic, pkg Package, db *sql.DB) {
	job.query = query
	job.args = args
	job.topic = topic
	job.pkg = pkg
	job.db = db
}

func (job *SqlJob) run() {
	fmt.Println("DOING SOME WORK")
	ExecAndDo(job.db, job.query, job.args, job.BuildPackage)
	fmt.Println("DONED QUERY")
	job.topic.Publish(job.pkg)
}

func (job *SqlJob) BuildPackage(results []sql.RawBytes) {
	line := make([]string, len(results))
	for index, field := range results {
		line[index] = string(field)
	}
	job.pkg.Put(line)
	fmt.Println("linea:")
	fmt.Println(line)
}
