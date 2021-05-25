package repository

import (
	"database/sql"
	"newTest/dependencyInjection/model"
)

type PersonRepository struct {
	database *sql.DB
}

func (repository *PersonRepository) FindAll() []*model.Person {
	rows, _ := repository.database.Query(
		`SELECT id, name, phone FROM users;`,
	)
	defer rows.Close()
	var people []*model.Person
	for rows.Next() {
		var (
			id    int
			name  string
			phone string
		)
		rows.Scan(&id, &name, &phone)
		people = append(people, &model.Person{
			Id:    id,
			Name:  name,
			Phone: phone,
		})
	}

	return people
}

func NewPersonRepository(database *sql.DB) *PersonRepository {
	return &PersonRepository{database: database}
}
