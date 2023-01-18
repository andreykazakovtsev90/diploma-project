package SMSData

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/providerReference"
)

type SMSData struct {
	Country      string
	Bandwidth    string
	ResponseTime string
	Provider     string
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
func ParseSMSData(fields []string) (*SMSData, bool) {
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
