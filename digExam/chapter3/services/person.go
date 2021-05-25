package services

import "digExam/chapter3/schema"

type Person interface {
	FindAll() []*schema.Person
}
