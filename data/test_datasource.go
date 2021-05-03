package data

import (
	"MemberRegisterApi/model"
)

type TestDatasource struct {
	members  []model.Member
	accesses []model.Access
	logs     []model.Log
}

func NewTestDatasource() *TestDatasource {
	members := []model.Member{
		{StudentId: "e2n20010", Name: "Shaka"},
		{StudentId: "e2n20011", Name: "Spygea"},
		{StudentId: "e2n20012", Name: "Stylishnoob"},
		{StudentId: "e2n20088", Name: "Yamaton"},
	}

	accesses := []model.Access{
		{StudentId: "e2n20010", Event: model.Register()},
		{StudentId: "e2n20011", Event: model.Entry()},
		{StudentId: "e2n20012", Event: model.Delete()},
		{StudentId: "e2n20088", Event: model.Exit()},
	}

	date, _ := model.ParseDate("20/04/12-12:20:11")
	logs := []model.Log{
		{Id: "0", StudentId: "e2n20010", Event: model.Register(), Date: date},
		{Id: "1", StudentId: "e2n20011", Event: model.Entry(), Date: date},
		{Id: "2", StudentId: "e2n20012", Event: model.Delete(), Date: date},
		{Id: "3", StudentId: "e2n20088", Event: model.Exit(), Date: date},
	}

	instance := &TestDatasource{
		members:  members,
		accesses: accesses,
		logs:     logs,
	}

	return instance
}

func (t *TestDatasource) SignUp(member model.Member) error {
	return nil
}

func (t *TestDatasource) GetAllMembers() []model.Member {
	return t.members
}

func (t *TestDatasource) GetMemberById(studentId string) *model.Member {
	return nil
}

func (t *TestDatasource) GetMemberByAliasIdSource(aliasIdSource string) *model.Member {
	return nil
}

func (t *TestDatasource) UpdateMember(member model.Member) error {
	return nil
}

func (t *TestDatasource) DeleteMember(studentId string) error {
	return nil
}

func (t *TestDatasource) TakeLog(log model.Log) error {
	return nil
}

func (t *TestDatasource) GetAllLogs() []model.Log {
	return t.logs
}

func (t *TestDatasource) GetAllAccess() []model.Access {
	return t.accesses
}

func (t *TestDatasource) GetAccessById(studentId string) *model.Access {
	for _, access := range t.accesses {
		if access.StudentId == studentId {
			return &access
		}
	}

	return nil
}

func (t *TestDatasource) UpdateAccess(access model.Access) error {
	return nil
}

func (t *TestDatasource) RegisterAccess(access model.Access) error {
	return nil
}
