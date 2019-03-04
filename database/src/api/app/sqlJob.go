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
}

func (job SqlJob) Init(args []interface{}, query string, topic Topic, pkg Package) {
	job.query = query
	job.args = args
	job.topic = topic
	job.pkg = pkg
}

func (job SqlJob) run() {
	fmt.Println("DOING SOME WORK")
	ExecAndDo(nil, job.query, job.args, job.BuildPackage)
	job.topic.Publish(job.pkg)
}

func (job SqlJob) BuildPackage(results []sql.RawBytes) {
	line := make([]string, len(results))
	for index, field := range results {
		line[index] = string(field)
	}
	job.pkg.Put(line)
}
