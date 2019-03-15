package main

import (
	"database/sql"
	"fmt"
	"github.com/mercadolibre/goTests/database/src/api/app"
	"github.com/mercadolibre/goTests/database/src/api/app/consumers"
	"github.com/mercadolibre/goTests/database/src/api/app/jobs"
	"github.com/mercadolibre/goTests/database/src/api/app/packages"
	sql2 "github.com/mercadolibre/goTests/database/src/api/app/sql"
	"github.com/mercadolibre/goTests/database/src/api/app/tools"
	"github.com/mercadolibre/goTests/database/src/api/app/topics"
	"github.com/mercadolibre/goTests/database/src/api/app/writers"
	"sync"
)



func main() {
	fmt.Println("start")

	//, manejo de errores y un buen log/reporte tests, sacar a tasks y merojar config

	//location of properties file in your machine
	path := "../../../../secure"
	dbConf := app.GetDbPropertes(path, "CORE_MLB")
    format := "%s:%s@tcp(%s:%d)/%s"
    dataSourceName := fmt.Sprintf(format, dbConf.User, dbConf.PassWord, dbConf.Url, dbConf.Port, dbConf.Schema)
	db, err := sql.Open("mysql", dataSourceName)
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	//db.SetMaxIdleConns(5)

	var itemsPerPackage = 20
	var workerSize = 5
	timeBetweenJobs := &tools.RandomWait{}
	timeBetweenJobs.Init(100, 200)

	workerTime := &tools.RandomWait{}
	workerTime.Init(50, 100)

	argsCount := []interface{}{"2019-02-08", "2019-02-09"}
	argsLimited := []interface{}{"2019-02-08", "2019-02-09", itemsPerPackage, 0}

	handleCount := &sql2.HandleSqlCount{}
	handleCount.Init(itemsPerPackage)

	sql2.ExecAndDo(db, sql2.CountUsuariosEntrantesMLB, argsCount, handleCount.CalculateLoops)

	jobsNumber := handleCount.GetLoopSize()

	writer := new(writers.CsvWriter)
	writer.Init("result.txt")
	defer writer.Close()

	topic := &topics.SqlTopic{}
	topic.Init(jobsNumber)

	taskToWait := &sync.WaitGroup{}
	taskToWait.Add(1)

	consumer := new(consumers.SqlConsumer)
	consumer.Init(jobsNumber, topic, writer, taskToWait)

	workers := new(app.Workers)
	workers.Init(jobsNumber, workerSize, workerTime)
	for i := 0; i < jobsNumber; i++ {
		argsLimitedAux := append([]interface{}(nil), argsLimited...)
		if i != 0 {
			offset := itemsPerPackage * i
			(argsLimitedAux)[3] = offset
		}
		job := &jobs.SqlJob{}
		pkg := &packages.SqlPackage{}
		pkg.Init(itemsPerPackage, i)
		job.Init(argsLimitedAux, sql2.SelectUsuariosEntrantesMLBLimited, topic, pkg, db)
		workers.AddWork(job)
		timeBetweenJobs.Wait()
	}

	taskToWait.Wait()

}
