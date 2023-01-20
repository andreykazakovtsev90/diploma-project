package response

import (
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/BillingData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/EmailData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/IncidentData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/MMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/SMSData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/SupportData"
	"github.com/andreykazakovtsev90/diploma-project/pkg/data/VoiceCallData"
)

type ResultT struct {
	Status bool        `json:"status"` // True, если все этапы сбора данных прошли успешно, False во всех остальных случаях
	Data   *ResultSetT `json:"data"`   // Заполнен, если все этапы сбора  данных прошли успешно, nil во всех остальных случаях
	Error  string      `json:"error"`  // Пустая строка, если все этапы сбора данных прошли успешно, в случае ошибки заполнено текстом ошибки
}

type ResultSetT struct {
	SMS       [][]SMSData.SMSData           `json:"sms"`
	MMS       [][]MMSData.MMSData           `json:"mms"`
	VoiceCall []VoiceCallData.VoiceCallData `json:"voice_call"`
	Email     [][]EmailData.EmailData       `json:"email"`
	Billing   BillingData.BillingData       `json:"billing"`
	Support   []int                         `json:"support"`
	Incidents []IncidentData.IncidentData   `json:"incident"`
}

func NewResultSetT() *ResultSetT {
	res := new(ResultSetT)
	res.SMS = make([][]SMSData.SMSData, 2)
	res.MMS = make([][]MMSData.MMSData, 2)
	res.VoiceCall = make([]VoiceCallData.VoiceCallData, 0)
	res.Email = make([][]EmailData.EmailData, 0)
	res.Billing = *BillingData.NewBillingData()
	res.Support = make([]int, 0)
	res.Incidents = make([]IncidentData.IncidentData, 0)
	return res
}

func (r *ResultSetT) SetSMS(data []SMSData.SMSData) {
	for i := range data {
		d := &data[i]
		d.ModifyCountry()
	}
	size := len(data)
	d1 := make([]SMSData.SMSData, size)
	copy(d1, data)
	r.SMS[0] = SMSData.SortByProvider(d1)
	d2 := make([]SMSData.SMSData, size)
	copy(d2, d1)
	r.SMS[1] = SMSData.SortByCountry(d2)
}

func (r *ResultSetT) SetMMS(data []MMSData.MMSData) {
	for i := range data {
		d := &data[i]
		d.ModifyCountry()
	}
	size := len(data)
	d1 := make([]MMSData.MMSData, size)
	copy(d1, data)
	r.MMS[0] = MMSData.SortByProvider(d1)
	d2 := make([]MMSData.MMSData, size)
	copy(d2, d1)
	r.MMS[1] = MMSData.SortByCountry(d2)
}

func (r *ResultSetT) SetVoiceCall(data []VoiceCallData.VoiceCallData) {
	r.VoiceCall = data
}

func (r *ResultSetT) SetEmail(data []EmailData.EmailData) {
	m := make(map[string][][]EmailData.EmailData)
	for i := range data {
		if v, ok := m[data[i].Country]; !ok {
			v = make([][]EmailData.EmailData, 2)
			v[0] = make([]EmailData.EmailData, 0)
			v[0] = append(v[0], data[i])
			v[1] = make([]EmailData.EmailData, 0)
			v[1] = append(v[1], data[i])
			m[data[i].Country] = v
		} else {
			if len(v[0]) < 3 {
				v[0] = append(v[0], data[i])
			} else {
				min := data[i]
				if v[0][0].DeliveryTime > min.DeliveryTime {
					v[0][0], min = min, v[0][0]
				}
				if v[0][1].DeliveryTime > min.DeliveryTime {
					v[0][1], min = min, v[0][1]
				}
				if v[0][2].DeliveryTime > min.DeliveryTime {
					v[0][2], min = min, v[0][2]
				}
			}
			if len(v[1]) < 3 {
				v[1] = append(v[1], data[i])
			} else {
				max := data[i]
				if v[1][0].DeliveryTime < max.DeliveryTime {
					v[1][0], max = max, v[1][0]
				}
				if v[1][1].DeliveryTime < max.DeliveryTime {
					v[1][1], max = max, v[1][1]
				}
				if v[1][2].DeliveryTime < max.DeliveryTime {
					v[1][2], max = max, v[1][2]
				}
			}
		}
	}
	r.Email = make([][]EmailData.EmailData, 0)
	for _, v := range m {
		r.Email = append(r.Email, v...)
	}
}

func (r *ResultSetT) SetSupport(data []*SupportData.SupportData) {
}

func (r *ResultSetT) SetIncidents(data []IncidentData.IncidentData) {
	r.Incidents = IncidentData.SortByStatus(data)
}
