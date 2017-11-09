package mocked

import "github.com/jclebreton/opensirene/domain"

type RW struct {
	FindDatabaseStatusRet
}

func (rw RW) FindDatabaseStatus() ([]domain.UpdateFileStatus, error) {
	return rw.FindDatabaseStatusRet.DBStatus, rw.FindDatabaseStatusRet.Err
}

type FindDatabaseStatusRet struct {
	DBStatus []domain.UpdateFileStatus
	Err      error
}
