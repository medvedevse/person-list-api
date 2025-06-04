package webapi

import (
	"github.com/medvedevse/person-list-api/internal/entity"

	"go.uber.org/zap"
)

// TODO: Probably it could be written better
type Person struct {
	Person entity.Person
}

func AddPersonData(l *zap.Logger, entityPerson *entity.Person) {
	l.Info("Calling a third-party API for data enrichment")
	// TODO: Probably it could be written better
	person := Person{Person: *entityPerson}
	person.AddPersonAge(l)
	person.AddPersonNationality(l)
	person.AddPersonGender(l)

	// TODO: Probably it could be written better
	*entityPerson = person.Person
	l.Info("Received data from a third-party API")
}
