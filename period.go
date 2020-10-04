package taxreturn

import "time"

// Period describes date period.
type Period struct {
	Start time.Time
	End   time.Time
}

// Days calculates amount of days in the period. Start and end dates are counted as part of period.
func (p Period) Days() int {
	days := p.End.Sub(p.Start).Hours() / 24
	return int(days) + 1
}

// FinancialYear describes financial year.
type FinancialYear Period

// FinancialYearStarting creates new financial year by starting year.
func FinancialYearStarting(year int) FinancialYear {
	start := time.Date(year, time.July, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year+1, time.June, 30, 0, 0, 0, 0, time.UTC)
	return FinancialYear{Start: start, End: end}
}

// FinancialYearEnding creates new financial year by ending year.
func FinancialYearEnding(year int) FinancialYear {
	start := time.Date(year-1, time.July, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(year, time.June, 30, 0, 0, 0, 0, time.UTC)
	return FinancialYear{Start: start, End: end}
}
