package service

import (
	"awesomeProject/core/model"
	"awesomeProject/domin/irepository"
	"awesomeProject/domin/iservice"
)

type personService struct {
	personRepository irepository.IPersonRepository
}

//NewPersonService ...
func NewPersonService(personRepository irepository.IPersonRepository) iservice.IPersonService {
	return &personService{
		personRepository: personRepository,
	}
}

//GetByID ...
func (r *personService) GetByID(ID int) (*model.Person, error) {
	return r.personRepository.FindByID(ID)
}

func (r *personService) GetAll() (*model.Persons, error) {
	return r.personRepository.FindAll()
}
