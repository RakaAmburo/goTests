package persist

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

func (pd PersonDAO) DeleteById(id int64) error{

	if err := pd.DB.Where("id = ?", id).Delete(&entities.Person{}).Error; err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}

}

func (pd PersonDAO) UpdateById(person *entities.Person){

}

func (pd PersonDAO) GetAllFiltered() (*[]entities.Person, error){

	people := make([]entities.Person, 0)
	if err := pd.DB.Find(&people).Error; err != nil {
		fmt.Println(err)
		return nil, err
	} else {
		return &people, nil
	}
}

