package taxreturn_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	taxreturn "github.com/antklim/tax-return"
)

func TestPaidPerDay(t *testing.T) {
	period := taxreturn.Period{
		Start: time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		End:   time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC),
	}
	bill := taxreturn.Bill{
		Period:    period,
		AmountDue: 100.0,
		Paid:      120.0,
	}
	actual := bill.PaidPerDay()
	assert.Equal(t, float32(40.0), actual)
}
