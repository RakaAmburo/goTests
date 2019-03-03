package main

import (
	"fmt"
	"sync"

	"github.com/mercadolibre/goTests/database/src/api/app"
)

func main() {
	fmt.Println("start")

	/*db, err := sql.Open("mysql", "ppaparini:FZS6RYQN@tcp(assetmgmtcore01.master.mlaws.com:6612)/asmgmtmlb")
	// if there is an error opening the connection, handle it
	if err != nil {
		fmt.Println("error")
		panic(err.Error())
	}
	defer db.Close()*/

	//execanddo to count
	//handlecount and extract info
	//calcular loops
	//instanciar todo y crear un publisher
	//crear jobs y asignarlos al worker con lo que necestia el exec and do

	writer := new(app.CsvWriter)
	writer.Init("resutl.txt")
	defer writer.Close()

	topic := &app.SqlTopic{}
	topic.Init(2)

	wait := &sync.WaitGroup{}
	wait.Add(1)
	consumer := new(app.SqlConsumer)
	consumer.Init(2, topic, writer, wait)

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

	wait.Wait()

}