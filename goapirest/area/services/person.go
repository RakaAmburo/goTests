package services

import (
	"github.com/mercadolibre/goTests/goapirest/area/dao"
	"github.com/mercadolibre/goTests/goapirest/area/entities"
)

type PersonService struct {
	PD dao.PersonDAO
}

func (ps PersonService) CreatePerson(person *entities.Person) {

	ps.PD.CreatePerson(person)

}

func (ps PersonService) GetPerson(id int64) (*entities.Person, error) {

	person, err := ps.PD.GetPersonById(id)
	if err == nil {
		return person, nil
	} else {
		return nil, err
	}

}