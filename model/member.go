package model

type Member struct {
	StudentId string `db:"student_id" json:"student_id"`
	Name      string `db:"name" json:"name"`
}
