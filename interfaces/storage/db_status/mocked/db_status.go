package mocked

import "github.com/jclebreton/opensirene/domain"

type RW struct {
	FindDatabaseStatusRet
	GetSuccessListRet
}

func (rw RW) FindDatabaseStatus() ([]domain.UpdateFileStatus, error) {
	return rw.FindDatabaseStatusRet.DBStatus, rw.FindDatabaseStatusRet.Err
}

type FindDatabaseStatusRet struct {
	DBStatus []domain.UpdateFileStatus
	Err      error
}

func (rw RW) GetSuccessList() ([]string, error) {
	return rw.GetSuccessListRet.DBStatus, rw.GetSuccessListRet.Err
}

type GetSuccessListRet struct {
	DBStatus []string
	Err      error
}
