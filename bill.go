package taxreturn

// Bill describes a bill for a period of time.
type Bill struct {
	Period    BillPeriod
	AmountDue float32
	Paid      float32
}

// PaidDaily shows an average payment amount per day.
func (b Bill) PaidDaily() float32 {
	days := b.Period.Days()
	return b.Paid / float32(days)
}

// PaidIn returns amount paid in period.
func (b Bill) PaidIn(p Period) float32 {
	var paidForPeriod float32 = 0.0
	switch {
	// TODO: add test case to check that the payment in other tax period does not included.
	// case b.Period.Start.After(tp.End), b.Period.End.Before(fy.Start):
	case b.Period.Start().Before(p.Start()) && b.Period.End().After(p.Start()):
		daysInPeriod := b.Period.End().Sub(p.Start()) + 1
		paidForPeriod = b.PaidDaily() * float32(daysInPeriod)
	case b.Period.Start().Before(p.End()) && b.Period.End().After(p.End()):
		daysInPeriod := p.End().Sub(b.Period.Start()) + 1
		paidForPeriod = b.PaidDaily() * float32(daysInPeriod)
	default:
		paidForPeriod = b.Paid
	}
	return paidForPeriod
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
