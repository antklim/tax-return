package taxreturn_test

import (
	"testing"
	"time"

	taxreturn "github.com/antklim/tax-return"
	"github.com/stretchr/testify/assert"
)

func TestBillPeriodDays(t *testing.T) {
	testCases := []struct {
		desc     string
		start    time.Time
		end      time.Time
		expected int
	}{
		{
			desc:     "period with the same start and end date has duration of 1 day",
			start:    time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			expected: 1,
		},
		{
			desc:     "period different start and end dates",
			start:    time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
			end:      time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC),
			expected: 3,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			period := taxreturn.NewBillPeriod(tC.start, tC.end)
			actual := period.Days()
			assert.Equal(t, tC.expected, actual)
		})
	}
}

func TestFinancialYearStarting(t *testing.T) {
	layout := "2006-01-02"
	start, _ := time.Parse(layout, "2020-07-01")
	end, _ := time.Parse(layout, "2021-06-30")

	fy := taxreturn.FinancialYearStarting(2020)
	assert.True(t, fy.Start().Equal(start))
	assert.True(t, fy.End().Equal(end))
}

func TestFinancialYearEnding(t *testing.T) {
	layout := "2006-01-02"
	start, _ := time.Parse(layout, "2019-07-01")
	end, _ := time.Parse(layout, "2020-06-30")

	fy := taxreturn.FinancialYearEnding(2020)
	assert.True(t, fy.Start().Equal(start))
	assert.True(t, fy.End().Equal(end))
}
