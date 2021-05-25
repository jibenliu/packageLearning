package controllers

import (
	"digExam/chapter2/services"
	"encoding/json"
	"net/http"
)

type Person struct {
	Service services.Person
}

func (p *Person) FindAll(w http.ResponseWriter, r *http.Request) {
	people := p.Service.FindAll()
	bytes, _ := json.Marshal(people)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(bytes)
}
