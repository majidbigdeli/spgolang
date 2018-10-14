package mssql

import (
	"awesomeProject/core/model"
	"awesomeProject/domin/irepository"

	"github.com/jmoiron/sqlx"
)

type personRepository struct {
	db *sqlx.DB
}

//NewProfileRepository ...
func NewPersonRepository(db *sqlx.DB) irepository.IPersonRepository {
	return &personRepository{
		db: db,
	}
}

func (ctx *personRepository) Save(person *model.Person) error {
	tx := ctx.db.MustBegin()
	_, err := tx.Exec("INSERT INTO person (first_name, last_name, email) VALUES ($1, $2, $3)", person.FirstName, person.LastName, person.Email)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (ctx *personRepository) Update(Id int, person *model.Person) error {
	tx := ctx.db.MustBegin()
	_, err := tx.Exec("UPDATE person SET first_name = $1, last_name = $2, email = $3 WHERE Id = $4", person.FirstName, person.LastName, person.Email, Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil

}
func (ctx *personRepository) Delete(Id int) error {
	tx := ctx.db.MustBegin()
	_, err := tx.Exec("DELETE FROM person WHERE Id = $1;", Id)
	if err != nil {
		tx.Rollback()
		return err
	}
	tx.Commit()
	return nil
}
func (ctx *personRepository) FindByID(Id int) (*model.Person, error) {
	p := model.Person{}
	err := ctx.db.Get(&p, "SELECT * FROM person where Id = $1", Id)
	if err != nil {
		return nil, err
	}
	return &p, nil

}
func (ctx *personRepository) FindByEmail(Email string) (*model.Person, error) {
	p := model.Person{}
	err := ctx.db.Get(&p, "SELECT * FROM person where email = $1", Email)
	if err != nil {
		return nil, err
	}
	return &p, nil
}
func (ctx *personRepository) FindAll() (*model.Persons, error) {
	pp := model.Persons{}

	err := ctx.db.Select(&pp, "SELECT * FROM person")
	if err != nil {
		return nil, err
	}
	return &pp, nil
}
