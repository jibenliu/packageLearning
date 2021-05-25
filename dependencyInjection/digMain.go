package main

import (
	"go.uber.org/dig"
	"newTest/dependencyInjection/config"
	"newTest/dependencyInjection/repository"
	"newTest/dependencyInjection/server"
	"newTest/dependencyInjection/service"
)

func BuildContainer() *dig.Container {
	container := dig.New()
	container.Provide(config.NewConfig)
	container.Provide(config.ConnectDatabase)
	container.Provide(repository.NewPersonRepository)
	container.Provide(service.NewPersonService)
	container.Provide(server.NewServer)
	return container
}

func main() {
	container := BuildContainer()
	err := container.Invoke(func(server *server.Server) {
		server.Run()
	})
	if err != nil {
		panic(err)
	}
}
