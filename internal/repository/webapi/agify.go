package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"person-service/internal/entity"

	"go.uber.org/zap"
)

func (p *Person) AddPersonAge(l *zap.Logger) error {
	l.Info("Making a request to a third-party API", zap.String("url", "https://api.agify.io"))
	url := fmt.Sprintf("https://api.agify.io/?name=%v", p.Person.Name)

	res, err := http.Get(url)
	if err != nil {
		l.Error("Error receiving age data", zap.Error(err))
		return err
	}
	l.Info("Request to third-party API successfully completed", zap.String("url", "https://api.agify.io"))
	defer res.Body.Close()

	var ageData *entity.PersonAge

	if err := json.NewDecoder(res.Body).Decode(&ageData); err != nil {
		l.Error("Error converting json data", zap.Error(err))
		return err
	}

	p.Person.Age = ageData.Age
	l.Info("Age data successfully loaded")
	return nil
}
