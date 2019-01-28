package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/mercadolibre/goTests/goapirest/area/entities"
	"github.com/mercadolibre/goTests/goapirest/area/services"
	"net/http"
	"strconv"
)

type PersonController struct{
	 Rg *gin.RouterGroup
	 Ps services.PersonService
}

type enhencedHandler struct{
	ps services.PersonService
}

type HandlerFunc func(*gin.Context, services.PersonService)

func (eh enhencedHandler)getEnhencedHandler(f HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
        f(c, eh.ps)
	}
}

func ( pc PersonController) RegisterEndpoints(){

	eh :=enhencedHandler{ps: pc.Ps}

	pc.Rg.GET("/person/:id", eh.getEnhencedHandler(getPersonById))
	pc.Rg.GET("/person", eh.getEnhencedHandler(getAllFiltered))

	pc.Rg.POST("/person", eh.getEnhencedHandler(setPerson))

	pc.Rg.DELETE("/person/:id", eh.getEnhencedHandler(deletePersonById))


}

func deletePersonById (c *gin.Context, ps services.PersonService){
	idStr := c.Param("id")
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		err := ps.DeletePerson(int64(id))
		c.JSON(201, err)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not provided"})
	}
}

func getPersonById(c *gin.Context, ps services.PersonService){

	idStr := c.Param("id")
	if idStr != "" {
		id, _ := strconv.Atoi(idStr)
		person, _ := ps.GetPerson(int64(id))
		c.JSON(200, person)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Id not provided"})
	}


	//c.String(200, "Getting person from id: " + c.Param("id") + " "+ c.Query("test"))
}

func setPerson(c *gin.Context, ps services.PersonService){

	p := &entities.Person{}
	err := c.ShouldBind(p)
	if err == nil {
		ps.CreatePerson(p)
		c.JSON(200, p)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

}

func getAllFiltered(c *gin.Context, ps services.PersonService){

	people, err := ps.GetAllFiltered()
	if err == nil {
		c.JSON(200, people)
	}
}