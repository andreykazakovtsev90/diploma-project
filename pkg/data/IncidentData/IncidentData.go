package IncidentData

import "strings"

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

func SortByStatus(data []IncidentData) []IncidentData {
	for i := 1; i < len(data); i++ {
		j := i
		for j > 0 && strings.Compare(data[j].Status, data[j-1].Status) < 0 {
			data[j], data[j-1] = data[j-1], data[j]
			j--
		}
	}
	return data
}
