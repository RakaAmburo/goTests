package main

import (
	"database/sql"
	"fmt"
	"github.com/mercadolibre/goTests/database/src/api/app"
	"sync"
)

func main() {
	fmt.Println("start")

	//falta ordenar, manejo de errores y un buen log

	db, err := sql.Open("mysql", "ppaparini:FZS6RYQN@tcp(assetmgmtcore01.master.mlaws.com:6612)/asmgmtmlb")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	//db.SetMaxIdleConns(5)

	var itemsPerPackage = 20
	var workerSize = 5
	timeBetweenJobs := &app.RandomWait{}
	timeBetweenJobs.Init(100, 200)

	workerTime := &app.RandomWait{}
	workerTime.Init(50, 100)

	argsCount := []interface{}{"2019-02-08", "2019-02-09"}
	argsLimited := []interface{}{"2019-02-08", "2019-02-09", itemsPerPackage, 0}

	handleCount := &app.HandleSqlCount{}
	handleCount.Init(itemsPerPackage)

	app.ExecAndDo(db, app.CountUsuariosEntrantesMLB, argsCount, handleCount.CalculateLoops)

	jobsNumber := handleCount.GetLoopSize()

	writer := new(app.CsvWriter)
	writer.Init("result.txt")
	defer writer.Close()

	topic := &app.SqlTopic{}
	topic.Init(jobsNumber)

	taskToWait := &sync.WaitGroup{}
	taskToWait.Add(1)

	consumer := new(app.SqlConsumer)
	consumer.Init(jobsNumber, topic, writer, taskToWait)

	workers := new(app.Workers)
	workers.Init(jobsNumber, workerSize, workerTime)
	for i := 0; i < jobsNumber; i++ {
		argsLimitedAux := append([]interface{}(nil), argsLimited...)
		if i != 0 {
			offset := itemsPerPackage * i
			(argsLimitedAux)[3] = offset
		}
		job := &app.SqlJob{}
		pkg := &app.SqlPackage{}
		pkg.Init(itemsPerPackage, i)
		job.Init(argsLimitedAux, app.SelectUsuariosEntrantesMLBLimited, topic, pkg, db)
		workers.AddWork(job)
		timeBetweenJobs.Wait()
	}

	taskToWait.Wait()

}