package data

import (
	"MemberRegisterApi/model"
	"errors"
	"fmt"
)

type DefaultLocalMemory struct {
	members  []model.Member
	logs     []model.Log
	accesses []model.Access
}

func NewLocalMemory() *DefaultLocalMemory {
	instance := &DefaultLocalMemory{}

	return instance
}

func (m *DefaultLocalMemory) SignUp(member model.Member) error {
	m.members = append(m.members, member)

	return nil
}

func (m *DefaultLocalMemory) GetAllMembers() []model.Member {
	return m.members
}

func (m *DefaultLocalMemory) GetMemberById(studentId string) *model.Member {
	index := getMemberIndexById(m.members, studentId)
	if index == nil {
		return nil
	}

	return &m.members[*index]
}

func (m *DefaultLocalMemory) UpdateMember(member model.Member) error {
	index := getMemberIndexById(m.members, member.StudentId)
	if index == nil {
		return errors.New(fmt.Sprintf("%s is not exist in members", member.StudentId))
	}

	target := &m.members[*index]

	*target = member

	return nil
}

func (m *DefaultLocalMemory) DeleteMember(studentId string) error {
	index := getMemberIndexById(m.members, studentId)
	if index == nil {
		return errors.New(fmt.Sprintf("%s is not exist in members", studentId))
	}

	m.members = append(m.members[:*index], m.members[*index+1:]...)

	return nil
}

func (m *DefaultLocalMemory) TakeLog(log model.Log) error {
	log.Id = fmt.Sprintf("%d", len(m.logs))

	m.logs = append(m.logs, log)

	return nil
}

func (m *DefaultLocalMemory) GetAllLogs() []model.Log {
	return m.logs
}

func (m *DefaultLocalMemory) GetAllAccess() []model.Access {
	return m.accesses
}

func (m *DefaultLocalMemory) GetAccessById(studentId string) *model.Access {
	index := getAccessIndexById(m.accesses, studentId)
	if index == nil {
		return nil
	}

	return &m.accesses[*index]
}

func (m *DefaultLocalMemory) UpdateAccess(access model.Access) {
	index := getAccessIndexById(m.accesses, access.StudentId)
	if index == nil {
		return
	}

	target := &m.accesses[*index]
	*target = access
}

func (m *DefaultLocalMemory) RegisterAccess(access model.Access) {
	m.accesses = append(m.accesses, access)
}

func getMemberIndexById(members []model.Member, studentId string) *int {
	for index, member := range members {
		if member.StudentId == studentId {
			return &index
		}
	}

	return nil
}

func getAccessIndexById(accesses []model.Access, studentId string) *int {
	for index, access := range accesses {
		if access.StudentId == studentId {
			return &index
		}
	}

	return nil
}
