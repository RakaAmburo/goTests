package app_test

import (
	"fmt"
	"github.com/mercadolibre/goTests/database/src/api/app"
	"testing"
	"time"
)

func Test_Topic(t *testing.T) {


	topic := app.SqlTopic{}
	topic.Init(2)
	pkg := app.SqlPackage{}
	pkg.Init(3, 0)

	pkg.Put([]string{"Country", "City", "Population222"})
	pkg.Put([]string{"Japan", "Tokyo", "923456"})
	pkg.Put([]string{"Australia", "Sydney", "789650"})

	pkg2 := app.SqlPackage{}
	pkg2.Init(3, 1)
	topic.Publish(&pkg)
	topic.Publish(&pkg2)

	time.Sleep(1 * time.Second)

	fmt.Println(topic.Get(0))

}