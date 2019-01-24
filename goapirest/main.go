package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/mercadolibre/goTests/goapirest/area/controllers"
	"github.com/mercadolibre/goTests/goapirest/area/dao"
	"github.com/mercadolibre/goTests/goapirest/area/services"
)

func main() {
	r := gin.Default()
	route := r.Group("/api")

	DB, err := gorm.Open("mysql", "root:1234qwer@tcp(127.0.0.1:3306)/mine?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("failed to connect database")
	}

	pd := dao.PersonDAO{DB: DB}
	ps := services.PersonService{PD: pd}
	pc := controllers.PersonController{Rg: route, Ps: ps}
	pc.RegisterEndpoints()

	r.Run(":8080")
}
