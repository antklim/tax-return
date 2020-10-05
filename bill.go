package taxreturn

import (
	"fmt"
	"strings"
)

// TODO: add reports

// Bill describes a bill for a period of time.
type Bill struct {
	Period BillPeriod
	Due    float32
	Paid   float32
}

// PaidDaily shows an average payment amount per day.
func (b Bill) PaidDaily() float32 {
	days := b.Period.Days()
	return b.Paid / float32(days)
}

// BilledDaysIn caclulates how many billed days were in provided period.
func (b Bill) BilledDaysIn(p Period) int {
	var days int

	switch {
	case PeriodOutside(b.Period, p):
		days = 0
	case PeriodOverlapsStart(b.Period, p):
		days = DaysInPeriod(p.Start(), b.Period.End()) + 1
	case PeriodOverlapsEnd(b.Period, p):
		days = DaysInPeriod(b.Period.Start(), p.End()) + 1
	default:
		days = b.Period.Days()
	}

	return days
}

// PaidIn returns amount paid in period.
func (b Bill) PaidIn(p Period) float32 {
	var paid float32

	if PeriodWithin(b.Period, p) {
		paid = b.Paid
	} else {
		paid = b.PaidDaily() * float32(b.BilledDaysIn(p))
	}

	return paid
}

// Bills describes a list of bills.
type Bills []Bill

// AmountPaidIn returns amount paid in financial year by the list of bills.
func (bills Bills) AmountPaidIn(p Period) float32 {
	var sum float32 = 0.0

	for _, b := range bills {
		sum += b.PaidIn(p)
	}

	return sum
}

// Report generates bills report.
func (bills Bills) Report(p Period) (string, error) {
	var b strings.Builder

	period := fmt.Sprintf("Financial period: %s\n\nBills periods:\n", p.String())
	if _, err := b.WriteString(period); err != nil {
		return "", err
	}

	for _, bill := range bills {
		s := fmt.Sprintf(
			"%s\t%10.3f (%d days x %6.3f per day)\n",
			bill.Period.String(),
			bill.PaidIn(p),
			bill.BilledDaysIn(p),
			bill.PaidDaily())

		if _, err := b.WriteString(s); err != nil {
			return "", err
		}
	}

	if _, err := b.WriteString("===\n"); err != nil {
		return "", err
	}

	total := fmt.Sprintf("Total paid in financial period\t$%.3f\n", bills.AmountPaidIn(p))
	if _, err := b.WriteString(total); err != nil {
		return "", err
	}

	return b.String(), nil
}
