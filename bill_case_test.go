package taxreturn_test

import (
	taxreturn "github.com/antklim/tax-return"
)

var billPeriod, _ = taxreturn.NewBillPeriod("2020-01-15", "2020-02-14")
var bill = taxreturn.Bill{
	Period: billPeriod,
	Due:    123.45,
	Paid:   123.45,
}

var billPeriodOverlaps, _ = taxreturn.NewBillPeriod("2020-06-15", "2020-07-14")
var billOverlaps = taxreturn.Bill{
	Period: billPeriodOverlaps,
	Due:    678.45,
	Paid:   567.45,
}

type billPaidInTestCase struct {
	desc     string
	bill     taxreturn.Bill
	period   taxreturn.Period
	expected float32
}

var billPaidInTestCases = []billPaidInTestCase{
	{
		desc:     "should not count payment when bill period before financial period",
		bill:     bill,
		period:   taxreturn.FinancialYearStarting(2020),
		expected: 0.0,
	},
	{
		desc:     "should not count payment when bill period after financial period",
		bill:     bill,
		period:   taxreturn.FinancialYearStarting(2018),
		expected: 0.0,
	},
	{
		desc:     "should count full payment when bill period is within financial period",
		bill:     bill,
		period:   taxreturn.FinancialYearStarting(2019),
		expected: 123.45,
	},
	{
		desc:     "should count part payment when bill period overlaps with financial period start",
		bill:     billOverlaps,
		period:   taxreturn.FinancialYearStarting(2020),
		expected: 264.810,
	},
	{
		desc:     "should count part payment when bill period overlaps with financial period end",
		bill:     billOverlaps,
		period:   taxreturn.FinancialYearEnding(2020),
		expected: 302.64,
	},
}
