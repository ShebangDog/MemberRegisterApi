package data

import (
	"MemberRegisterApi/model"
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

type (
	DefaultLocalDatabase struct {
		database    *sql.DB
		memberTable string
		logTable    string
		accessTable string
		tables      []string
	}

	//Event(interface)の都合でScanできないためdtoを使用してmodelへ変換する
	//row => modelができないからdtoを使用してmodelへ変換する
	//row => logDTO => model
	//https://stackoverflow.com/questions/57827120/golang-sql-scan-interface-scan-into-interface-field-type
	logDTO struct {
		Id        string
		StudentId string
		Event     string
		Date      string
	}

	accessDTO struct {
		StudentId string
		Event     string
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
	instance.accessTable = "access"

	instance.tables = []string{
		instance.logTable,
		instance.memberTable,
		instance.accessTable,
	}

	instance.createMemberTable()
	instance.createLogTable()
	instance.createAccessTable()

	return instance
}

func (db *DefaultLocalDatabase) SignUp(member model.Member) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (student_id, name) VALUES("%s", "%s") `,
		db.memberTable,
		member.StudentId,
		member.Name,
	)

	_, err := db.execDb(query)
	return err
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

func (db *DefaultLocalDatabase) UpdateMember(member model.Member) error {
	query := fmt.Sprintf(
		`UPDATE %s SET name = '%s' WHERE student_id = '%s'`,
		db.memberTable,
		member.Name,
		member.StudentId,
	)

	_, err := db.execDb(query)
	return err
}

func (db *DefaultLocalDatabase) DeleteMember(studentId string) error {
	query := fmt.Sprintf(
		`DELETE FROM %s WHERE student_id = '%s'`,
		db.memberTable,
		studentId,
	)

	_, err := db.execDb(query)
	return err
}

func (db *DefaultLocalDatabase) TakeLog(log model.Log) error {
	query := fmt.Sprintf(
		`INSERT INTO %s (student_id, event, date) VALUES('%s', '%s', '%s')`,
		db.logTable,
		log.StudentId,
		log.Event.Value(),
		log.Date.ToString(),
	)

	_, err := db.execDb(query)
	return err
}

func (db *DefaultLocalDatabase) GetAllLogs() []model.Log {
	query := fmt.Sprintf(
		`SELECT * FROM %s`, db.logTable,
	)

	var d logDTO
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

	var d logDTO

	rows, _ := db.query(query)
	for rows.Next() {
		rows.Scan(&d.Id, &d.StudentId, &d.Event, &d.Date)

		log := d.toModel()
		return &log
	}

	return nil
}

func (db *DefaultLocalDatabase) GetAllAccess() []model.Access {
	query := fmt.Sprintf(
		`SELECT * FROM %s`,
		db.accessTable,
	)

	var accesses []model.Access
	var d accessDTO

	rows, _ := db.query(query)
	for rows.Next() {
		rows.Scan(&d.StudentId, &d.Event)
		accesses = append(accesses, d.toModel())
	}

	return accesses
}

func (db *DefaultLocalDatabase) GetAccessById(studentId string) *model.Access {
	query := fmt.Sprintf(
		`SELECT * FROM %s WHERE student_id = '%s'`,
		db.accessTable,
		studentId,
	)

	var d accessDTO

	rows, _ := db.query(query)
	for rows.Next() {
		rows.Scan(&d.StudentId, &d.Event)
		access := d.toModel()

		return &access
	}

	return nil
}

func (db *DefaultLocalDatabase) UpdateAccess(access model.Access) {
	query := fmt.Sprintf(
		`UPDATE %s SET event = '%s' WHERE student_id = '%s'`,
		db.accessTable,
		access.Event.Value(),
		access.StudentId,
	)

	db.execDb(query)
}

func (db *DefaultLocalDatabase) RegisterAccess(access model.Access) {
	query := fmt.Sprintf(
		`INSERT INTO %s (student_id, event) VALUES('%s', '%s')`,
		db.accessTable,
		access.StudentId,
		access.Event.Value(),
	)

	db.execDb(query)
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

func (db *DefaultLocalDatabase) createAccessTable() {
	query := fmt.Sprintf(
		`CREATE TABLE IF NOT EXISTS %s (student_id STRING PRIMARY KEY, event STRING NOT NULL)`, db.accessTable,
	)

	db.execDb(query)
}

func (db *DefaultLocalDatabase) execDb(query string) (sql.Result, error) {
	return db.database.Exec(query)
}

func (db *DefaultLocalDatabase) query(query string) (*sql.Rows, error) {
	return db.database.Query(query)
}

func (d logDTO) toModel() model.Log {
	event, _ := model.ParseEvent(d.Event)
	date, _ := model.ParseDate(d.Date)

	return model.Log{Id: d.Id, StudentId: d.StudentId, Event: event, Date: date}
}

func (d accessDTO) toModel() model.Access {
	event, _ := model.ParseEvent(d.Event)

	return model.Access{StudentId: d.StudentId, Event: event}
}
