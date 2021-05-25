package services

import "digExam/chapter2/schema"

type Person interface {
	FindAll() []*schema.Person
}
