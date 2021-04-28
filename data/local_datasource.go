package data

import "MemberRegisterApi/model"

type LocalDatasource interface {
	SignUp(member model.Member) error
	GetAllMembers() []model.Member
	GetMemberById(studentId string) *model.Member
	UpdateMember(member model.Member) error
	DeleteMember(studentId string) error
	TakeLog(log model.Log) error
	GetAllLogs() []model.Log
	GetAllAccess() []model.Access
	GetAccessById(studentId string) *model.Access
	UpdateAccess(access model.Access) error
	RegisterAccess(access model.Access) error
}
