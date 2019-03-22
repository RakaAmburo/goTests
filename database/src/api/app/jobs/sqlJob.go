package jobs

import (
	"database/sql"
	"github.com/mercadolibre/goTests/database/src/api/app"
	sql2 "github.com/mercadolibre/goTests/database/src/api/app/sql"
)

type SqlJob struct {
	query string
	args  []interface{}
	topic app.Topic
	pkg   app.Package
	db    *sql.DB
}

func (job *SqlJob) Init(args []interface{}, query string, topic app.Topic, pkg app.Package, db *sql.DB) {
	job.query = query
	job.args = args
	job.topic = topic
	job.pkg = pkg
	job.db = db
}

func (job *SqlJob) Run() {
	sql2.ExecAndDo(job.db, job.query, job.args, job.BuildPackage)
	job.topic.Publish(job.pkg)
}

func (job *SqlJob) GetPackageNumber() int {
	return job.pkg.GetIndex()
}

func (job *SqlJob) BuildPackage(results []sql.RawBytes) {
	line := make([]string, len(results))
	for index, field := range results {
		line[index] = string(field)
	}
	job.pkg.Put(line)
}
