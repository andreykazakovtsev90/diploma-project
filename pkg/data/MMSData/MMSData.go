package MMSData

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/providerReference"
	"strings"
)

type MMSData struct {
	Country      string `json:"country"`       // alpha-2 — код страны
	Provider     string `json:"provider"`      // название компании-провайдера
	Bandwidth    string `json:"bandwidth"`     // пропускная способность канала от 0 до 100%
	ResponseTime string `json:"response_time"` // среднее время ответа в миллисекундах
}

// Возвращает список валидных данных о системе MMS
func (d *MMSData) IsValid() bool {
	if !countryReference.Contains(d.Country) {
		return false
	}
	provider, ok := providerReference.Get(d.Provider)
	if !ok || !provider.IsMMS {
		return false
	}
	return true
}

func (d *MMSData) ModifyCountry() {
	d.Country, _ = countryReference.Get(d.Country)
}

func SortByProvider(data []MMSData) []MMSData {
	for i := 1; i < len(data); i++ {
		j := i
		for j > 0 && strings.Compare(data[j].Provider, data[j-1].Provider) < 0 {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
	return data
}

func SortByCountry(data []MMSData) []MMSData {
	for i := 1; i < len(data); i++ {
		j := i
		for j > 0 && strings.Compare(data[j].Country, data[j-1].Country) < 0 {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
	return data
}
