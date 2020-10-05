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

// NewBillPeriod creates a new bill period, start and end dates are in fomat YYYY-MM-DD.
func NewBillPeriod(periodStart, periodEnd string) (BillPeriod, error) {
	layout := "2006-01-02"

	start, err := time.Parse(layout, periodStart)
	if err != nil {
		return BillPeriod{}, err
	}

	end, err := time.Parse(layout, periodEnd)
	if err != nil {
		return BillPeriod{}, err
	}

	return BillPeriod{start: start, end: end}, nil
}

// Days calculates amount of days in the period. Start and end dates are counted as part of period.
func (p BillPeriod) Days() int {
	return DaysInPeriod(p.Start(), p.End()) + 1
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

// DaysInPeriod calculates amount of days between start and end date (end date excluded).
func DaysInPeriod(start, end time.Time) int {
	days := end.Sub(start).Hours() / 24
	return int(days)
}
