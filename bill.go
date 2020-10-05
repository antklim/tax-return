package taxreturn

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
// func (bills Bills) Report(p Period) (string, error) {
// 	var b strings.Builder

// 	// total := bb.AmountPaidIn(p)

// 	for _, b := range bills {
// 		// s := b.Period.String()
// 	}

// 	// return fmt.Sprintln(total)

// 	return b.String(), nil
// }
