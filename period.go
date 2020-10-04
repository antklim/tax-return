package taxreturn

import "time"

// Period describes interface of the generic time period.
type Period interface {
	Start() time.Time
	End() time.Time
}

// BillPeriod describes date period.
type BillPeriod struct {
	start time.Time
	end   time.Time
}

// NewBillPeriod creates a new bill period.
func NewBillPeriod(start, end time.Time) BillPeriod {
	return BillPeriod{start, end}
}

// Days calculates amount of days in the period. Start and end dates are counted as part of period.
func (p BillPeriod) Days() int {
	days := p.End().Sub(p.Start()).Hours() / 24
	return int(days) + 1
}

// Start ...
func (p BillPeriod) Start() time.Time {
	return p.start
}

// End ...
func (p BillPeriod) End() time.Time {
	return p.end
}

// FinancialYear describes financial year.
type FinancialYear struct {
	start time.Time
	end   time.Time
}

// FinancialYearStarting creates new financial year by starting year.
func FinancialYearStarting(year int) FinancialYear {
	start := time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, time.June, 30, 0, 0, 0, 0, time.UTC)
	return FinancialYear{start, end}
}

// FinancialYearEnding creates new financial year by ending year.
func FinancialYearEnding(year int) FinancialYear {
	start := time.Date(year-1, time.July, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.June, 30, 0, 0, 0, 0, time.UTC)
	return FinancialYear{start, end}
}

// Start start date of financial year.
func (fy FinancialYear) Start() time.Time {
	return fy.start
}

// End end date of financial year.
func (fy FinancialYear) End() time.Time {
	return fy.end
}
