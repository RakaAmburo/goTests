package main

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/mercadolibre/goTests/database/src/api/app"
)

func main() {
	fmt.Println("start")

	writer := new(app.CsvWriter)
	writer.Init("resutl.txt")
	defer writer.Close()

	topic := &app.SqlTopic{}
	topic.Init(2)

	task := &sync.WaitGroup{}
	task.Add(1)
	consumer := new(app.SqlConsumer)
	consumer.Init(2, topic, writer, task)

	pkg := new(app.SqlPackage)
	pkg.Init(3, 0)

	pkg.Put([]string{"Country", "City", "Population"})
	pkg.Put([]string{"Japan", "Tokyo", "923456"})
	pkg.Put([]string{"Australia", "Sydney", "789650"})

	pkg2 := new(app.SqlPackage)
	pkg2.Init(3, 1)

	pkg2.Put([]string{"Argentina", "Buenos Aires", "1"})
	pkg2.Put([]string{"France", "paris", "923456"})
	pkg2.Put([]string{"Spain", "madrid", "789650"})

	topic.Publish(pkg)
	topic.Publish(pkg2)

	//---------

	db, err := sql.Open("mysql", "ppaparini:FZS6RYQN@tcp(assetmgmtcore01.master.mlaws.com:6612)/asmgmtmlb")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()

	handleCount := app.HandleSqlCount{}

	app.ExecAndDo(db, app.CountUsuariosEntrantes, []interface{}{}, handleCount.CalculateLoops)

	jobSize := handleCount.GetResult()
	workers := new(app.Workers)
	workers.Init(jobSize)
	for i := 0; i < jobSize; i++ {
		job := &app.SqlJob{}
		job.Init([]interface{}{}, "", topic, nil, db)
		workers.AddWork(job)
	}


	//execanddo to count
	//handlecount and extract info
	//calcular loops
	//instanciar todo y crear un publisher <- NO SE VA A ENCARGAR EL JOB
	//crear jobs y asignarlos al worker con lo que necestia el exec and do


	//---------

	task.Wait()

}