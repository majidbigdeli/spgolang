package irepository

import "awesomeProject/core/model"

//ProfileRepository ...
type IPersonRepository interface {
	Save(person *model.Person) error
	Update(int,*model.Person) error
	Delete(int) error
	FindByID(int) (*model.Person, error)
	FindByEmail(string) (*model.Person, error)
	FindAll() (*model.Persons, error)
}
