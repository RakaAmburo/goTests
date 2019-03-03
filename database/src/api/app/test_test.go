package app_test

import (
	"fmt"
	"github.com/mercadolibre/goTests/database/src/api/app"
	"testing"
	"time"
)

func Test_Test(t *testing.T) {

	test := app.Test{}
	test.Init(2)
	pkg := app.SqlPackage{}
	pkg.Init(3, 0)

	pkg.Put([]string{"Country", "City", "Population222"})
	pkg.Put([]string{"Japan", "Tokyo", "923456"})
	pkg.Put([]string{"Australia", "Sydney", "789650"})

	pkg2 := app.SqlPackage{}
	pkg2.Init(3, 1)
	test.Publish(&pkg)
	test.Publish(&pkg2)

	time.Sleep(1 * time.Second)

	fmt.Println(test.Get(0))

}
