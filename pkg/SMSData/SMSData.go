package SMSData

type SMSData struct {
	Сountry      string
	Bandwidth    string
	ResponseTime string
	Provider     string
}

func NewSMSData(country string, bandwidth string, responseTime string, provider string) *SMSData {
	data := new(SMSData)
	data.Сountry = country
	data.Bandwidth = bandwidth
	data.ResponseTime = responseTime
	data.Provider = provider
	return data
}
