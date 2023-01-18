package IncidentData

const accendentStatusActive = "active"
const accendentStatusClosed = "closed"

type IncidentData struct {
	Topic  string `json:"topic"`
	Status string `json:"status"` // возможные статусы active и closed
}

// Возвращает список валидных данных о системе истории инцидентов
func (d *IncidentData) IsValid() bool {
	if d.Status != accendentStatusActive && d.Status != accendentStatusClosed {
		return false
	}
	return true
}
