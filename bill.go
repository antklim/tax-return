package taxreturn

// Bill describes a bill for a period of time.
type Bill struct {
	Period    Period
	AmountDue float32
	Paid      float32
}

// PaidPerDay shows an average payment amount per day.
func (b Bill) PaidPerDay() float32 {
	days := b.Period.Days()
	return b.Paid / float32(days)
}

// AmountPaidIn returns amount paid in financial year.
func (b Bill) AmountPaidIn(fy FinancialYear) float32 {
	var paidForPeriod float32 = 0.0
	switch {
	// TODO: add test case to check that the payment in other tax period does not included.
	// case b.Period.Start.After(tp.End), b.Period.End.Before(fy.Start):
	case b.Period.Start.Before(fy.Start) && b.Period.End.After(fy.Start):
		daysInPeriod := b.Period.End.Sub(fy.Start) + 1
		paidForPeriod = b.PaidPerDay() * float32(daysInPeriod)
	case b.Period.Start.Before(fy.End) && b.Period.End.After(fy.End):
		daysInPeriod := fy.End.Sub(b.Period.Start) + 1
		paidForPeriod = b.PaidPerDay() * float32(daysInPeriod)
	default:
		paidForPeriod = b.Paid
	}
	return paidForPeriod
}

// Bills describes a list of bills.
type Bills []Bill

// AmountPaidIn returns amount paid in financial year by the list of bills.
func (bb Bills) AmountPaidIn(fy FinancialYear) float32 {
	var sum float32 = 0.0

	for _, b := range bb {
		sum += b.AmountPaidIn(fy)
	}

	return sum
}
