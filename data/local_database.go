package data

import (
	"MemberRegisterApi/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type (
	LocalDatabase interface {
		SignUp(member model.Member) (sql.Result, error)
		GetAllMembers() []model.Member
		GetMemberById(studentId string) *model.Member
		UpdateMember(member model.Member) (sql.Result, error)
		DeleteMember(studentId string) (sql.Result, error)
	}

	DefaultLocalDatabase struct {
		database    *sql.DB
		memberTable string
		logTable    string
		tables      []string
	}
)

func NewLocalDatabase() *DefaultLocalDatabase {
	var err error

	instance := &DefaultLocalDatabase{}
	instance.database, err = sql.Open(
		"sqlite3",
		"./nuno-lab-entry.db",
	)

	if err != nil {
		panic(err)
	}

	instance.memberTable = "member"
	instance.logTable = "log"

	instance.tables = []string{
		instance.logTable,
		instance.memberTable,
	}

	instance.createMemberTable()
	instance.createLogTable()

	return instance
}

func (db *DefaultLocalDatabase) SignUp(member model.Member) (sql.Result, error) {
	query := fmt.Sprintf(
		`INSERT INTO %s (student_id, name) VALUES("%s", "%s") `,
		db.memberTable,
		member.StudentId,
		member.Name,
	)

	return db.execDb(query)
}

func (db *DefaultLocalDatabase) GetAllMembers() []model.Member {
	query := "SELECT * FROM member"

	var studentId string
	var name string
	var members []model.Member

	rows, err := db.query(query)
	if err != nil {
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&studentId, &name)
		if err != nil {
			return nil
		}
		member := model.Member{StudentId: studentId, Name: name}
		members = append(members, member)
	}

	return members
}

func (db *DefaultLocalDatabase) GetMemberById(studentId string) *model.Member {
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE student_id = '%s'`,
		db.memberTable,
		studentId,
	)

	var member model.Member

	rows, err := db.query(query)
	if err != nil {
		return nil
	}

	for rows.Next() {
		err := rows.Scan(&member.StudentId, &member.Name)
		if err == nil {
			return &member
		}
	}

	return nil
}

func (db *DefaultLocalDatabase) UpdateMember(member model.Member) (sql.Result, error) {
	query := fmt.Sprintf(
		`UPDATE %s SET name = '%s' WHERE student_id = '%s'`,
		db.memberTable,
		member.Name,
		member.StudentId,
	)

	return db.execDb(query)
}

func (db *DefaultLocalDatabase) DeleteMember(studentId string) (sql.Result, error) {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE student_id = '%s'`,
		db.memberTable,
		studentId,
	)

	return db.execDb(query)
}

func (db *DefaultLocalDatabase) createMemberTable() {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (student_id string PRIMARY KEY, name string NOT NULL)`,
		db.memberTable,
	)

	db.execDb(query)
}

func (db *DefaultLocalDatabase) createLogTable() {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (id INTEGER PRIMARY KEY AUTOINCREMENT, student_id STRING NOT NULL, event STRING NOT NULL, date STRING NOT NULL)`,
		db.logTable,
	)

	db.execDb(query)
}

func (db *DefaultLocalDatabase) execDb(query string) (sql.Result, error) {
	return db.database.Exec(query)
}

func (db *DefaultLocalDatabase) query(query string) (*sql.Rows, error) {
	return db.database.Query(query)
}
