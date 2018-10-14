package iservice

import "awesomeProject/core/model"

//ProfileService interface
type IPersonService interface {
	//	SaveProfile(*model.Person) (*model.Persons, error)
	//	UploadProfile(int,*model.Person) (*model.Person, error)
	GetByID(int) (*model.Person, error)
	GetAll() (*model.Persons, error)
	//	GetByEmail(string) (*model.Person, error)
}
