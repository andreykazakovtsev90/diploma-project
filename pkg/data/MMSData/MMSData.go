package MMSData

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/references/providerReference"
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
