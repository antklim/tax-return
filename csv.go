package taxreturn

import (
	"encoding/csv"
	"os"
	"strconv"
)

type csvField int

const (
	csvFieldStart csvField = iota
	csvFieldEnd
	csvFieldDue
	csvFieldPaid
)

// ReadCsv reads CSV file and maps it to Bills.
func ReadCsv(file *os.File, hasHeader bool) (Bills, error) {
	reader := csv.NewReader(file)

	if hasHeader {
		_, err := reader.Read()
		if err != nil {
			return nil, err
		}
	}

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	bills := make(Bills, len(records))
	for i, record := range records {
		bill, err := csvRecordToBill(record)
		if err != nil {
			return nil, err
		}
		bills[i] = *bill
	}

	return bills, nil
}

func csvRecordToBill(record []string) (*Bill, error) {
	period, err := NewBillPeriod(record[csvFieldStart], record[csvFieldEnd])
	if err != nil {
		return nil, err
	}

	paid, err := strconv.ParseFloat(record[csvFieldPaid], 32)
	if err != nil {
		return nil, err
	}

	due, err := strconv.ParseFloat(record[csvFieldDue], 32)
	if err != nil {
		return nil, err
	}

	return &Bill{
		Period: period,
		Due:    float32(due),
		Paid:   float32(paid),
	}, nil
}
