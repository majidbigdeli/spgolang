package model

import (
	"database/sql"
)

//Person Model
type Person struct {
	ID        int    `db:"id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string
}

//Place Model
type Place struct {
	Country string
	City    sql.NullString
	TelCode int
}

//Persons Model
type Persons []Person
