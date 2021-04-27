package model

import "encoding/json"

type Access struct {
	StudentId string `db:"student_id" json:"student_id"`
	Event     Event
}

func (a *Access) UnmarshalJSON(data []byte) error {
	type Alias Access
	aux := &struct {
		Event string `json:"event"`
		*Alias
	}{
		Alias: (*Alias)(a),
	}

	json.Unmarshal(data, &aux)
	a.Event, _ = ParseEvent(aux.Event)

	return nil
}

func (a *Access) MarshalJSON() ([]byte, error) {
	result := struct {
		StudentId string `json:"student_id"`
		Event     string `json:"event"`
	}{
		StudentId: a.StudentId,
		Event:     a.Event.Value(),
	}

	return json.Marshal(result)
}
