package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/medvedevse/person-list-api/internal/entity"

	"go.uber.org/zap"
)

func (p *Person) AddPersonNationality(l *zap.Logger) error {
	l.Info("Making a request to a third-party API", zap.String("url", "https://api.nationalize.io"))
	url := fmt.Sprintf("https://api.nationalize.io/?name=%v", p.Person.Name)

	res, err := http.Get(url)
	if err != nil {
		l.Error("Error receiving nationality data", zap.Error(err))
		return err
	}

	l.Info("Request to third-party API successfully completed", zap.String("url", "https://api.nationalize.io"))
	defer res.Body.Close()

	var nationalityData *entity.NationalityCountry

	if err := json.NewDecoder(res.Body).Decode(&nationalityData); err != nil {
		l.Error("Error converting json data", zap.Error(err))
		return err
	}

	p.Person.Nationality = nationalityData.Country[0].CountryID
	l.Info("Nationality data successfully loaded")
	return nil
}
