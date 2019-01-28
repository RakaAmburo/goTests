package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mercadolibre/goTests/goapirest/area/controllers"
	"github.com/mercadolibre/goTests/goapirest/area/persist"
	"github.com/mercadolibre/goTests/goapirest/area/services"
)

func main() {
	r := gin.Default()
	route := r.Group("/api")

	db := GetProps().Database
	user := db.User
	pass := db.PassWord
	host := db.Host
	port := db.Port
	url := fmt.Sprintf(db.Url, user, pass, host, port)


	DB, err := gorm.Open("mysql", url)

	if err != nil {
		panic("failed to connect database")
	}

	pd := persist.PersonDAO{DB: DB}
	ps := services.PersonService{PD: pd}
	pc := controllers.PersonController{Rg: route, Ps: ps}
	pc.RegisterEndpoints()

	r.Run(":8080")
}
