package services

import (
	"github.com/mercadolibre/goTests/goapirest/area/persist"
	"github.com/mercadolibre/goTests/goapirest/area/entities"
)

type PersonService struct {
	PD persist.PersonDAO
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

func (ps PersonService) GetAllFiltered() (*[]entities.Person, error){

	people, err := ps.PD.GetAllFiltered()
	if err == nil {
		return people, nil
	} else {
		return nil, err
	}
}

func (ps PersonService) DeletePerson(id int64) error{
	err := ps.PD.DeleteById(id)
	return err
}