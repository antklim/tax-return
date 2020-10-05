package taxreturn

import (
	"fmt"
	"time"
)

const (
	parseLayout  string = "2006-01-02"
	stringLayout string = "2006-01-02"
)

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
	start, err := time.Parse(parseLayout, periodStart)
	if err != nil {
		return BillPeriod{}, err
	}

	end, err := time.Parse(parseLayout, periodEnd)
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

func (p BillPeriod) String() string {
	start := p.Start().Format(stringLayout)
	end := p.End().Format(stringLayout)
	return fmt.Sprintf("%s - %s", start, end)
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

func (fy FinancialYear) String() string {
	start := fy.Start().Format(stringLayout)
	end := fy.End().Format(stringLayout)
	return fmt.Sprintf("%s - %s", start, end)
}

// DaysInPeriod calculates amount of days between start and end date (end date excluded).
func DaysInPeriod(start, end time.Time) int {
	days := end.Sub(start).Hours() / 24
	return int(days)
}

// PeriodWithin retuns true when period a is within period b including period start and end dates.
func PeriodWithin(a, b Period) bool {
	return (a.Start().After(b.Start()) || a.Start().Equal(b.Start())) &&
		(a.End().Before(b.End()) || a.End().Equal(b.End()))
}

// PeriodOutside retuns true when period a is outside period b including period start and end dates.
func PeriodOutside(a, b Period) bool {
	return a.Start().After(b.End()) || a.End().Before(b.Start())
}

// PeriodOverlapsStart retuns true when period a ovelaps with the start of period b.
func PeriodOverlapsStart(a, b Period) bool {
	return a.Start().Before(b.Start()) && (a.End().After(b.Start()) || a.End().Equal(b.Start()))
}

// PeriodOverlapsEnd retuns true when period a ovelaps with the end of period b.
func PeriodOverlapsEnd(a, b Period) bool {
	return (a.Start().Before(b.End()) || a.Start().Equal(b.End())) && a.End().After(b.End())
}
