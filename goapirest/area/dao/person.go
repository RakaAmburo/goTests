package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/mercadolibre/goTests/goapirest/area/entities"
)

type PersonDAO struct{
	DB *gorm.DB
}

func (pd PersonDAO) CreatePerson(person *entities.Person) {

   pd.DB.Create(&person)

}

func (pd PersonDAO) GetPersonById(id int64) (*entities.Person, error){

	person := entities.Person{}
	if err :=pd.DB.Where("id = ?", id).First(&person).Error; err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return &person, nil
	}

}