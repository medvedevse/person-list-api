package entity

type NationalityCountryID struct {
	CountryID string `json:"country_id"`
}

type NationalityCountry struct {
	Country []NationalityCountryID `json:"country"`
}
