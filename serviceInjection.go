package main

import (
	"awesomeProject/domin/repository/mssql"
	"awesomeProject/domin/service"
	"awesomeProject/ui/controller"

	"github.com/kataras/iris/mvc"
)

//Person ...
func Person(app *mvc.Application) {

	// Create our movie repository with some (memory) data from the datasource.
	repo := mssql.NewPersonRepository(GetDB())
	// Create our movie service, we will bind it to the movie app's dependencies.
	personService := service.NewPersonService(repo)
	app.Register(personService)

	// serve our movies controller.
	// Note that you can serve more than one controller
	// you can also create child mvc apps using the `movies.Party(relativePath)` or `movies.Clone(app.Party(...))`
	// if you want.
	app.Handle(new(controller.PersonController))
}
