package countryRef

import (
	"encoding/json"
	"io/ioutil"
)

type CountryRef struct {
	countries map[string]string
}

type ref struct {
	Countries []*Country `json:"countries"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func NewCountryRef() *CountryRef {
	countryRef := new(CountryRef)
	countryRef.countries = make(map[string]string, 0)
	return countryRef
}

func (c *CountryRef) Init(filename string) error {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	r := ref{}
	err = json.Unmarshal(file, &r)
	if err != nil {
		return err
	}
	for _, country := range r.Countries {
		c.countries[country.Code] = country.Name
	}
	return nil
}

func (r *CountryRef) Contains(code string) bool {
	_, ok := r.countries[code]
	return ok
}
