package main

import (
	"dependencyInjection/config"
	"dependencyInjection/repository"
	"dependencyInjection/server"
	"dependencyInjection/service"
	"go.uber.org/dig"
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
