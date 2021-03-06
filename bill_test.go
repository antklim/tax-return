package taxreturn_test

import (
	"testing"

	taxreturn "github.com/antklim/tax-return"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPaidDaily(t *testing.T) {
	period, err := taxreturn.NewBillPeriod("2020-01-01", "2020-01-03")
	require.NoError(t, err)

	bill := taxreturn.Bill{
		Period: period,
		Due:    100.0,
		Paid:   120.0,
	}
	actual := bill.PaidDaily()
	assert.Equal(t, float32(40.0), actual)
}

func TestBilledDaysIn(t *testing.T) {
	for _, tC := range billedDaysInTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.bill.BilledDaysIn(tC.period)
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestPaidIn(t *testing.T) {
	for _, tC := range paidInTestCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := tC.bill.PaidIn(tC.period)
			assert.InDelta(t, tC.expected, actual, 0.0001)
		})
	}
}
