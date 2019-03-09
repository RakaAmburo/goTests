package main

import (
	"database/sql"
	"fmt"
	"sync"

	"github.com/mercadolibre/goTests/database/src/api/app"
)

func main() {
	fmt.Println("start")



	/*pkg := new(app.SqlPackage)
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
	topic.Publish(pkg2)*/

	//---------

	db, err := sql.Open("mysql", "ppaparini:FZS6RYQN@tcp(assetmgmtcore01.master.mlaws.com:6612)/asmgmtmlb")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()
	//db.SetMaxIdleConns(5)

	argsCount := []interface{}{"2019-02-08", "2019-02-09"}
	argsLimited := []interface{}{"2019-02-08", "2019-02-09", 20, 0}

	handleCount := &app.HandleSqlCount{}
	handleCount.Init(20)

	app.ExecAndDo(db, app.CountUsuariosEntrantes, argsCount, handleCount.CalculateLoops)

	jobSize := handleCount.GetLoopSize()

	writer := new(app.CsvWriter)
	writer.Init("result.txt")
	defer writer.Close()

	topic := &app.SqlTopic{}
	topic.Init(jobSize)

	task := &sync.WaitGroup{}
	task.Add(1)
	consumer := new(app.SqlConsumer)
	consumer.Init(jobSize, topic, writer, task)

	workers := new(app.Workers)
	workers.Init(jobSize, 5)
	for i := 0; i < jobSize; i++ {
		job := &app.SqlJob{}
		pkg := &app.SqlPackage{}
		pkg.Init(10, i)
		job.Init(argsLimited, app.UsuariosEntrantesMLB, topic, pkg, db)
		workers.AddWork(job)
	}


	//---------

	task.Wait()

}