package SMSData

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/providerReference"
)

type SMSData struct {
	Country      string // alpha-2 — код страны
	Bandwidth    string // пропускная способность канала от 0 до 100%
	ResponseTime string // среднее время ответа в миллисекундах
	Provider     string // название компании-провайдера
}

func NewSMSData(country string, bandwidth string, responseTime string, provider string) *SMSData {
	data := new(SMSData)
	data.Country = country
	data.Bandwidth = bandwidth
	data.ResponseTime = responseTime
	data.Provider = provider
	return data
}

// Возвращает список валидных данных о системе SMS
func Parse(fields []string) (*SMSData, bool) {
	if len(fields) != 4 {
		return nil, false
	}
	if !countryReference.Contains(fields[0]) {
		return nil, false
	}
	provider, ok := providerReference.Get(fields[3])
	if !ok || !provider.IsSMS {
		return nil, false
	}
	d := NewSMSData(fields[0], fields[1], fields[2], fields[3])
	return d, true
}
