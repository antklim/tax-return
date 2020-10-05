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

// PaidIn returns amount paid in period.
func (b Bill) PaidIn(p Period) float32 {
	var paid float32
	switch {
	case b.Period.Start().After(p.End()), b.Period.End().Before(p.Start()):
		paid = 0.0
	case b.Period.Start().Before(p.Start()) && b.Period.End().After(p.Start()):
		days := DaysInPeriod(p.Start(), b.Period.End()) + 1
		paid = b.PaidDaily() * float32(days)
	case b.Period.Start().Before(p.End()) && b.Period.End().After(p.End()):
		days := DaysInPeriod(b.Period.Start(), p.End()) + 1
		paid = b.PaidDaily() * float32(days)
	default:
		paid = b.Paid
	}
	return paid
}

// Bills describes a list of bills.
type Bills []Bill

// AmountPaidIn returns amount paid in financial year by the list of bills.
func (bb Bills) AmountPaidIn(p Period) float32 {
	var sum float32 = 0.0

	for _, b := range bb {
		sum += b.PaidIn(p)
	}

	return sum
}

// Report generates bills report.
// func (bb Bills) Report(p Period) string {
// 	total := bb.AmountPaidIn(p)

// 	for _, b := range bb {
// 		period := fmt.Sprintf()
// 	}

// 	return fmt.Sprintln(total)
// }
