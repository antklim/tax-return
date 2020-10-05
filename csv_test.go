package taxreturn_test

import (
	"os"
	"testing"

	taxreturn "github.com/antklim/tax-return"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var billPeriod1, _ = taxreturn.NewBillPeriod("2020-01-14", "2020-02-13")
var billPeriod2, _ = taxreturn.NewBillPeriod("2020-02-14", "2020-03-13")

func TestReadCsv(t *testing.T) {
	testCases := []struct {
		desc      string
		csvFile   string
		hasHeader bool
		expected  taxreturn.Bills
	}{
		{
			desc:      "reads file with header",
			csvFile:   "test/records.csv",
			hasHeader: true,
			expected: taxreturn.Bills{
				{
					Period: billPeriod1,
					Due:    100.01,
					Paid:   101.00,
				},
			},
		},
		{
			desc:      "reads file without header",
			csvFile:   "test/no_header_records.csv",
			hasHeader: false,
			expected: taxreturn.Bills{
				{
					Period: billPeriod2,
					Due:    123.45,
					Paid:   123.50,
				},
			},
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			file, err := os.Open(tC.csvFile)
			require.NoError(t, err)
			actual, err := taxreturn.ReadCsv(file, tC.hasHeader)
			require.NoError(t, err)
			assert.ElementsMatch(t, tC.expected, actual)
		})
	}
}
