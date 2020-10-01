package taxreturn

// Bill describes a bill for a period of time and amount.
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
