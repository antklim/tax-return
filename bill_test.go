package taxreturn_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	taxreturn "github.com/antklim/tax-return"
)

func TestPaidDaily(t *testing.T) {
	period := taxreturn.NewBillPeriod(
		time.Date(2020, 01, 01, 0, 0, 0, 0, time.UTC),
		time.Date(2020, 01, 03, 0, 0, 0, 0, time.UTC))

	bill := taxreturn.Bill{
		Period:    period,
		AmountDue: 100.0,
		Paid:      120.0,
	}
	actual := bill.PaidDaily()
	assert.Equal(t, float32(40.0), actual)
}

func TestPaidIn(t *testing.T) {

}
