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
		TakeLog(log model.Log) (sql.Result, error)
		GetAllLogs() []model.Log
	}

	DefaultLocalDatabase struct {
		database    *sql.DB
		memberTable string
		logTable    string
		tables      []string
	}

	//Event(interface)の都合でScanできないためdtoを使用してmodelへ変換する
	//row => modelができないからdtoを使用してmodelへ変換する
	//row => eventDTO => model
	//https://stackoverflow.com/questions/57827120/golang-sql-scan-interface-scan-into-interface-field-type
	eventDTO struct {
		Id        string
		StudentId string
		Event     string
		Date      string
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

	var member model.Member
	var members []model.Member

	rows, err := db.query(query)
	if err != nil {
		return nil
	}
	for rows.Next() {
		err := rows.Scan(&member.StudentId, &member.Name)
		if err != nil {
			return nil
		}
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

func (db *DefaultLocalDatabase) TakeLog(log model.Log) (sql.Result, error) {
	query := fmt.Sprintf(
		`INSERT INTO %s (student_id, event, date) VALUES('%s', '%s', '%s')`,
		db.logTable,
		log.StudentId,
		log.Event.Value(),
		log.Date.ToString(),
	)

	return db.execDb(query)
}

func (db *DefaultLocalDatabase) GetAllLogs() []model.Log {
	query := fmt.Sprintf(
		`SELECT * FROM %s`, db.logTable,
	)

	var d eventDTO
	var logs []model.Log

	rows, _ := db.query(query)

	for rows.Next() {
		rows.Scan(&d.Id, &d.StudentId, &d.Event, &d.Date)

		logs = append(logs, d.toModel())
	}

	return logs
}

func (db *DefaultLocalDatabase) GetLogById(id int) *model.Log {
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE id = "%d"`,
		db.logTable,
		id,
	)

	var d eventDTO

	rows, _ := db.query(query)
	for rows.Next() {
		rows.Scan(&d.Id, &d.StudentId, &d.Event, &d.Date)

		log := d.toModel()
		return &log
	}

	return nil
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

func (d eventDTO) toModel() model.Log {
	event, _ := model.ParseEvent(d.Event)
	date, _ := model.ParseDate(d.Date)

	return model.Log{Id: d.Id, StudentId: d.StudentId, Event: event, Date: date}
}
