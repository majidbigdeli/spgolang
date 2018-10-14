package controller

import (
	"awesomeProject/core/model"
	"awesomeProject/domin/iservice"
)

//PersonController ...
type PersonController struct {
	Service iservice.IPersonService
}

//func NewPersonController(service iservice.IPersonService ) *personController {
//	return &personController{
//		service : service,
//	}
//}

//GetBy ...
func (c *PersonController) GetBy(id int) (*model.Person, error) {
	return c.Service.GetByID(id)
}

//GetAll ...
func (c *PersonController) Get() (*model.Persons, error) {
	return c.Service.GetAll()
}
