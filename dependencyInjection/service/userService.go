package service

import (
	"dependencyInjection/config"
	"dependencyInjection/model"
	"dependencyInjection/repository"
)

type PersonService struct {
	config     *config.Config
	repository *repository.PersonRepository
}

func (service *PersonService) FindAll() []*model.Person {
	if service.config.Enabled {
		return service.repository.FindAll()
	}

	return []*model.Person{}
}

func NewPersonService(config *config.Config, repository *repository.PersonRepository) *PersonService {
	return &PersonService{config: config, repository: repository}
}
