package model

import "encoding/json"

type Log struct {
	Id        string `db:"id" json:"id"`
	StudentId string `db:"student_id" json:"student_id"`
	Event     Event
	Date      Date
}

func (l *Log) UnmarshalJSON(data []byte) error {
	type Alias Log
	aux := &struct {
		Event string `json:"event"`
		Date  string `json:"date"`
		*Alias
	}{
		Alias: (*Alias)(l),
	}

	json.Unmarshal(data, &aux)
	l.Event, _ = ParseEvent(aux.Event)
	l.Date, _ = ParseDate(aux.Date)

	return nil
}

func (l *Log) MarshalJSON() ([]byte, error) {
	result := struct {
		Id        string `json:"id"`
		StudentId string `json:"student_id"`
		Event     string `json:"event"`
		Date      string `json:"date"`
	}{
		Id:        l.Id,
		StudentId: l.StudentId,
		Event:     l.Event.Value(),
		Date:      l.Date.ToString(),
	}

	return json.Marshal(result)
}
