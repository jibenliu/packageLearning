package main

import (
	"dependencyInjection/config"
	"dependencyInjection/repository"
	"dependencyInjection/server"
	"dependencyInjection/service"
)

func main() {
	cfg := config.NewConfig()
	db, err := config.ConnectDatabase(cfg)
	if err != nil {
		panic(err)
	}

	personRepository := repository.NewPersonRepository(db)
	personService := service.NewPersonService(cfg, personRepository)
	sv := server.NewServer(cfg, personService)
	sv.Run()
}
