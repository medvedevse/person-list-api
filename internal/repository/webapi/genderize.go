package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"person-service/internal/entity"

	"go.uber.org/zap"
)

func (p *Person) AddPersonGender(l *zap.Logger) error {
	l.Info("Making a request to a third-party API", zap.String("url", "https://api.genderize.io"))
	url := fmt.Sprintf("https://api.genderize.io/?name=%v", p.Person.Name)

	res, err := http.Get(url)
	if err != nil {
		l.Error("Error receiving gender data", zap.Error(err))
		return err
	}

	l.Info("Request to third-party API successfully completed", zap.String("url", "https://api.genderize.io"))
	defer res.Body.Close()

	var genderData *entity.PersonGender

	if err := json.NewDecoder(res.Body).Decode(&genderData); err != nil {
		l.Error("Error converting json data", zap.Error(err))
		return err
	}

	p.Person.Gender = genderData.Gender
	l.Info("Gender data successfully loaded")
	return nil
}
