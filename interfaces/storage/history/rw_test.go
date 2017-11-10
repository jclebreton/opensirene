package history_test

import (
	"math/rand"
	"testing"
	"time"

	"github.com/alecthomas/assert"
	"github.com/jclebreton/opensirene/domain"
	"github.com/jclebreton/opensirene/interfaces/storage/history"
	tH "github.com/jclebreton/opensirene/testHelpers"
)

func TestFindHistories(t *testing.T) {
	rand.Seed(time.Now().Unix())
	db := tH.GetConn()
	defer db.Close()

	rw := history.RW{GormClient: db}
	tH.Reset(rw.GormClient, history.History{})

	exp := []domain.History{}

	for i := 0; i < 10; i++ {
		h := genRandomHistoryRecord()
		rw.GormClient.Create(h)
		exp = append(exp, *h.ToUC())
	}

	returned, err := rw.FindHistories()

	assert.NoError(t, err, "findHistory failed")
	assert.NoError(t, tH.CompareHistorySlices(exp, returned), "slices are different")

}

func genRandomHistoryRecord() history.History {
	return history.History{
		ID:        int32(rand.Intn(100000) + 1),
		Datetime:  time.Now().Add(time.Hour * -time.Duration(rand.Intn(100000))),
		Filename:  tH.RandString(26),
		IsSuccess: rand.Intn(2) == 1,
		Err:       tH.RandString(26),
	}
}
