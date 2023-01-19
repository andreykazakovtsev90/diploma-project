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
	SMS       [][]SMSData.SMSData                `json:"sms"`
	MMS       [][]MMSData.MMSData                `json:"mms"`
	VoiceCall []VoiceCallData.VoiceCallData      `json:"voice_call"`
	Email     map[string][][]EmailData.EmailData `json:"email"`
	Billing   BillingData.BillingData            `json:"billing"`
	Support   []int                              `json:"support"`
	Incidents []IncidentData.IncidentData        `json:"incident"`
}

func NewResultSetT() *ResultSetT {
	res := new(ResultSetT)
	res.SMS = make([][]SMSData.SMSData, 2)
	res.MMS = make([][]MMSData.MMSData, 2)
	res.VoiceCall = make([]VoiceCallData.VoiceCallData, 0)
	res.Email = make(map[string][][]EmailData.EmailData, 0)
	res.Billing = *BillingData.NewBillingData()
	res.Support = make([]int, 0)
	res.Incidents = make([]IncidentData.IncidentData, 0)
	return res
}

func (r *ResultSetT) SetSMS(data []*SMSData.SMSData) {
	r.SMS[0] = make([]SMSData.SMSData, 0)
	r.SMS[1] = make([]SMSData.SMSData, 0)
	for _, d := range data {
		r.SMS[0] = append(r.SMS[0], *d)
	}
}

func (r *ResultSetT) SetMMS(data []*MMSData.MMSData) {
	r.MMS[0] = make([]MMSData.MMSData, 0)
	r.MMS[1] = make([]MMSData.MMSData, 0)
	for _, d := range data {
		r.MMS[0] = append(r.MMS[0], *d)
	}
}

func (r *ResultSetT) SetVoiceCall(data []*VoiceCallData.VoiceCallData) {
	for _, d := range data {
		r.VoiceCall = append(r.VoiceCall, *d)
	}
}

func (r *ResultSetT) SetEmail(data []*EmailData.EmailData) {
}

func (r *ResultSetT) SetSupport(data []*SupportData.SupportData) {
}

func (r *ResultSetT) SetIncidents(data []*IncidentData.IncidentData) {
	for _, d := range data {
		r.Incidents = append(r.Incidents, *d)
	}
}
