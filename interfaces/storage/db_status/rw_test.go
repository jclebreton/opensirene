package db_status_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/storage/db_status"
	tH "github.com/jclebreton/opensirene/testHelpers"
)

func TestFindDatabaseStatus(t *testing.T) {
	rand.Seed(time.Now().Unix())
	db := tH.GetConn()
	defer db.Close()

	rw := db_status.RW{GormClient: db}
	tH.Reset(rw.GormClient, db_status.UpdateStatus{})

	exp := []domain.UpdateFileStatus{}

	for i := 0; i < 10; i++ {
		h := genRandomHistoryRecord()
		rw.GormClient.Create(h)
		exp = append(exp, *h.ToUC())
	}

	returned, err := rw.FindDatabaseStatus()

	assert.NoError(t, err, "FindDatabaseStatus failed")
	assert.NoError(t, tH.CompareUpdateStatusSlices(exp, returned), "slices are different")

}

func genRandomHistoryRecord() db_status.UpdateStatus {
	return db_status.UpdateStatus{
		ID:        int32(rand.Intn(100000) + 1),
		Datetime:  time.Now().Add(time.Hour * -time.Duration(rand.Intn(100000))),
		Filename:  tH.RandString(26),
		IsSuccess: rand.Intn(2) == 1,
		Err:       tH.RandString(26),
	}
}
