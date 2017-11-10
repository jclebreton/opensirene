package database

import (
	"os"
	"testing"

	"encoding/csv"

	"reflect"

	"time"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var csvPath = "/tmp/pgx_copy_test.csv"
var csvRows = [][]string{{"A", "B", "C"}, {"1234", "5678", "FOO"}, {"4321", "8765", "BAR"}}

func createCSVForTest() (*os.File, error) {
	var file *os.File
	var err error

	if file, err = os.Create(csvPath); err != nil {
		return nil, err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Comma = ';'
	defer writer.Flush()

	for _, value := range csvRows {
		if err := writer.Write(value); err != nil {
			return nil, err
		}
	}

	return file, nil
}

func Test_PgxCopyFrom_Success(t *testing.T) {
	file, err := createCSVForTest()
	defer os.Remove(file.Name())
	assert.NoError(t, err)

	rows := &PgxCopyFrom{
		Path: csvPath,
		File: file,
		CallBackTriggerOnColName: []string{"C"},
		CallBackFunc: func(colValue string) (interface{}, error) {
			return colValue + "XXX", nil
		},
	}

	err = rows.Prepare()
	assert.NoError(t, err)

	var outputRows [][]interface{}
	for rows.Next() {
		row, err := rows.Values()
		assert.NoError(t, err)
		outputRows = append(outputRows, row)
	}
	assert.NoError(t, rows.Err())

	inputRows := [][]interface{}{
		{csvRows[1][0], csvRows[1][1], csvRows[1][2] + "XXX"},
		{csvRows[2][0], csvRows[2][1], csvRows[2][2] + "XXX"},
	}

	if !reflect.DeepEqual(inputRows, outputRows) {
		t.Errorf("Input rows and output rows do not equal: %v -> %v", inputRows, outputRows)
	}
}

func Test_colHasTrigger(t *testing.T) {
	rows := &PgxCopyFrom{callBackTriggerOnKeys: []int{1, 4}}
	assert.True(t, rows.colHasTrigger(1))
	assert.False(t, rows.colHasTrigger(2))
	assert.False(t, rows.colHasTrigger(3))
	assert.True(t, rows.colHasTrigger(4))
}

func Test_callTrigger(t *testing.T) {
	rows := &PgxCopyFrom{
		callBackTriggerOnKeys: []int{4},
		CallBackFunc: func(colValue string) (interface{}, error) {
			return colValue + "XXX", nil
		},
	}
	var values []interface{}
	values = append(values, true)

	v, err := rows.callTrigger(values, 4, "foo")
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{true, "fooXXX"}, v)

	v, err = rows.callTrigger(values, 1, "bar")
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{true, "bar"}, v)
}

func Test_callTrigger_with_date(t *testing.T) {
	rows := &PgxCopyFrom{
		callBackTriggerOnKeys: []int{4},
		CallBackFunc: func(colValue string) (interface{}, error) {
			if colValue == "" {
				return colValue, nil
			}
			return time.Now(), nil
		},
	}
	var values []interface{}
	values = append(values, true)
	v, err := rows.callTrigger(values, 4, "foo")
	assert.NoError(t, err)
	assert.IsType(t, time.Time{}, v[1])
}

func Test_Values(t *testing.T) {
	rows := &PgxCopyFrom{err: errors.New("foo")}
	v, err := rows.Values()
	assert.Error(t, err, "foo")
	assert.Nil(t, v)

	rows = &PgxCopyFrom{values: []interface{}{true, "bar"}}
	v, err = rows.Values()
	assert.NoError(t, err)
	assert.Equal(t, []interface{}{true, "bar"}, v)
}

func Test_Err(t *testing.T) {
	rows := &PgxCopyFrom{err: errors.New("foo")}
	assert.Error(t, rows.Err(), "foo")
}
