package countryReference

import (
	"encoding/json"
	"io/ioutil"
)

var countryRef *CountryReference

type CountryReference struct {
	countries map[string]string
}

type ref struct {
	Countries []*Country `json:"countries"`
}

type Country struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

func Init(filename string) error {
	countryRef = new(CountryReference)
	countryRef.countries = make(map[string]string, 0)
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
		countryRef.countries[country.Code] = country.Name
	}
	return nil
}

func Get(code string) (string, bool) {
	c, ok := countryRef.countries[code]
	return c, ok
}

func Contains(code string) bool {
	_, ok := countryRef.countries[code]
	return ok
}
