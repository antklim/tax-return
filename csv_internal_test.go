package taxreturn

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stretchr/testify/require"
)

func TestCsvRecordToBill(t *testing.T) {
	start := "2020-01-14"
	end := "2020-02-13"
	due := "123.45"
	paid := "678.9"

	record := []string{start, end, due, paid}
	period, err := NewBillPeriod(start, end)
	require.NoError(t, err)

	expected := Bill{
		Period: period,
		Due:    123.45,
		Paid:   678.9,
	}

	bill, err := csvRecordToBill(record)
	require.NoError(t, err)
	assert.Equal(t, expected, *bill)
}
