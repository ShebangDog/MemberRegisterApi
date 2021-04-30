package model

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
)

type Member struct {
	StudentId string `db:"student_id" json:"student_id"`
	AliasId   HashedIdm
	Name      string `db:"name" json:"name"`
}

func (m *Member) UnmarshalJSON(data []byte) error {
	type Alias Member
	aux := &struct {
		AliasIdSource string `json:"alias_id"`
		*Alias
	}{
		Alias: (*Alias)(m),
	}

	json.Unmarshal(data, &aux)

	m.AliasId = HashIdm(aux.AliasIdSource)

	return nil
}

func (m *Member) MarshalJSON() ([]byte, error) {
	result := struct {
		StudentId string    `json:"student_id"`
		AliasId   HashedIdm `json:"alias_id"`
		Name      string    `json:"name"`
	}{
		StudentId: m.StudentId,
		AliasId:   m.AliasId,
		Name:      m.Name,
	}

	return json.Marshal(result)
}

type HashedIdm struct {
	Value string
}

func HashIdm(idm string) HashedIdm {
	h := sha256.New()
	h.Write([]byte(idm))
	aliasId := fmt.Sprintf("%x", h.Sum(nil))

	return HashedIdm{Value: aliasId}
}
