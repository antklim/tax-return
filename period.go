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
