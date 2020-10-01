package taxreturn_test

import (
	"testing"
	"time"

	taxreturn "github.com/antklim/tax-return"
	"github.com/stretchr/testify/assert"
)

func TestPeriodDays(t *testing.T) {
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
			period := taxreturn.Period{Start: tC.start, End: tC.end}
			actual := period.Days()
			assert.Equal(t, tC.expected, actual)
		})
	}
}
