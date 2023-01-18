package VoiceCallData

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/countryReference"
	"github.com/andreykazakovtsev90/diploma-project/pkg/providerReference"
	"strconv"
)

type VoiceCallData struct {
	Country             string  // alpha-2 — код страны
	Bandwidth           int     // текущая нагрузка в процентах
	ResponseTime        int     // среднее время ответа в миллисекундах
	Provider            string  // название компании-провайдера
	ConnectionStability float32 // стабильность соединения
	TTFB                int     // TTFB
	VoicePurity         int     // чистота TTFB-связи
	Median              int     // медиана длительности звонка
}

func NewVoiceCallData(country string, bandwidth int, responseTime int, provider string, connectionStability float32, ttfb int, voicePurity int, median int) *VoiceCallData {
	data := new(VoiceCallData)
	data.Country = country
	data.Bandwidth = bandwidth
	data.ResponseTime = responseTime
	data.Provider = provider
	data.ConnectionStability = connectionStability
	data.TTFB = ttfb
	data.VoicePurity = voicePurity
	data.Median = median
	return data
}

// Возвращает список валидных данных о системе VoiceCall
func Parse(fields []string) (*VoiceCallData, bool) {
	if len(fields) != 8 {
		return nil, false
	}
	if !countryReference.Contains(fields[0]) {
		return nil, false
	}
	bandwidth, err := strconv.Atoi(fields[1])
	if err != nil {
		return nil, false
	}
	responseTime, err := strconv.Atoi(fields[2])
	if err != nil {
		return nil, false
	}
	provider, ok := providerReference.Get(fields[3])
	if !ok || !provider.IsVoiceCall {
		return nil, false
	}
	conStab, err := strconv.ParseFloat(fields[4], 32)
	if err != nil {
		return nil, false
	}
	ttfb, err := strconv.Atoi(fields[5])
	if err != nil {
		return nil, false
	}
	purity, err := strconv.Atoi(fields[6])
	if err != nil {
		return nil, false
	}
	median, err := strconv.Atoi(fields[7])
	if err != nil {
		return nil, false
	}

	d := NewVoiceCallData(fields[0], bandwidth, responseTime, fields[3], float32(conStab), ttfb, purity, median)
	return d, true
}
